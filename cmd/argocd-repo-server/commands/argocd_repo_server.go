package commands

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/vathsalashetty96/pkg/stats"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/vathsalashetty96/argo-cd/common"
	"github.com/vathsalashetty96/argo-cd/reposerver"
	"github.com/vathsalashetty96/argo-cd/reposerver/apiclient"
	reposervercache "github.com/vathsalashetty96/argo-cd/reposerver/cache"
	"github.com/vathsalashetty96/argo-cd/reposerver/metrics"
	"github.com/vathsalashetty96/argo-cd/reposerver/repository"
	cacheutil "github.com/vathsalashetty96/argo-cd/util/cache"
	"github.com/vathsalashetty96/argo-cd/util/cli"
	"github.com/vathsalashetty96/argo-cd/util/env"
	"github.com/vathsalashetty96/argo-cd/util/errors"
	"github.com/vathsalashetty96/argo-cd/util/gpg"
	"github.com/vathsalashetty96/argo-cd/util/healthz"
	ioutil "github.com/vathsalashetty96/argo-cd/util/io"
	"github.com/vathsalashetty96/argo-cd/util/tls"
)

const (
	// CLIName is the name of the CLI
	cliName         = "argocd-repo-server"
	gnuPGSourcePath = "/app/config/gpg/source"

	defaultPauseGenerationAfterFailedGenerationAttempts = 3
	defaultPauseGenerationOnFailureForMinutes           = 60
	defaultPauseGenerationOnFailureForRequests          = 0
)

func getGnuPGSourcePath() string {
	if path := os.Getenv("ARGOCD_GPG_DATA_PATH"); path != "" {
		return path
	} else {
		return gnuPGSourcePath
	}
}

func getPauseGenerationAfterFailedGenerationAttempts() int {
	return env.ParseNumFromEnv(common.EnvPauseGenerationAfterFailedAttempts, defaultPauseGenerationAfterFailedGenerationAttempts, 0, math.MaxInt32)
}

func getPauseGenerationOnFailureForMinutes() int {
	return env.ParseNumFromEnv(common.EnvPauseGenerationMinutes, defaultPauseGenerationOnFailureForMinutes, 0, math.MaxInt32)
}

func getPauseGenerationOnFailureForRequests() int {
	return env.ParseNumFromEnv(common.EnvPauseGenerationRequests, defaultPauseGenerationOnFailureForRequests, 0, math.MaxInt32)
}

func NewCommand() *cobra.Command {
	var (
		logFormat              string
		logLevel               string
		parallelismLimit       int64
		listenPort             int
		metricsPort            int
		cacheSrc               func() (*reposervercache.Cache, error)
		tlsConfigCustomizerSrc func() (tls.ConfigCustomizer, error)
		redisClient            *redis.Client
	)
	var command = cobra.Command{
		Use:               cliName,
		Short:             "Run ArgoCD Repository Server",
		Long:              "ArgoCD Repository Server is an internal service which maintains a local cache of the Git repository holding the application manifests, and is responsible for generating and returning the Kubernetes manifests.  This command runs Repository Server in the foreground.  It can be configured by following options.",
		DisableAutoGenTag: true,
		RunE: func(c *cobra.Command, args []string) error {
			cli.SetLogFormat(logFormat)
			cli.SetLogLevel(logLevel)

			tlsConfigCustomizer, err := tlsConfigCustomizerSrc()
			errors.CheckError(err)

			cache, err := cacheSrc()
			errors.CheckError(err)

			metricsServer := metrics.NewMetricsServer()
			cacheutil.CollectMetrics(redisClient, metricsServer)
			server, err := reposerver.NewServer(metricsServer, cache, tlsConfigCustomizer, repository.RepoServerInitConstants{
				ParallelismLimit: parallelismLimit,
				PauseGenerationAfterFailedGenerationAttempts: getPauseGenerationAfterFailedGenerationAttempts(),
				PauseGenerationOnFailureForMinutes:           getPauseGenerationOnFailureForMinutes(),
				PauseGenerationOnFailureForRequests:          getPauseGenerationOnFailureForRequests(),
			})
			errors.CheckError(err)

			grpc := server.CreateGRPC()
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
			errors.CheckError(err)

			healthz.ServeHealthCheck(http.DefaultServeMux, func(r *http.Request) error {
				if val, ok := r.URL.Query()["full"]; ok && len(val) > 0 && val[0] == "true" {
					// connect to itself to make sure repo server is able to serve connection
					// used by liveness probe to auto restart repo server
					// see https://github.com/argoproj/argo-cd/issues/5110 for more information
					conn, err := apiclient.NewConnection(fmt.Sprintf("localhost:%d", listenPort), 60)
					if err != nil {
						return err
					}
					defer ioutil.Close(conn)
					client := grpc_health_v1.NewHealthClient(conn)
					res, err := client.Check(r.Context(), &grpc_health_v1.HealthCheckRequest{})
					if err != nil {
						return err
					}
					if res.Status != grpc_health_v1.HealthCheckResponse_SERVING {
						return fmt.Errorf("grpc health check status is '%v'", res.Status)
					}
					return nil
				}
				return nil
			})
			http.Handle("/metrics", metricsServer.GetHandler())
			go func() { errors.CheckError(http.ListenAndServe(fmt.Sprintf(":%d", metricsPort), nil)) }()

			if gpg.IsGPGEnabled() {
				log.Infof("Initializing GnuPG keyring at %s", common.GetGnuPGHomePath())
				err = gpg.InitializeGnuPG()
				errors.CheckError(err)

				log.Infof("Populating GnuPG keyring with keys from %s", getGnuPGSourcePath())
				added, removed, err := gpg.SyncKeyRingFromDirectory(getGnuPGSourcePath())
				errors.CheckError(err)
				log.Infof("Loaded %d (and removed %d) keys from keyring", len(added), len(removed))

				go func() { errors.CheckError(reposerver.StartGPGWatcher(getGnuPGSourcePath())) }()
			}

			log.Infof("argocd-repo-server %s serving on %s", common.GetVersion(), listener.Addr())
			stats.RegisterStackDumper()
			stats.StartStatsTicker(10 * time.Minute)
			stats.RegisterHeapDumper("memprofile")
			err = grpc.Serve(listener)
			errors.CheckError(err)
			return nil
		},
	}

	command.Flags().StringVar(&logFormat, "logformat", "text", "Set the logging format. One of: text|json")
	command.Flags().StringVar(&logLevel, "loglevel", "info", "Set the logging level. One of: debug|info|warn|error")
	command.Flags().Int64Var(&parallelismLimit, "parallelismlimit", 0, "Limit on number of concurrent manifests generate requests. Any value less the 1 means no limit.")
	command.Flags().IntVar(&listenPort, "port", common.DefaultPortRepoServer, "Listen on given port for incoming connections")
	command.Flags().IntVar(&metricsPort, "metrics-port", common.DefaultPortRepoServerMetrics, "Start metrics server on given port")

	tlsConfigCustomizerSrc = tls.AddTLSFlagsToCmd(&command)
	cacheSrc = reposervercache.AddCacheFlagsToCmd(&command, func(client *redis.Client) {
		redisClient = client
	})
	return &command
}

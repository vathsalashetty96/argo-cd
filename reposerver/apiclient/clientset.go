package apiclient

import (
	"crypto/tls"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	argogrpc "github.com/vathsalashetty96/argo-cd/util/grpc"
	"github.com/vathsalashetty96/argo-cd/util/io"
)

const (
	// MaxGRPCMessageSize contains max grpc message size
	MaxGRPCMessageSize = 100 * 1024 * 1024
)

// Clientset represents repository server api clients
type Clientset interface {
	NewRepoServerClient() (io.Closer, RepoServerServiceClient, error)
}

type clientSet struct {
	address        string
	timeoutSeconds int
}

func (c *clientSet) NewRepoServerClient() (io.Closer, RepoServerServiceClient, error) {
	conn, err := NewConnection(c.address, c.timeoutSeconds)
	if err != nil {
		return nil, nil, err
	}
	return conn, NewRepoServerServiceClient(conn), nil
}

func NewConnection(address string, timeoutSeconds int) (*grpc.ClientConn, error) {
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithMax(3),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1000 * time.Millisecond)),
	}
	unaryInterceptors := []grpc.UnaryClientInterceptor{grpc_retry.UnaryClientInterceptor(retryOpts...)}
	if timeoutSeconds > 0 {
		unaryInterceptors = append(unaryInterceptors, argogrpc.WithTimeout(time.Duration(timeoutSeconds)*time.Second))
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(retryOpts...)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryInterceptors...)),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxGRPCMessageSize), grpc.MaxCallSendMsgSize(MaxGRPCMessageSize)),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Errorf("Unable to connect to repository service with address %s", address)
		return nil, err
	}
	return conn, nil
}

// NewRepoServerClientset creates new instance of repo server Clientset
func NewRepoServerClientset(address string, timeoutSeconds int) Clientset {
	return &clientSet{address: address, timeoutSeconds: timeoutSeconds}
}

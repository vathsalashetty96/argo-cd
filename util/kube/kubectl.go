package kube

import (
	"os"

	"github.com/vathsalashetty96/argo-cd/util/log"

	"github.com/vathsalashetty96/gitops-engine/pkg/utils/kube"
	"github.com/vathsalashetty96/gitops-engine/pkg/utils/tracing"
	"github.com/sirupsen/logrus"
)

var tracer tracing.Tracer = &tracing.NopTracer{}

func init() {
	if os.Getenv("ARGOCD_TRACING_ENABLED") == "1" {
		tracer = tracing.NewLoggingTracer(log.NewLogrusLogger(logrus.New()))
	}
}

func NewKubectl() kube.Kubectl {
	return &kube.KubectlCmd{Tracer: tracer, Log: log.NewLogrusLogger(logrus.New())}
}

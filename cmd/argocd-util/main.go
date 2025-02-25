package main

import (
	"fmt"
	"os"

	"github.com/vathsalashetty96/argo-cd/cmd/argocd-util/commands"

	// load the gcp plugin (required to authenticate against GKE clusters).
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// load the oidc plugin (required to authenticate with OpenID Connect).
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// load the azure plugin (required to authenticate with AKS clusters).
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
)

func main() {
	if err := commands.NewCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

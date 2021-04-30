package mocks

import (
	apiclient "github.com/vathsalashetty96/argo-cd/reposerver/apiclient"

	io "github.com/vathsalashetty96/argo-cd/util/io"
)

type Clientset struct {
	RepoServerServiceClient apiclient.RepoServerServiceClient
}

func (c *Clientset) NewRepoServerClient() (io.Closer, apiclient.RepoServerServiceClient, error) {
	return io.NopCloser, c.RepoServerServiceClient, nil
}

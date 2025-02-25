package test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	argoexec "github.com/vathsalashetty96/pkg/exec"
	"github.com/stretchr/testify/assert"

	"github.com/vathsalashetty96/argo-cd/test/fixture/test"
)

func TestKustomizeVersion(t *testing.T) {
	test.CIOnly(t)
	out, err := argoexec.RunCommand("kustomize", argoexec.CmdOpts{}, "version")
	assert.NoError(t, err)
	assert.Contains(t, out, "Version:kustomize/v3", "kustomize should be version 3")
}

// TestBuildManifests makes sure we are consistent in naming, and all kustomization.yamls are buildable
func TestBuildManifests(t *testing.T) {
	err := filepath.Walk("../manifests", func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		switch filepath.Base(path) {
		case "kustomization.yaml":
			// noop
		case "Kustomization", "kustomization.yml":
			// These are valid, but we want to be consistent with filenames
			return fmt.Errorf("Please name file 'kustomization.yaml' instead of '%s'", filepath.Base(path))
		case "Kustomize", "kustomize.yaml", "kustomize.yml":
			// These are not even valid kustomization filenames but sometimes get mistaken for them
			return fmt.Errorf("'%s' is not a valid kustomize name", filepath.Base(path))
		default:
			return nil
		}
		dirName := filepath.Dir(path)
		_, err = argoexec.RunCommand("kustomize", argoexec.CmdOpts{}, "build", dirName)
		return err
	})
	assert.NoError(t, err)
}

package e2e

import (
	"testing"

	. "github.com/vathsalashetty96/gitops-engine/pkg/sync/common"

	. "github.com/vathsalashetty96/argo-cd/pkg/apis/application/v1alpha1"
	. "github.com/vathsalashetty96/argo-cd/test/e2e/fixture/app"
	"github.com/vathsalashetty96/argo-cd/test/fixture/test"
)

func TestPatch(t *testing.T) {
	test.LocalOnly(t)
	Given(t).
		Path("two-nice-pods").
		When().
		AddFile("pod-3.yaml", `apiVersion: v1
kind: Pod
metadata:
  name: pod-3
spec:
  containers:
    - name: main
      image: alpine:3.10.2
      imagePullPolicy: IfNotPresent
      command:
        - "true"
  restartPolicy: Never
`).
		Create().
		Sync().
		Then().
		Expect(OperationPhaseIs(OperationSucceeded)).
		Expect(SyncStatusIs(SyncStatusCodeSynced)).
		When().
		DeleteFile("pod-1.yaml").
		PatchFile("pod-2.yaml", `[{"op": "add", "path": "/metadata/annotations", "value": {"bar": "Baz"}}]`).
		AddFile("pod-4.yaml", `apiVersion: v1
kind: Pod
metadata:
  name: pod-4
spec:
  containers:
    - name: main
      image: alpine:3.10.2
      imagePullPolicy: IfNotPresent
      command:
        - "true"
  restartPolicy: Never
`).
		Refresh(RefreshTypeHard).
		Then().
		Expect(SyncStatusIs(SyncStatusCodeOutOfSync))
}

package e2e

import (
	"testing"

	"github.com/vathsalashetty96/gitops-engine/pkg/health"
	. "github.com/vathsalashetty96/gitops-engine/pkg/sync/common"

	. "github.com/vathsalashetty96/argo-cd/pkg/apis/application/v1alpha1"
	. "github.com/vathsalashetty96/argo-cd/test/e2e/fixture/app"
)

// ensure that cluster scoped objects, like a cluster role, as a hok, can be successfully deployed
func TestClusterRoleBinding(t *testing.T) {
	Given(t).
		Path("cluster-role").
		When().
		Create().
		Sync().
		Then().
		Expect(OperationPhaseIs(OperationSucceeded)).
		Expect(HealthIs(health.HealthStatusHealthy)).
		Expect(SyncStatusIs(SyncStatusCodeSynced))
}

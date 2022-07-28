package e2e

import (
	. "github.com/onsi/ginkgo"
	capi_e2e "sigs.k8s.io/cluster-api/test/e2e"
)

var _ = PDescribe("When testing MachineDeployment scale out/in", func() {
	capi_e2e.MachineDeploymentScaleSpec(ctx, func() capi_e2e.MachineDeploymentScaleSpecInput {
		return capi_e2e.MachineDeploymentScaleSpecInput{
			E2EConfig:             e2eConfig,
			ClusterctlConfigPath:  clusterctlConfigPath,
			BootstrapClusterProxy: bootstrapClusterProxy,
			ArtifactFolder:        artifactFolder,
			SkipCleanup:           skipCleanup,
			Flavor:                "internal",
		}
	})
})

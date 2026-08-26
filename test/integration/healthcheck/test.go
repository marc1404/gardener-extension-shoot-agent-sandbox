// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

/**
	Overview
		- Tests the health checks for the shoot-agent-sandbox extension.
	Prerequisites
		- A Shoot exists.
	Test-case:
		1) Extension CRD
			1.1) HealthCondition Type: ShootSystemComponentsHealthy
				-  update the ManagedResource 'extension-shoot-agent-sandbox-shoot' and verify the health check conditions in the Extension CRD status.
 **/

package healthcheck

import (
	"time"

	ginkgo "github.com/onsi/ginkgo/v2"
)

const (
	timeout = 5 * time.Minute
)

var _ = ginkgo.Describe("Extension-shoot-agent-sandbox integration test: health checks", func() {
	// f := framework.NewShootFramework(nil)

	ginkgo.Context("Extension", func() {
		ginkgo.Context("Condition type: ShootSystemComponentsHealthy", func() {
			// f.Serial().Release().CIt(fmt.Sprintf("Extension CRD should contain unhealthy condition due to ManagedResource '%s' is unhealthy", constants.ManagedResourceNamesAgentShoot), func(ctx context.Context) {
			// 	err := healthcheckoperation.ExtensionHealthCheckWithManagedResource(ctx, timeout, f, "shoot-agent-sandbox", constants.ManagedResourceNamesControllerSeed, gardencorev1beta1.ShootObservabilityComponentsHealthy)
			// 	framework.ExpectNoError(err)
			// }, timeout)
		})
	})
})

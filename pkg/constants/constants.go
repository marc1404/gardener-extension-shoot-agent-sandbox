// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package constants

const (
	// ExtensionType is the name of the extension type.
	ExtensionType = "shoot-agent-sandbox"
	// ExtensionsName is the name of this Gardener extension
	ExtensionName = "gardener-extension-" + ExtensionType
	// ServiceName is the name of the service.
	ServiceName = ExtensionType

	// WebhookName is the name of this extensions webhook
	WebhookName = "shoot-agent-sandbox"
	// WebhookPath is the path the webhook is exposed on
	WebhookPath = "/webhooks/shoot-agent-sandbox"

	// ApplicationName is the name for resource describing the components deployed by the extension controller.
	ApplicationName = "agent-sandbox"

	extensionServiceName = "extension-" + ServiceName
	// NamespaceKubeSystem kube-system namespace
	NamespaceKubeSystem = "kube-system"

	// ManagedResourceNamesControllerSeed is the name used to describe the managed seed resources for the controller.
	ManagedResourceNamesControllerSeed = extensionServiceName + "-controller-seed"

	// ShootAccessSecretName is the name of the shoot access secret in the seed.
	ShootAccessSecretName = extensionServiceName
	// ShootAccessServiceAccountName is the name of the service account used for accessing the shoot.
	ShootAccessServiceAccountName = ShootAccessSecretName

	// AgentSandboxControllerChartNameSeed is the chart name for agent-sandbox controller resources in the seed.
	AgentSandboxControllerChartNameSeed = "shoot-agent-sandbox-controller-seed"
	// AgentSandboxControllerChartNameShoot is the chart name for agent-sandbox controller resources in the shoot.
	AgentSandboxControllerChartNameShoot = "shoot-agent-sandbox-controller-shoot"

	// NamespaceAgentSandbox is the namespace name for the agent-sandbox controller
	NamespaceAgentSandbox = "agent-sandbox-system"
	// ReleaseAgentSandbox is the release name for the agent-sandbox controller
	ReleaseAgentSandbox = "agent-sandbox"
	// ManagedResourceNamesAgentSandbox is the name for the managed resource for the agent-sandbox controller
	ManagedResourceNamesAgentSandbox = "agent-sandbox"
	// ManagedResourceNamesShootComponents is the name for the managed resource for the shoot-components
	ManagedResourceNamesShootComponents = "agent-sandbox-shoot-components"

	// ReleaseShootComponents is the release name for the shoot-components
	ReleaseShootComponents = "agent-sandbox-shoot-components"
)

// SPDX-FileCopyrightText: Copyright Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"os"

	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	heartbeatcmd "github.com/gardener/gardener/extensions/pkg/controller/heartbeat/cmd"
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"

	pkgcmd "github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/cmd"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/constants"
)

// ExtensionName is the name of the extension.
const ExtensionName = "shoot-agent-sandbox"

// Options holds configuration passed to the agent-sandbox controller.
type Options struct {
	generalOptions      *controllercmd.GeneralOptions
	agentSandboxOptions *pkgcmd.AgentSandboxOptions
	restOptions         *controllercmd.RESTOptions
	managerOptions      *controllercmd.ManagerOptions
	controllerOptions   *controllercmd.ControllerOptions
	lifecycleOptions    *controllercmd.ControllerOptions
	heartbeatOptions    *heartbeatcmd.Options
	controllerSwitches  *controllercmd.SwitchOptions
	reconcileOptions    *controllercmd.ReconcilerOptions
	webhookOptions      *webhookcmd.AddToManagerOptions
	optionAggregator    controllercmd.OptionAggregator
}

// NewOptions creates a new Options instance.
func NewOptions() *Options {

	webhookServerOptions := &webhookcmd.ServerOptions{
		Namespace: os.Getenv("WEBHOOK_CONFIG_NAMESPACE"),
	}

	webhookSwitches := pkgcmd.WebHookSwitchOptions()

	webhookOptions := webhookcmd.NewAddToManagerOptions(
		constants.ServiceName,
		"",
		nil,
		nil,
		webhookServerOptions,
		webhookSwitches,
	)

	options := &Options{
		generalOptions: &controllercmd.GeneralOptions{},
		restOptions:    &controllercmd.RESTOptions{},
		managerOptions: &controllercmd.ManagerOptions{
			// These are default values.
			LeaderElection:          true,
			LeaderElectionID:        controllercmd.LeaderElectionNameID(ExtensionName),
			LeaderElectionNamespace: os.Getenv("LEADER_ELECTION_NAMESPACE"),
			WebhookServerPort:       443,
			WebhookCertDir:          "/tmp/gardener-extensions-cert",
		},
		lifecycleOptions: &controllercmd.ControllerOptions{
			// This is a default value.
			MaxConcurrentReconciles: 5,
		},
		controllerOptions: &controllercmd.ControllerOptions{
			// This is a default value.
			MaxConcurrentReconciles: 5,
		},
		heartbeatOptions: &heartbeatcmd.Options{
			// This is a default value.
			ExtensionName:        ExtensionName,
			RenewIntervalSeconds: 30,
			Namespace:            os.Getenv("LEADER_ELECTION_NAMESPACE"),
		},
		reconcileOptions:    &controllercmd.ReconcilerOptions{},
		controllerSwitches:  pkgcmd.ControllerSwitches(),
		agentSandboxOptions: &pkgcmd.AgentSandboxOptions{},
		webhookOptions:      webhookOptions,
	}

	options.optionAggregator = controllercmd.NewOptionAggregator(
		options.generalOptions,
		options.agentSandboxOptions,
		options.restOptions,
		options.managerOptions,
		options.controllerOptions,
		controllercmd.PrefixOption("lifecycle-", options.lifecycleOptions),
		controllercmd.PrefixOption("heartbeat-", options.heartbeatOptions),
		options.controllerSwitches,
		options.reconcileOptions,
		options.webhookOptions,
	)

	return options
}

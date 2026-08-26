// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"

	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/constants"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/webhook"
)

// WebHookSwitchOptions returns the *webhookcmd.SwitchOptions for the webhook
func WebHookSwitchOptions() *webhookcmd.SwitchOptions {
	return webhookcmd.NewSwitchOptions(
		webhookcmd.Switch(constants.WebhookName, webhook.AddToManager),
	)
}

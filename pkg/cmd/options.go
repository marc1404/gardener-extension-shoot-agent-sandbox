// SPDX-FileCopyrightText: Copyright Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gardener/gardener/extensions/pkg/controller/cmd"
	extensionshealthcheck "github.com/gardener/gardener/extensions/pkg/controller/healthcheck"
	extensionsheartbeatcontroller "github.com/gardener/gardener/extensions/pkg/controller/heartbeat"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/apis/config"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/apis/config/v1alpha1"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/controller/healthcheck"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/controller/lifecycle"
)

var (
	scheme  *runtime.Scheme
	decoder runtime.Decoder
)

func init() {
	scheme = runtime.NewScheme()
	utilruntime.Must(config.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))

	decoder = serializer.NewCodecFactory(scheme).UniversalDecoder()
}

// AgentSandboxOptions holds options related to the agent-sandbox controller.
type AgentSandboxOptions struct {
	ConfigLocation string
	config         *AgentSandboxConfig
}

// AddFlags implements Flagger.AddFlags.
func (o *AgentSandboxOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.ConfigLocation, "config", "", "Path to agent-sandbox configuration")
}

// Complete implements Completer.Complete.
func (o *AgentSandboxOptions) Complete() error {
	if o.ConfigLocation == "" {
		return errors.New("config location is not set")
	}
	data, err := os.ReadFile(o.ConfigLocation)
	if err != nil {
		return err
	}

	config := config.Configuration{}
	_, _, err = decoder.Decode(data, nil, &config)
	if err != nil {
		return fmt.Errorf("unable to decode agent-sandbox config: %w", err)
	}

	o.config = &AgentSandboxConfig{
		config: config,
	}

	return nil
}

// Completed returns the decoded AgentSandboxConfig instance. Only call this if `Complete` was successful.
func (o *AgentSandboxOptions) Completed() *AgentSandboxConfig {
	return o.config
}

// AgentSandboxConfig contains configuration information about the agent-sandbox controller.
type AgentSandboxConfig struct {
	config config.Configuration
}

// ApplyAgentSandboxConfig applies the AgentSandboxOptions to the passed ControllerOptions instance.
func (c *AgentSandboxConfig) ApplyAgentSandboxConfig(config *config.Configuration) {
	*config = c.config
}

// ControllerSwitches are the cmd.ControllerSwitches for the extension controllers.
func ControllerSwitches() *cmd.SwitchOptions {
	return cmd.NewSwitchOptions(
		cmd.Switch(lifecycle.Name, lifecycle.AddToManager),
		cmd.Switch(extensionshealthcheck.ControllerName, healthcheck.AddToManager),
		cmd.Switch(extensionsheartbeatcontroller.ControllerName, extensionsheartbeatcontroller.AddToManager),
	)
}

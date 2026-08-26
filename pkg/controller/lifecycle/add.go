// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package lifecycle

import (
	"context"
	"time"

	"github.com/gardener/gardener/extensions/pkg/controller/extension"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/apis/config"
	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/constants"
)

const (
	// Type is the type of Extension resource.
	Type = constants.ExtensionType
	// Name is the name of the lifecycle controller.
	Name = "shoot_agent_sandbox_lifecycle_controller"
	// FinalizerSuffix is the finalizer suffix for the agent-sandbox controller.
	FinalizerSuffix = constants.ExtensionType
)

// DefaultAddOptions contains configuration for the agent-sandbox controller.
var DefaultAddOptions = AddOptions{}

// AddOptions are options to apply when adding the agent-sandbox controller to the manager.
type AddOptions struct {
	// ControllerOptions contains options for the controller.
	ControllerOptions controller.Options
	// ServiceConfig contains configuration for the agent-sandbox controller.
	ServiceConfig config.Configuration
	// IgnoreOperationAnnotation specifies whether to ignore the operation annotation or not.
	IgnoreOperationAnnotation bool
}

// AddToManager adds an agent-sandbox Lifecycle controller to the given Controller Manager.
func AddToManager(ctx context.Context, mgr manager.Manager) error {
	return extension.Add(mgr, extension.AddArgs{
		Actuator:          NewActuator(mgr, DefaultAddOptions.ServiceConfig),
		ControllerOptions: DefaultAddOptions.ControllerOptions,
		Name:              Name,
		FinalizerSuffix:   FinalizerSuffix,
		Resync:            60 * time.Minute,
		Predicates:        extension.DefaultPredicates(ctx, mgr, DefaultAddOptions.IgnoreOperationAnnotation),
		Type:              constants.ExtensionType,
	})
}

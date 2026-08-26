// SPDX-FileCopyrightText: Copyright Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"context"
	"errors"
	"fmt"

	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	extensionswebhookcontext "github.com/gardener/gardener/extensions/pkg/webhook/context"
	grmv1alpha1 "github.com/gardener/gardener/pkg/apis/config/resourcemanager/v1alpha1"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/constants"
)

const (
	configYamlKey = "config.yaml"
)

var (
	scheme *runtime.Scheme
	codec  runtime.Codec
)

func init() {
	scheme = runtime.NewScheme()
	utilruntime.Must(grmv1alpha1.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))

	var (
		ser = json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme, scheme, json.SerializerOptions{
			Yaml:   true,
			Pretty: false,
			Strict: false,
		})
		versions = schema.GroupVersions([]schema.GroupVersion{
			grmv1alpha1.SchemeGroupVersion,
			corev1.SchemeGroupVersion,
		})
	)

	codec = serializer.NewCodecFactory(scheme).CodecForVersions(ser, ser, versions, versions)
}

type mutator struct {
	client client.Client
	logger logr.Logger
}

// NewMutator creates a new agent-sandbox mutator.
func NewMutator(
	mgr manager.Manager,
	logger logr.Logger,
) extensionswebhook.Mutator {
	return &mutator{
		client: mgr.GetClient(),
		logger: logger.WithName("mutator"),
	}
}

// Mutate mutates the given object.
func (m *mutator) Mutate(ctx context.Context, new, old client.Object) error {
	// If the object does have a deletion timestamp then we don't want to mutate anything.
	if new.GetDeletionTimestamp() != nil {
		return nil
	}

	gctx := extensionswebhookcontext.NewGardenContext(m.client, new)

	cluster, err := gctx.GetCluster(ctx)
	if err != nil {
		return fmt.Errorf("failed to obtain cluster resource: %w", err)
	}

	var agentSandboxExtensionEnabled bool

	for _, extension := range cluster.Shoot.Spec.Extensions {
		if extension.Type == constants.ExtensionType {
			agentSandboxExtensionEnabled = true
			break
		}
	}

	if !agentSandboxExtensionEnabled {
		return nil
	}

	if x, ok := new.(*corev1.ConfigMap); ok {
		var oldConfigMap *corev1.ConfigMap
		if old != nil {
			var ok bool
			oldConfigMap, ok = old.(*corev1.ConfigMap)
			if !ok {
				return errors.New("could not cast old object to ConfigMap")
			}
		}

		// TODO: check if we really want to erase the immutable field from grms configmap
		x.Immutable = ptr.To(false)

		if agentSandboxExtensionEnabled {
			return mutateGardenerResourceManagerConfig(ctx, m.logger, x, oldConfigMap)
		}
	}
	return nil
}

func mutateGardenerResourceManagerConfig(_ context.Context, logger logr.Logger, configMap, _ *corev1.ConfigMap) error {
	if configMap == nil || configMap.Data == nil {
		return nil
	}

	extensionswebhook.LogMutation(logger, "ConfigMap", configMap.Namespace, configMap.Name)

	grmConfigRaw, ok := configMap.Data[configYamlKey]
	if !ok {
		return fmt.Errorf("GRM configmap does not contain key %q", configYamlKey)
	}

	grmConfig := &grmv1alpha1.ResourceManagerConfiguration{}
	_, _, err := codec.Decode([]byte(grmConfigRaw), nil, grmConfig)
	if err != nil {
		return fmt.Errorf("failed to decode GRM configuration: %w", err)
	}

	if grmConfig.TargetClientConnection != nil {
		grmConfig.TargetClientConnection.Namespaces = append(grmConfig.TargetClientConnection.Namespaces, constants.NamespaceAgentSandbox)
	}

	data, err := runtime.Encode(codec, grmConfig)
	if err != nil {
		return fmt.Errorf("failed to encode mutated GRM configuration: %w", err)
	}

	configMap.Data[configYamlKey] = string(data)

	return nil
}

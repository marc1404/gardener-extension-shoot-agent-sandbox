// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"strings"

	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/constants"
)

var logger = log.Log.WithName(constants.WebhookName)

const (
	grmConfigMapPrefix = "gardener-resource-manager-"
)

// AddToManager returns a new mutating webhook for the agent-sandbox extension.
func AddToManager(mgr manager.Manager) (*extensionswebhook.Webhook, error) {
	logger.Info("Adding webhook to manager")

	mutator := NewMutator(
		mgr,
		logger,
	)

	objTypes := []extensionswebhook.Type{
		{Obj: &corev1.ConfigMap{}},
	}

	handler, err := extensionswebhook.NewBuilder(mgr, logger).WithPredicates(isGrmConfigMap()).WithMutator(mutator, objTypes...).Build()
	if err != nil {
		return nil, err
	}

	webhook := &extensionswebhook.Webhook{
		Name:     extensionswebhook.PrefixedName(constants.WebhookName, false),
		Provider: "",
		Action:   extensionswebhook.ActionMutating,
		Path:     constants.WebhookPath,
		Target:   extensionswebhook.TargetSeed,
		Webhook:  &admission.Webhook{Handler: handler},
		Types:    objTypes,
		NamespaceSelector: &metav1.LabelSelector{
			MatchLabels: map[string]string{v1beta1constants.LabelExtensionPrefix + constants.ExtensionType: "true"},
		},
	}

	return webhook, nil
}

// isGrmConfigMap returns a predicate that filters ConfigMaps belonging to Gardener Resource Manager
func isGrmConfigMap() predicate.Predicate {
	return predicate.NewPredicateFuncs(func(obj client.Object) bool {
		configMap, ok := obj.(*corev1.ConfigMap)
		if !ok {
			return false
		}
		return strings.HasPrefix(configMap.Name, grmConfigMapPrefix)
	})
}

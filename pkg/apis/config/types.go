// SPDX-FileCopyrightText: Copyright Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	healthcheckconfigv1alpha1 "github.com/gardener/gardener/extensions/pkg/apis/config/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Configuration contains information about the GPU operator configuration.
type Configuration struct {
	metav1.TypeMeta

	// HealthCheckConfig is the config for the health check controller.
	HealthCheckConfig *healthcheckconfigv1alpha1.HealthCheckConfig
}

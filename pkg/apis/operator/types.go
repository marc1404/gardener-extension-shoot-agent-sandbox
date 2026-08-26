// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package operator

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AgentSandbox contains the configuration for the agent-sandbox controller.
type AgentSandbox struct {
	metav1.TypeMeta

	// Extensions contains the configuration for the agent-sandbox controller extensions.
	Extensions *Extensions
}

// Extensions contains the configuration for the agent-sandbox controller extensions.
type Extensions struct {
	// Enable indicates whether the agent-sandbox controller extensions should be enabled or not.
	Enable bool
}

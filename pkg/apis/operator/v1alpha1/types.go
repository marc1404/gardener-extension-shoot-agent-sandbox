// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AgentSandbox contains the configuration for the agent-sandbox controller.
type AgentSandbox struct {
	metav1.TypeMeta `json:",inline"`

	// Extensions contains the configuration for the agent-sandbox controller extensions.
	// +optional
	Extensions *Extensions `json:"extensions"`
}

// Extensions contains the configuration for the agent-sandbox controller extensions.
type Extensions struct {
	// Enable indicates whether the agent-sandbox controller extensions should be enabled or not.
	Enable bool `json:"enable"`
}

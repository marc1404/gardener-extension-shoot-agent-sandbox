// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

// +k8s:deepcopy-gen=package
// +k8s:conversion-gen=github.com/gardener/gardener-extension-shoot-agent-sandbox/pkg/apis/config
// +k8s:defaulter-gen=TypeMeta
// +k8s:openapi-gen=true

//go:generate sh -c "gen-crd-api-reference-docs -api-dir . -config ../../../../hack/api-reference/config.json -template-dir ${GARDENER_HACK_DIR}/api-reference/template -out-file ../../../../hack/api-reference/config.md"

// Package v1alpha1 contains the shoot agent-sandbox extension configuration.
// +groupName=shoot-agent-sandbox.extensions.config.gardener.cloud
package v1alpha1

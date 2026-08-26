// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

//go:generate sh -c "extension-generator --name=extension-shoot-agent-sandbox --provider-type=shoot-agent-sandbox --component-category=extension --extension-oci-repository=europe-docker.pkg.dev/gardener-project/public/charts/gardener/extensions/shoot-agent-sandbox:$(cat ../VERSION) --destination=./extension/base/extension.yaml"
//go:generate sh -c "$TOOLS_BIN_DIR/kustomize build ./extension -o ./extension.yaml"

package example

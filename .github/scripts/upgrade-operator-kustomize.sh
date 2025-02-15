#!/usr/bin/env bash

# standard bash error handling
set -o errexit  # exit immediately when a command fails.
set -E          # needs to be set if we want the ERR trap
set -o pipefail # prevents errors in a pipeline from being masked

# Expected variables:
IMG=${IMG?"Define IMG env"} # operator img
MODULE_VERSION=${MODULE_VERSION?"Define MODULE_VERSION env"} # module version used to set common labels

PROJECT_ROOT=${PWD}
CONFIG_OPERATOR=${PROJECT_ROOT}/config/operator/base

echo "ensure kustomize..."
PROJECT_ROOT=${PROJECT_ROOT} make -f hack/tools.mk kustomize

echo "upgrade ${CONFIG_OPERATOR}..."

echo "upgrade image to ${IMG}..."
cd ${CONFIG_OPERATOR}/deployment && ${PROJECT_ROOT}/bin/kustomize edit set image controller=${IMG}

echo "upgrade module version to ${MODULE_VERSION}..."
cd ${CONFIG_OPERATOR} && ${PROJECT_ROOT}/bin/kustomize edit add label app.kubernetes.io/version:${MODULE_VERSION} --force --without-selector --include-templates

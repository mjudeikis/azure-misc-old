#!/bin/bash -ex

if [[ $# -eq 0 ]]; then
    echo "usage: $0 resourcegroup [image] [custom_rg]"
    exit 1
fi

make

RESOURCE_GROUP="$1"
IMAGE="$2"
CUSTOM_RG="$3"

if [[ -z "$IMAGE" ]]; then
    IMAGE=$(az image list -g images -o json --query "[?starts_with(name, 'centos7-3.10-') && tags.valid=='true'].name | sort(@) | [-1]" | tr -d '"')
fi

if [[ -z "$CUSTOM_RG" ]]; then
    CUSTOM_RG=images
fi

RESOURCE_GROUP="$RESOURCE_GROUP" IMAGE="$IMAGE" CUSTOM_RG="$CUSTOM_RG" envsubst <_input/openshift-template.json >_input/openshift.json

time bin/acs-engine deploy -f --resource-group $1 --location eastus --subscription-id $SUBSCRIPTION_ID _input/openshift.json

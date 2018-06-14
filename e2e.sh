#!/bin/bash

#docker run --rm \
docker run -it \
  -v  ${PWD}:/go/src/github.com/Azure/acs-engine:z \
  -w /go/src/github.com/Azure/acs-engine \
  -e ORCHESTRATOR_VERSION=unstable \
  -e CLUSTER_DEFINITION=examples/openshift.json \
  -e CLIENT_ID=${CLIENT_ID} \
  -e CLIENT_SECRET=${CLIENT_SECRET} \
  -e TENANT_ID=${TENANT_ID} \
  -e SUBSCRIPTION_ID=${SUBSCRIPTION_ID} \
  -e ORCHESTRATOR=openshift \
  -e TIMEOUT=30m \
  -e CLEANUP_ON_EXIT=false \
  -e IMAGE_NAME=centos7-3.10-201806081432 \
  -e IMAGE_RESOURCE_GROUP=images \
  -e REGIONS=eastus \
  -e DISTRO=openshift39_centos \
  registry.svc.ci.openshift.org/ci/acs-engine-tests:v3.9 bash
  #-e NAME=kwoodsoncluster \
  #"make test-openshift"


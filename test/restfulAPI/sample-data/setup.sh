#!/bin/bash

# Test environment
oc new-project if-templates
oc create -f ./ovms-template.yaml
oc create -f ./storage-s3-template.yaml
oc create -f ./odhintegration-crd.yaml
oc create -f ./odhintegration-cr-openvino.yaml

# create sample data
oc new-project test-if
oc create -f ./s3-test-secret.yaml
#!/bin/bash

# Test environment
oc new-project test-if
oc create -f ./ovms-template.yaml
oc create -f ./storage-s3-template.yaml
oc create -f ./odhintegration-crd.yaml

# create sample data
oc create -f ./odhintegration-cr-openvino.yaml
oc create -f ./s3-test-secret.yaml
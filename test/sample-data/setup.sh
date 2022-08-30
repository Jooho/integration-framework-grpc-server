#!/bin/bash
cd "$(dirname "$0")"

# Test environment
oc project if-templates || oc new-project if-templates
oc create -f ../templates/ovms-template.yaml -n if-templates
oc create -f ../templates/storage-s3-template.yaml -n if-templates 
oc create -f ../templates/odhintegration-crd.yaml -n if-templates
oc create -f ../templates/odhintegration-cr-openvino.yaml -n if-templates

# create sample data
oc project test-if || oc new-project test-if
oc create -f ./s3-test-secret.yaml -n test-if
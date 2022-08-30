#!/bin/bash
cd "$(dirname "$0")"

source ./env.sh

# Install ODH operator
oc create -f ./scripts/odh-subscription.yaml

# Install Red Hat Pipeline Operator
oc create -f ./scripts/pipeline-subscription.yaml

# Deploy KFDEF for pipeline/modelmesh 
opendatahub_operator_running=1
pipeline_operator_running=1
operator_ready=1
while [ $operator_ready == 1 ]
do
  oc get pods --field-selector status.phase=Running -n openshift-operators | grep opendatahub > /dev/null
  opendatahub_operator_running=$(echo $?)

  oc get pods --field-selector status.phase=Running -n openshift-operators | grep openshift-pipelines > /dev/null
  pipeline_operator_running=$(echo $?)
  if [ ${opendatahub_operator_running} == '0' ] &&  [ ${pipeline_operator_running} == '0' ]
  then
    oc new-project odh-applications
    oc create -f ./scripts/kdf.yaml -n odh-applications
    operator_ready=0
  fi
done

# Deploy Minio pod that has a model
# oc project $namespace || oc new-project $namespace

## REDHAT MODELMESH SETUP
# sed "s/%MINIO_SECRET_KEY%/$SECRETKEY/g" ./scripts/minio-template.yaml > /tmp/minio.yaml
# oc apply -f /tmp/minio.yaml -n $namespace
# oc apply -f scripts/predictor.yaml -n $namespace
# oc apply -f scripts/inferenceservice.yaml -n $namespace

## Upstream MODELMESH SETUP
# Above is using RHODS modelmesh but I am using upstream version to use the latest api
oc project $namespace || oc new-project $namespace
cd /tmp
RELEASE=release-0.9
git clone -b $RELEASE --depth 1 --single-branch https://github.com/kserve/modelmesh-serving.git
cd modelmesh-serving

 oc adm policy add-cluster-role-to-user cluster-admin system:serviceaccount:test-if:default

./scripts/install.sh --namespace $namespace --quickstart

# # Deploy integration framework server
oc create -f ../../deploy -n namespace

# # Create sample data for openvino/modelmesh
$test_path/sample-data/setup.sh

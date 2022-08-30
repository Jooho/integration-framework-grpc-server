#!/bin/bash
cd "$(dirname "$0")"

source ./env.sh

# oc delete -f /tmp/minio.yaml -n $namespace
# oc delete -f scripts/inferenceservice.yaml -n $namespace


cd /tmp/modelmesh-serving
./scripts/delete.sh --namespace $namespace
# Model Mesh with Integration Framework Server Demo
Integration Framework Server can work with modelmesh runtime that is in ODH/RHODS.
This show how it can work together.

## Environment
Set different namespace if you want to change (Default: test-if)
~~~
vi env.sh
namespace=modelmesh-demo
~~~

## Setup Test Environment
~~~
* Install ODH operator
* Install Red Hat Pipeline Operator
* Deploy KFDEF for pipeline/modelmesh 
* Deploy Minio pod that has a model
* Deploy integration framework server
* Create sample data for openvino/modelmesh

./env-setup.sh
~~~


## Demo Flow
- Create a s3 storage for minio using integration framework server
  
  * Red Hat modelmesh case
    ~~~
    source env.sh
    AWS_SECRET_ACCESS_KEY=${SECRETKEY}
    AWS_S3_ENDPOINT="http:\/\/minio:9000"
    domain='apps.jlee-test.w1ak.s1.devshift.org'

    # using oc cli
    #oc process storage-s3 -n if-templates -p AWS_ACCESS_KEY_ID=THEACCESSKEY -p AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} -p AWS_S3_ENDPOINT=${AWS_S3_ENDPOINT} -p AWS_DEFAULT_BUCKET=modelmesh-example-models -p AWS_DEFAULT_REGION=us-south -p NAME=test-mlserver-s3 |oc create -f -n ${namespace}

    # using integration framework server
    sed "s/%AWS_SECRET_ACCESS_KEY%/$SECRETKEY/g" ./samples/secret-data.json > /tmp/secret-data.json
    sed "s/%AWS_S3_ENDPOINT%/$AWS_S3_ENDPOINT/g"  -i /tmp/secret-data.json
    curl -XPOST https://integration-framework-server-test-if.${domain}/api/v1/ns/test-if/storage/s3 --data "@/tmp/secret-data.json"|jq -r '.manifest|@base64d'|oc create -f -
    ~~~
  * Upstream ModelMesh
    ~~~
    domain='apps.jlee-test.w1ak.s1.devshift.org'

    AWS_ACCESS_KEY_ID=$(oc get secret storage-config -o json |jq -r '.data.localMinIO|@base64d'|jq -r '.access_key_id') 
    AWS_DEFAULT_BUCKET=$(oc get secret storage-config -o json |jq -r '.data.localMinIO|@base64d'|jq -r '.default_bucket') 
    AWS_DEFAULT_REGION=$(oc get secret storage-config -o json |jq -r '.data.localMinIO|@base64d'|jq -r '.region') 
    AWS_S3_ENDPOINT=$(oc get secret storage-config -o json |jq -r '.data.localMinIO|@base64d'|jq -r '.endpoint_url'| sed 's/\//\\\//g') 
    AWS_SECRET_ACCESS_KEY=$(oc get secret storage-config -o json |jq -r '.data.localMinIO|@base64d'|jq -r '.secret_access_key')
    STORAGE_TYPE=$(oc get secret storage-config -o json |jq -r '.data.localMinIO|@base64d'|jq -r '.type|ascii_upcase')

    domain='apps.jlee-test.w1ak.s1.devshift.org'
  
    sed "s/%AWS_SECRET_ACCESS_KEY%/$SECRETKEY/g" ./sample-files/secret-data.json > /tmp/secret-data.json
    sed "s/%AWS_ACCESS_KEY_ID%/$AWS_ACCESS_KEY_ID/g"  -i /tmp/secret-data.json
    sed "s/%AWS_DEFAULT_BUCKET%/$AWS_DEFAULT_BUCKET/g"  -i /tmp/secret-data.json
    sed "s/%AWS_DEFAULT_REGION%/$AWS_DEFAULT_REGION/g"  -i /tmp/secret-data.json
    sed "s/%AWS_S3_ENDPOINT%/$AWS_S3_ENDPOINT/g"  -i /tmp/secret-data.json
  
    curl -XPOST https://integration-framework-server-test-if.${domain}/api/v1/ns/test-if/storage/s3 --data "@/tmp/secret-data.json"|jq -r '.manifest|@base64d'|oc create -f -
    ~~~  

- Deploy a model with modelmesh using integration framework server
  ~~~
  cat << EOF > /tmp/params.json
  {
    "appName": "mlserver",
    "storageName": "test-mlserver-s3",
    "parameters": {
      "MODEL_PATH": "sklearn/mnist-svm.joblib",
      "MODEL_TYPE": "sklearn",
      "MODEL_SERVER_RESOURCE_NAME": "example-sklearn-isvc"
      }
  }
  EOF

  curl -XPOST  https://integration-framework-server-test-if.${domain}/api/v1/ns/test-if/modelserving  --data "@/tmp/params.json" |jq -r '.manifest|@base64d'| oc create -f -
  ~~~


TODO
-- Modelmesh is using different way about secret. It contains mutiple secret key containing parameters.
At this moment, modelserving logic try to use a secret that has parametes indivisually as a key.



  ~~~
  oc process mlserver-template -n if-templates -p STORAGE_NAME=test-mlserver-s3 -p MODEL_TYPE=sklearn -p MODEL_PATH=sklearn/mnist-svm.joblib -p MODEL_SERVER_RESOURCE_NAME=example-sklearn-isvc -p  AWS_DEFAULT_BUCKET=modelmesh-example-models -p STORAGE_TYPE=s3|oc create -f -
  ~~~

# KubeFlow Pipeline with Integration Framework Server

## Flow
Get required parameters -> Generate a pipeline yaml file

## Get Required Parameters
get_required_data.py is provided to gather required parameters to deploy a model through integration framework server.

**Arguments**
- app-name: model serving application name
- storage-name: storage name that is the name of secret to access the storage
- namespace: namespace that has storage-name

**Return**
- required parameters

**Usage**
~~~
python get_required_data.py --app-name <modelserving application name> --storage-name <secret name for storage> --namespace <namespace>
~~~

**Example**(`openvino`):
~~~
python get_required_data.py --app-name openvino --storage-name test-openvino-s3 --namespace test-if
~~~

## Generate a pipeline yaml file

1. Add the following required parameters and codes into your kubeflow pipeline python script.
~~~
 ### Add required arguments ###
 appName = "openvino"
 storageName = "test-openvino-s3"
 namespace = "test-if"
 parameters = {
     "MODEL_PATH": "model-path-value",
     "MODEL_NAME": "model-name-value",
     "BATCH_SIZE": "batch-size-value",
     "SHAPE": "shape-value",
     "MODEL_SERVER_RESOURCE_NAME": "model-server-resource-name-value",
     "MODEL_SERVER_NAMESPACE": "model-server-namespace",
 }
 
 ### Add this code snippet ###
 deploy_op = components.load_component_from_file("./component.yaml")
 def kfp_if_pipeline():
     ### Add the following function into your kubeflow pipeline python script. ###
     rhods_deploy = deploy_op(
         app_name=appName,
         storage_name=storageName,
         namespace=namespace,
         parameters=parameters,
     )
~~~

2. Execute a kubeflow pipeline python file.
  
   kfp-test image is the helper image that uses root user.

   - Example(RHODS KUBEFLOW PIPELINE)
    ~~~
    # Execute a helper container (kfp-test)
    make run-kfp-env-img

    # Generate a modelserving kubeflow pipeline yaml file
    python kfp-deploy.py

    # Review generated yaml file and upload it to kubeflow tekton pipeline server.
    cat kfp-deploy.yaml
    ~~~

## Deploy a Model manually using the script for debugging
deploy-model.py in quay.io/jooholee/rhods-integration-framework-deploy:latest is a script help deploying a model by creating a CR of a modelserving server operator.
~~~
RECREATE=true make run-kfp-image

python /home/deploy-model.py --app-name openvino --storage-name test-openvino-s3 --namespace test-if --parameters  '{"BATCH_SIZE": "batch-size-value", "MODEL_NAME": "model-name-value","MODEL_PATH": "model-path-value","MODEL_SERVER_NAMESPACE": "model-server-namespace","MODEL_SERVER_RESOURCE_NAME": "model-server-resource-name-value", "SHAPE":"shape-value"}'
~~~
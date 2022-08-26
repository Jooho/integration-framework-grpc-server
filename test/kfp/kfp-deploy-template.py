from kfp import dsl
from kfp import components

##### Update this part #####
# appName = "openvino"
# storageName = "test-openvino-s3"
# namespace = "test-if"
# parameters = {
#     "MODEL_PATH": "model-path-value",
#     "MODEL_NAME": "model-name-value",
#     "BATCH_SIZE": "batch-size-value",
#     "SHAPE": "shape-value",
#     "MODEL_SERVER_RESOURCE_NAME": "model-server-resource-name-value",
#     "MODEL_SERVER_NAMESPACE": "model-server-namespace",
# }

deploy_op = components.load_component_from_file("./component.yaml")

# Create a pipeline yaml file
@dsl.pipeline(
    name="KFP for OpenShift ODH Integration Framework Server",
    description="Kubeflow pipelines deploy model serving application through integration framework server.",
)
def kfp_if_pipeline():
    rhods_deploy = deploy_op(
        app_name=appName,
        storage_name=storageName,
        namespace=namespace,
        parameters=parameters,
    )


if __name__ == "__main__":
    # compile the pipeline
    from kfp_tekton.compiler import TektonCompiler

    TektonCompiler().compile(kfp_if_pipeline, __file__.replace(".py", ".yaml"))

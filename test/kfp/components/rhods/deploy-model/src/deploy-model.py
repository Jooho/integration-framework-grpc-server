import requests
import json
import base64

# Check necessary parameters for model serving server
# appName="openvino"
# storageName="test-openvino-s3"
# namespace="test-if"
# parameters_to_send={}

def deploy(args):
  # **Variable Example**
  # - appName="openvino"
  # - storageName="test-openvino-s3"
  # - namespace="test-if"
  # - parameters={
  #      "MODEL_PATH": "model-path-value",
  #      "MODEL_NAME": "model-name-value",
  #      "BATCH_SIZE": "batch-size-value",
  #      "SHAPE": "shape-value",
  #      "MODEL_SERVER_RESOURCE_NAME": "model-server-resource-name-value",
  #      "MODEL_SERVER_NAMESPACE": "model-server-namespace"
  #   }
  appName=args.app_name
  storageName=args.storage_name
  namespace=args.namespace
  parameters_to_send=json.loads(args.parameters)

  data_to_send={
    "appName": appName,
    "storageName": storageName,
    "parameters": parameters_to_send    
  }
  
  getRenderedManifest = requests.post("http://localhost:8000/api/v1/ns/{}/modelserving".format(namespace),data=json.dumps(data_to_send))
  
  jsonManifest = json.loads(base64.b64decode(getRenderedManifest.json()["manifest"]).decode('utf-8'))
  
  createManifest = requests.post("http://localhost:8000/api/v1/ns/{}/create".format(namespace),data=json.dumps({"manifest": getRenderedManifest.json()["manifest"]}))

  print(createManifest.json())

if __name__ == "__main__":
    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument('--app-name', type=str, required=True)
    parser.add_argument('--storage-name', type=str, required=True)
    parser.add_argument('--namespace', type=str, required=True)
    parser.add_argument('--parameters', type=str, required=True)
    args = parser.parse_args()
    deploy(args)
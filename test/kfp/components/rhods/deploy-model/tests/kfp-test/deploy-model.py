import requests
import json
import base64

# Check necessary parameters for model serving server
# appName="openvino"
# storageName="test-openvino-s3"
# namespace="test-if"
# parameters_to_send={}

def deploy(args):
  # getParameters = requests.get("http://localhost:8000/api/v1/modelserving/{}?storageName={}&namespace={}".format(appName,storageName, namespace))
  # jsonParams = json.loads(base64.b64decode(getParameters.json()["parameters"]).decode('utf-8'))

  # parameters_to_send={}
  # for param in jsonParams:
  #   required = str(param["required"]) if "required" in param else "false"
  #   print(param["name"] + ": " + param["description"] +" (required:" +required +")")
  #   parameters_to_send[param["name"]]=""

  # print("\n\nRequired JSON Data: "+json.dumps(parameters_to_send))


  # Get rendered manifest
  # parameters_to_send={
  #     "MODEL_PATH": "model-path-value",
  #     "MODEL_NAME": "model-name-value",
  #     "BATCH_SIZE": "batch-size-value",
  #     "SHAPE": "shape-value",
  #     "MODEL_SERVER_RESOURCE_NAME": "model-server-resource-name-value",
  #     "MODEL_SERVER_NAMESPACE": "model-server-namespace"
  # }
  appName=args.app_name
  storageName=args.storage_name
  namespace=args.namespace
  parameters_to_send=json.loads(args.parameters)

  data_to_send={
    "appName": appName,
    "storageName": storageName,
    "parameters": parameters_to_send    
  }
  print(data_to_send)
  getRenderedManifest = requests.post("http://localhost:8000/api/v1/ns/{}/modelserving".format(namespace),data=json.dumps(data_to_send))
  # print(renderedManifest.json())
  jsonManifest = json.loads(base64.b64decode(getRenderedManifest.json()["manifest"]).decode('utf-8'))
  # print("Manifest\n")
  # print(json.dumps(jsonManifest))

  # # Create the manifest
  # renderedManifest_to_send={
  #   "manifest": getRenderedManifest.json()["manifest"]
  # }
  createManifest = requests.post("http://localhost:8000/api/v1/ns/{}/create".format(namespace),data=json.dumps({"manifest": getRenderedManifest.json()["manifest"]}))

  print(createManifest.json())

# appName="openvino"
# storageName="test-openvino-s3"
# namespace="test-if"
if __name__ == "__main__":
    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument('--app-name', type=str, required=True)
    parser.add_argument('--storage-name', type=str, required=True)
    parser.add_argument('--namespace', type=str, required=True)
    parser.add_argument('--parameters', type=str, required=True)
    args = parser.parse_args()
    deploy(args)
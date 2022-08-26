import requests
import json
import base64

# Check necessary parameters for model serving server
appName="openvino"
storageName="test-openvino-s3"
namespace="test-if"
parameters_to_send={}
getParameters = requests.get("http://localhost:8000/api/v1/modelserving/{}?storageName={}&namespace={}".format(appName,storageName, namespace))
jsonParams = json.loads(base64.b64decode(getParameters.json()["parameters"]).decode('utf-8'))

parameters_to_send={}
for param in jsonParams:
  required = str(param["required"]) if "required" in param else "false"
  print(param["name"] + ": " + param["description"] +" (required:" +required +")")
  parameters_to_send[param["name"]]=""

print("\n\nRequired JSON Data: "+json.dumps(parameters_to_send))


# Get rendered manifest
parameters_to_send={
	  "MODEL_PATH": "model-path-value",
	  "MODEL_NAME": "model-name-value",
	  "BATCH_SIZE": "batch-size-value",
	  "SHAPE": "shape-value",
	  "MODEL_SERVER_RESOURCE_NAME": "model-server-resource-name-value",
	  "MODEL_SERVER_NAMESPACE": "model-server-namespace"
 }

data_to_send={
  "appName": appName,
  "storageName": storageName,
  "parameters": parameters_to_send    
}


# #Content type must be included in the header
# header = {"content-type": "application/json"}
getRenderedManifest = requests.post("http://localhost:8000/api/v1/ns/{}/modelserving".format(namespace),data=json.dumps(data_to_send))
# print(renderedManifest.json())
jsonManifest = json.loads(base64.b64decode(getRenderedManifest.json()["manifest"]).decode('utf-8'))
print("Manifest\n")
print(json.dumps(jsonManifest))

# Create the manifest
renderedManifest_to_send={
  "manifest": getRenderedManifest.json()["manifest"]
}
createManifest = requests.post("http://localhost:8000/api/v1/ns/{}/create".format(namespace),data=json.dumps(renderedManifest_to_send))

print(createManifest.json())

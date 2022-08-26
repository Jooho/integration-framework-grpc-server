import requests
import json
import base64

# Check necessary parameters for model serving server
def execute(args):
  """
  execute function needs 3 parameters

  :param app-name: model serving application name
  :param storage-name: storage name that is the name of secret to access the storage
  :param namespace: namespace that has storage-name
  :return: parameter list

  example command: python get_required_data.py --app-name openvino --storage-name test-openvino-s3 --namespace test-if
  """
  appName=args.app_name
  storageName=args.storage_name
  namespace=args.namespace
  getParameters = requests.get("http://localhost:8000/api/v1/modelserving/{}?storageName={}&namespace={}".format(appName,storageName, namespace))
  jsonParams = json.loads(base64.b64decode(getParameters.json()["parameters"]).decode('utf-8'))

  parameters_to_send={}
  print("************ Required Data *********")
  for param in jsonParams:
    required = str(param["required"]) if "required" in param else "false"
    print(param["name"] + ": " + param["description"] +" (required:" +required +")")
    parameters_to_send[param["name"]]=""

  print("\n\nRequired JSON Data: "+json.dumps(parameters_to_send))


if __name__ == "__main__":
    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument('--app-name', type=str, required=True)
    parser.add_argument('--storage-name', type=str, required=True)
    parser.add_argument('--namespace', type=str, required=True)
    args = parser.parse_args()
    execute(args)
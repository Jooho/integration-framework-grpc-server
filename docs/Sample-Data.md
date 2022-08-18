# Sample Data to Test
This doc contains sample test data for gGPC API call. You can use grpcui or grpcurl but grpcui would be easier. Please check [this doc for installation of the tools](./Tools.md)

## Storage
- GetRenderedStorageManifest
  ~~~
  {
  "storageType": "s3"
  }
  ~~~
- GetStorageParams
  ~~~
  {
  "storageType": "s3",
  "namespace": "test-if",
  "parameters": {
        "AWS_ACCESS_KEY_ID": "aws-key-id-value",
        "AWS_SECRET_ACCESS_KEY": "aws-secret-value",
        "AWS_S3_ENDPOINT": "aws-s3-endpoint-value",
        "AWS_DEFAULT_REGION": "aws-default-region-value",
        "AWS_DEFAULT_BUCKET": "aws-default-bucket-value",
        "NAME": "name-value" 
      }
  }
  ~~~
  
## ModelServing
- GetAppCustomResource 
  ~~~
  {
  "appName": "modelserving-openvino",
  "storageType": "s3",
  "storageName": "name-value",
  "storageNamespace": "test-if",
  "parameters": {
	  "MODEL_PATH": "model-path-value",
	  "MODEL_NAME": "model-name-value",
	  "BATCH_SIZE": "batch-size-value",
	  "SHAPE": "shape-value",
	  "MODEL_SERVER_RESOURCE_NAME": "model-server-resource-name-value",
	  "MODEL_SERVER_NAMESPACE": "model-server-namespace",
	  "STORAGE_TYPE": "S3"
    }
  }
  ~~~
- GetAppParams
  ~~~
  {
  "appName": "openvino"
  }
  ~~~
- ListApp
  ~~~
  {}
  ~~~

# K8SCall
- CreateObjectByStringJson
  ~~~

  ~~~
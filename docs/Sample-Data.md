# Sample Data to Test
This doc contains sample test data for gGPC API call. You can use grpcui or grpcurl but grpcui would be easier. Please check [this doc for installation of the tools](./Tools.md)

## Storage
- GetRenderedStorageManifest
  ~~~
  {
  "type": "s3",
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
- GetStorageParams
  ~~~
  {
  "storageType": "s3"
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
- CreateObjectByJson
  ~~~
  {
  "manifest": "eyJhcGlWZXJzaW9uIjoiaW50ZWwuY29tL3YxYWxwaGExIiwia2luZCI6Ik1vZGVsU2VydmVyIiwibWV0YWRhdGEiOnsibmFtZSI6Im1vZGVsLXNlcnZlci1yZXNvdXJjZS1uYW1lLXZhbHVlIn0sInNwZWMiOnsiZGVwbG95bWVudF9wYXJhbWV0ZXJzIjp7Im9wZW5zaGlmdF9zZXJ2aWNlX21lc2giOmZhbHNlLCJyZXBsaWNhcyI6MSwicmVzb3VyY2VzIjp7ImxpbWl0cyI6eyJ4cHVfZGV2aWNlX3F1YW50aXR5IjoiMSJ9LCJyZXF1ZXN0cyI6eyJ4cHVfZGV2aWNlX3F1YW50aXR5IjoiMSJ9fX0sImltYWdlX25hbWUiOiJyZWdpc3RyeS5jb25uZWN0LnJlZGhhdC5jb20vaW50ZWwvb3BlbnZpbm8tbW9kZWwtc2VydmVyQHNoYTI1NjpmNjcwYWEzZGMwMTRiODc4NmU1NTRiOGEzYmI3ZTJlODQ3NTc0NGQ1ODhlNWU3MmQ1NTQ2NjBiNzQ0MzBhOGM1IiwibW9kZWxzX3JlcG9zaXRvcnkiOnsiYXdzX2FjY2Vzc19rZXlfaWQiOiJhd3Mta2V5LWlkLXZhbHVlIiwiYXdzX3JlZ2lvbiI6ImF3cy1kZWZhdWx0LXJlZ2lvbi12YWx1ZSIsImF3c19zZWNyZXRfYWNjZXNzX2tleSI6ImF3cy1zZWNyZXQtdmFsdWUiLCJhenVyZV9zdG9yYWdlX2Nvbm5lY3Rpb25fc3RyaW5nIjoiIiwiZ2NwX2NyZWRzX3NlY3JldF9uYW1lIjoiIiwiaHR0cF9wcm94eSI6IiIsImh0dHBzX3Byb3h5IjoiIiwibW9kZWxzX2hvc3RfcGF0aCI6IiIsIm1vZGVsc192b2x1bWVfY2xhaW0iOiIiLCJzM19jb21wYXRfYXBpX2VuZHBvaW50IjoiYXdzLXMzLWVuZHBvaW50LXZhbHVlIiwic3RvcmFnZV90eXBlIjoiY3pNPSJ9LCJtb2RlbHNfc2V0dGluZ3MiOnsiYmF0Y2hfc2l6ZSI6ImJhdGNoLXNpemUtdmFsdWUiLCJjb25maWdfY29uZmlnbWFwX25hbWUiOiIiLCJpZGxlX3NlcXVlbmNlX2NsZWFudXAiOmZhbHNlLCJpc19zdGF0ZWZ1bCI6ZmFsc2UsImxheW91dCI6IiIsImxvd19sYXRlbmN5X3RyYW5zZm9ybWF0aW9uIjp0cnVlLCJtYXhfc2VxdWVuY2VfbnVtYmVyIjowLCJtb2RlbF9jb25maWciOiIiLCJtb2RlbF9uYW1lIjoibW9kZWwtbmFtZS12YWx1ZSIsIm1vZGVsX3BhdGgiOiJtb2RlbC1wYXRoLXZhbHVlIiwibW9kZWxfdmVyc2lvbl9wb2xpY3kiOiJ7XCJsYXRlc3RcIjogeyBcIm51bV92ZXJzaW9uc1wiOjEgfX0iLCJuaXJlcSI6MCwicGx1Z2luX2NvbmZpZyI6IntcIkNQVV9USFJPVUdIUFVUX1NUUkVBTVNcIjoxfSIsInNoYXBlIjoic2hhcGUtdmFsdWUiLCJzaW5nbGVfbW9kZWxfbW9kZSI6dHJ1ZSwidGFyZ2V0X2RldmljZSI6IkNQVSJ9LCJzZXJ2ZXJfc2V0dGluZ3MiOnsiZmlsZV9zeXN0ZW1fcG9sbF93YWl0X3NlY29uZHMiOjAsImdycGNfd29ya2VycyI6MSwibG9nX2xldmVsIjoiSU5GTyIsInJlc3Rfd29ya2VycyI6MCwic2VxdWVuY2VfY2xlYW5lcl9wb2xsX3dhaXRfbWludXRlcyI6MH0sInNlcnZpY2VfcGFyYW1ldGVycyI6eyJncnBjX3BvcnQiOjgwODAsInJlc3RfcG9ydCI6ODA4MSwic2VydmljZV90eXBlIjoiQ2x1c3RlcklQIn19fQ==",
  "namespace": "test-if"
  }
  ~~~
# Tools
In order to develop `Integration Framework Server` efficiently, many tools can be used. This doc explains useful tools for better development.




## gRPC API tools
These tools are for checking gRPC API request/response. grpcui provide UI console so it is easier to test but grpcurl is simpler to call.
- [grpcui](https://github.com/fullstorydev/grpcui)
- [grpcurl](https://github.com/fullstorydev/grpcurl)

### gRPC UI
- Installation
~~~
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
~~~

- How to use
~~~
grpcui -plaintext localhost:9000
~~~

### gRPC Curl
- Installation
~~~
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
~~~

- How to use
~~~
grpcurl -plaintext -d '{"storageType": "s3"}' --cd  localhost:9000 api.Storage/GetStorageParams
grpcurl -plaintext -d '{"storageType": "s3", "variables": [ {"value":"val1"},{"value":"val1"}]}'  localhost:9000 api.Storage/GetRenderedStorageManifest

grpcurl -d '{"id": 1234, "tags": ["foo","bar"]}' \
    grpc.server.com:443 my.custom.server.Service/Method
~~~
# Development Guide

## Setup Development Environment
- clean library (make lib-clean)
- download grpc gen well know type/gateway googleapi(make lib-download)
- install grpc related binaries(make lib-install)
~~~
make setup
~~~

## Start Grpc Server
~~~
make run 
~~~

## Compile Proto messages (v1)
~~~
# generate pb files
make gen-v1-proto

# clean pb files
make clean-v1-proto
~~~





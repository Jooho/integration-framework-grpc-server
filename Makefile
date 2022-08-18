
# Versions
PROTOC_VER=21.2
PROTOC_GEN_GO_GRPC=1.2
PROTOC_GEN_GO=1.28
PROTOC_GEN_GATEWAY=2.11.2

# Default Env
MODE?=local
LOGLEVEL?=0
# Path
TEMP_DIR=/tmp

.PHONY: lib-clean
lib-clean:
	rm -rf protoc3 /tmp/protoc3 /tmp/pkg
	# rm -f protoc-${PROTOC_VER}-linux-x86_64.zip
	go clean

.PHONY: lib-install
lib-install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO}
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC}
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v${PROTOC_GEN_GATEWAY} 
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v${PROTOC_GEN_GATEWAY} 

.PHONY: lib-download
lib-download:
	# Protoc download
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip --output /tmp/protoc-${PROTOC_VER}-linux-x86_64.zip
	unzip /tmp/protoc-${PROTOC_VER}-linux-x86_64.zip  -d /tmp/protoc3
	mkdir protoc3
	mv /tmp/protoc3/include/* protoc3
	# rm -f protoc-${PROTOC_VER}-linux-x86_64.zip

	# googleapis for gateway
	curl -OL https://github.com/googleapis/googleapis/archive/master.tar.gz --output /tmp/googleapis.tar.gz
	tar xvf /tmp/googleapis.tar.gz --one-top-level=/tmp/gateway
	mkdir protoc3/google/api -p
	mv /tmp/gateway/googleapis-master/google/api/annotations.proto proto/protoc3/google/api/.
	mv /tmp/gateway/googleapis-master/google/api/field_behavior.proto proto/protoc3/google/api/.
	mv /tmp/gateway/googleapis-master/google/api/http.proto proto/protoc3/google/api/.
	mv /tmp/gateway/googleapis-master/google/api/httpbody.proto proto/protoc3/google/api/.
	

.PHONY: setup
setup: lib-clean lib-install 
	make lib-download

.PHONY: gen-v1-protos
gen-v1-proto:
	hacks/scripts/generate_proto.sh v1 ${TEMP_DIR}
	go mod tidy
	go mod vendor
	 
.PHONY: clean-proto
clean-v1-proto: 
	find ./pkg/api/v1/ -name "*.go"  |xargs rm -rf
	rm -rf /tmp/pkg

.PHONY: run
run:
	go run ./cmd/server/main.go --mode=$(MODE) --log-level=$(LOGLEVEL)
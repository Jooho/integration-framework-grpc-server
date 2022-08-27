# Include various params
include $(shell pwd)/env

RECREATE?=""
.PHONY: build-kfp-image
build-kfp-env-image:
	cd ./hack; ./build-kfp-image.sh

.PHONY: build-component-image
build-component-image:
	cd ${component_path}; ./build-image.sh

.PHONY: push-component-image
push-component-image:
	podman push ${full_image_name}
	podman tag ${full_image_name} quay.io/jooholee/rhods-integration-framework-deploy:latest
	podman push quay.io/jooholee/rhods-integration-framework-deploy:latest

.PHONY: run-deploy-model-img
run-deploy-model-img:
	cd ./hack; ./run-deploy-model-img.sh ${RECREATE}

.PHONY: run-kfp-env-img
run-kfp-env-img:
	cd ./hack; ./run-kfp-env-img.sh ${RECREATE}
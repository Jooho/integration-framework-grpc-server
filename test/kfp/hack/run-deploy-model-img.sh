source ../env.sh

# Test deploy-model.py in quay.io/jooholee/rhods-integration-framework-deploy:latest
if [[ $1 != "" ]] 
then 
  echo "Overwrite is set: delete the tests folder"
  rm -rf ../$component_path/tests/kfp-test
fi

if [[ ! -d ../$component_path/tests/kfp-test ]]
then
  echo "Create the tests folder"  
  mkdir ../$component_path/tests/kfp-test
  cp ../$component_path/src/deploy-model.py ../$component_path/tests/kfp-test/.
  cp ../$component_path/component.yaml ../$component_path/tests/kfp-test/.
  cp ../kfp-deploy.py ../$component_path/tests/kfp-test/.
else
  echo "Overwrite is not set or test folder exist"
fi

podman run --net=host --privileged -v ../$component_path/tests/kfp-test:/home -it $full_image_name /bin/bash
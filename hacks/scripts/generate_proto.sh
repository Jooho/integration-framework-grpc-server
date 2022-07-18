#! /bin/bash
version=$1
tmp_dir=$2
dest_dir=$3

export pb_dir=""

if [[ z${version} == z ]] 
then
    echo "version is required: v1 or v2"
    echo "usage: $0 v1 /tmp/proto  ./pkg/api"
    exit 1
fi

if [[ z${tmp_dir} ==  z ]]
then 
    tmp_dir=/tmp
fi

if [[ z${dest_dir} == z ]]
then    
    dest_dir=./pkg/api
fi

pb_dir=/tmp/proto

find ./proto/v1 -name "*.proto" | xargs -I % sh -c "protoc --proto_path=. --proto_path=proto/protoc3 --proto_path=./vendor  --go_out ${tmp_dir} --go_opt paths=source_relative --go-grpc_out ${tmp_dir} --go-grpc_opt paths=source_relative %"

for file in $(find ${pb_dir}/${version} -type f -name "*.go")
do
  export file_name=$(echo $file|awk -F/ '{print $NF}')
  export file_name_without_ext=$(echo $file|awk -F/ '{print $NF}'|cut -d. -f1)
  export dir_name=${file_name_without_ext}

  if [[ $file_name_without_ext =~ "grpc" ]]
  then
    dir_name=$(echo ${file_name_without_ext}|cut -d_ -f1)
  fi
  mkdir -p ${dest_dir}/${version}/${dir_name}
  cp ${pb_dir}/${version}/${file_name} ${dest_dir}/${version}/${dir_name}
done

source ../../../env.sh

cd "$(dirname "$0")"

echo "$full_image_name" .
podman build -t "$full_image_name" .

#Output the strict image name (which contains the sha256 image digest)
#This name can be used by the subsequent steps to refer to the exact image that was built even if another image with the same name was pushed.
image_name_with_digest=$(podman inspect --format="{{index .RepoDigests 0}}" "$full_image_name")
strict_image_name_output_file=./versions/image_digests_for_tags/$image_tag
mkdir -p "$(dirname "$strict_image_name_output_file")"
echo $image_name_with_digest | tee "$strict_image_name_output_file"
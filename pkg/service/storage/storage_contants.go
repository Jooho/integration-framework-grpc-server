package storage

var(
	StorageTypes= []string{}
	S3 = []string{"AWS","MINIO","CEPH"}
	GCP = "GCP"
	AZURE ="AZURE"
	PVC = "PVC"
)

func init(){
	StorageTypes = append(StorageTypes, S3...)
	StorageTypes = append(StorageTypes, GCP ,AZURE, PVC)
}
package storage

import (
	"context"

	v1storage "github.com/Jooho/integration-framework-server/pkg/api/v1/storage"
	"github.com/Jooho/integration-framework-server/pkg/logger"

	"google.golang.org/grpc"

	// appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

type storageServer struct {
	v1storage.StorageServer
	clientset *kubernetes.Clientset
}

func NewStorageServer(s grpc.Server, c *kubernetes.Clientset) {
	logger.Log.Info("Adding Storage Service to gRPC server...")

	storage := &storageServer{}
	storage.clientset = c

	v1storage.RegisterStorageServer(&s, storage)
}

func (s *storageServer) GetStorageParams(ctx context.Context, req *v1storage.GetStorageParamsRequest) (*v1storage.GetStorageParamResponse, error) {
	logger.Log.Debug("Entry storage.go - GetApplications")

	
	return &v1storage.GetStorageParamResponse{}, nil
}

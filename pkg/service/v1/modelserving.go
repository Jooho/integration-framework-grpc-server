package v1

import (
	"context"

	v1ms "github.com/Jooho/integration-framework-server/pkg/api/v1/modelserving"
	"github.com/Jooho/integration-framework-server/pkg/logger"
	"google.golang.org/grpc"
)

type modelServingServer struct {
	v1ms.ModelServingServer
}

func NewModelServingServer(s grpc.Server) {
	v1ms.RegisterModelServingServer(&s, &modelServingServer{})
	logger.Log.Info("Adding Model Serving Service to gRPC server...")
}

func (s *modelServingServer) GetApplications(ctx context.Context, req *v1ms.GetAppRequest) (*v1ms.GetAppResponse, error) {
	logger.Log.Debug("Entry modelserving.go - GetApplications")

	


	return &v1ms.GetAppResponse{}, nil
}

func (s *modelServingServer) GetAppParams(ctx context.Context, req *v1ms.GetAppParamsRequest) (*v1ms.GetAppParamsResponse, error) {
	return &v1ms.GetAppParamsResponse{}, nil
}

func (s *modelServingServer) GetAppCustomResource(ctx context.Context, req *v1ms.GetAppCRRequest) (*v1ms.GetAppCRResponse, error) {
	return &v1ms.GetAppCRResponse{}, nil
}

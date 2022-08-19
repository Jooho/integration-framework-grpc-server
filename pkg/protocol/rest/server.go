package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jooho/integration-framework-server/pkg/logger"
	"github.com/Jooho/integration-framework-server/pkg/protocol/rest/middleware"
	userv1 "github.com/Jooho/integration-framework-server/pkg/api/user/v1"
	storagev1 "github.com/Jooho/integration-framework-server/pkg/api/storage/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	// "github.com/amsokol/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	// "github.com/amsokol/go-grpc-http-rest-microservice-tutorial/pkg/logger"
	// "github.com/amsokol/go-grpc-http-rest-microservice-tutorial/pkg/protocol/rest/middleware"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := userv1.RegisterUserHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		logger.Log.Fatal("failed to add gRPC User Service to HTTP gateway", zap.String("reason", err.Error()))
	}
	if err := storagev1.RegisterStorageHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		logger.Log.Fatal("failed to add gRPC Storage Service to HTTP gateway", zap.String("reason", err.Error()))
	}
	srv := &http.Server{
		Addr: ":" + httpPort,
		// add handler with middleware
		Handler: middleware.AddRequestID(
			middleware.AddLogger(logger.Log, mux)),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
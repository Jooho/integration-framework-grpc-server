package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/Jooho/integration-framework-server/pkg/apis/user"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const portNumber = "9000"

func customMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		log.Print("Requested at:", time.Now())

		resp, err := handler(ctx, req)
		return resp, err
	}
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

  
	// grpcServer := grpc.NewServer()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			customMiddleware(),
			
			)),
		)

	reflection.Register(grpcServer)

	user.NewUserServer(*grpcServer)

	

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

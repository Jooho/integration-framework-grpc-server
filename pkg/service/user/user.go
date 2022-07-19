package user

import (
	"context"

	v1user "github.com/Jooho/integration-framework-server/pkg/api/v1/user"
	"github.com/Jooho/integration-framework-server/pkg/logger"
	"github.com/Jooho/integration-framework-server/test/data"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
)


type userServer struct {
	v1user.UserServer
	clientset *kubernetes.Clientset
}

func NewUserServer(s grpc.Server, c *kubernetes.Clientset) {
	logger.Log.Info("Adding User Service to gRPC server...")
	user := &userServer{}
	user.clientset=c
	
	v1user.RegisterUserServer(&s, user)
}

// GetUser returns user message by user_id
func (s *userServer) GetUser(ctx context.Context, req *v1user.GetUserRequest) (*v1user.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *v1user.UserMessage
	for _, u := range data.UserData {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}

	return &v1user.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}


// ListUsers returns all user messages
func (s *userServer) ListUsers(ctx context.Context, req *v1user.ListUsersRequest) (*v1user.ListUsersResponse, error) {
	userMessages := make([]*v1user.UserMessage, len(data.UserData))	
	for i, u := range data.UserData {
		userMessages[i] = u
	}
	return &v1user.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}
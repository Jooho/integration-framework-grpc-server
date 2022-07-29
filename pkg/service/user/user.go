package user

import (
	"context"

	userv1 "github.com/Jooho/integration-framework-server/pkg/api/user/v1"
	"github.com/Jooho/integration-framework-server/pkg/logger"
	"github.com/Jooho/integration-framework-server/test/data"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
)


type userServer struct {
	userv1.UserServer
	clientset *kubernetes.Clientset
}

func NewUserServer(s grpc.Server, c *kubernetes.Clientset) {
	logger.Log.Info("Adding User Service to gRPC server...")
	user := &userServer{}
	user.clientset=c
	
	userv1.RegisterUserServer(&s, user)
}

// GetUser returns user message by user_id
func (s *userServer) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *userv1.UserMessage
	for _, u := range data.UserData {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}

	return &userv1.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}


// ListUsers returns all user messages
func (s *userServer) ListUsers(ctx context.Context, req *userv1.ListUsersRequest) (*userv1.ListUsersResponse, error) {
	userMessages := make([]*userv1.UserMessage, len(data.UserData))	
	for i, u := range data.UserData {
		userMessages[i] = u
	}
	return &userv1.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}
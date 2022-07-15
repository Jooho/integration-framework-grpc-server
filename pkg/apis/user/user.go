package user

import (
	"context"

	v1pb "github.com/Jooho/integration-framework-server/pkg/api/v1"
	"github.com/Jooho/integration-framework-server/test/data"
	"google.golang.org/grpc"
)


type userServer struct {
	v1pb.UserServer
}

func NewUserServer(s grpc.Server) {
	v1pb.RegisterUserServer(&s, &userServer{})
}

// GetUser returns user message by user_id
func (s *userServer) GetUser(ctx context.Context, req *v1pb.GetUserRequest) (*v1pb.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *v1pb.UserMessage
	for _, u := range data.UserData {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}

	return &v1pb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}


// ListUsers returns all user messages
func (s *userServer) ListUsers(ctx context.Context, req *v1pb.ListUsersRequest) (*v1pb.ListUsersResponse, error) {
	userMessages := make([]*v1pb.UserMessage, len(data.UserData))	
	for i, u := range data.UserData {
		userMessages[i] = u
	}
	return &v1pb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}
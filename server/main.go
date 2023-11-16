package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"github.com/diogo-esteves/go_grpc/pb/user"
)

type userServiceServer struct{}

func (s *userServiceServer) GetUserById(ctx context.Context, req *user.UserRequest) (*user.User, error) {
	// Logic to fetch user details from the database
	user := &user.User{
		Id:   req.UserId,
		Name: "Diogo Esteves", // Assuming a static name for simplicity
	}
	return user, nil
}

func main() {
	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &userServiceServer{})

	fmt.Println("gRPC server is running on port 5001")
	if err := server.Serve(listen); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

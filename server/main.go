package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/diogo-esteves/go_grpc/pb"

	"google.golang.org/grpc"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer // Embed the UnimplementedServiceServer
}

func (s *userServiceServer) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	// Logic to fetch user details from the database
	user := &pb.User{
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
	pb.RegisterUserServiceServer(server, &userServiceServer{})

	fmt.Println("gRPC server is running on port 5001")
	if err := server.Serve(listen); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

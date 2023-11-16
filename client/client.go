
package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "github.com/diogo-esteves/go_grpc/pb"
)

func main() {
	connection, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	userDetails, err := client.GetUserById(context.Background(), &pb.UserRequest{UserId: 1})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("User Details: %+v\n", userDetails)
}

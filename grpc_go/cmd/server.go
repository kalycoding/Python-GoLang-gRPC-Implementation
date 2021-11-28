package main

import (
	"context"
	pb "example/grpc/grpc_go"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement message.CreateUser.
type server struct {
	pb.UnimplementedCreateUserServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCreateUserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) Create(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("User Created %v", in)
	msg := pb.MsgToUserResponse(in)
	return msg, nil
}

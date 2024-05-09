package main

import (
	"context"
	"log"
	"net"

	pb "github.com/knut1027/go-gRPC"
	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name + ", agein"}, nil
}

func (s *server) SayHelloWithID(ctx context.Context, in *pb.HelloWithIDRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Id)
	return &pb.HelloReply{Message: "Hello " + in.Id}, nil
}

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

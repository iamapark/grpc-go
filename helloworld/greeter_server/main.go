package main

import (
	"golang.org/x/net/context"
	pb "github.com/iamapark/grpc-go/helloworld/helloworld"
	"net"
	"log"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error)  {
	log.Printf("HelloRequest: %v", in)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	log.Println("listen port", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

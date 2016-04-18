package main

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/iamapark/grpc-go/helloworld/helloworld"
	"os"
	"golang.org/x/net/context"
)

const (
	address = "localhost:50051"
	defaultName = "JYP"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

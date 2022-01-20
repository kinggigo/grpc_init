package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc_test/greet/greetpb"
	"log"
	"net"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("hello this is server %v", req.Greeting.FirstName)
	firstname := req.GetGreeting().GetFirstName()

	result := "hello " + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("hello world")

	a, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Fail to listen : %v\n", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(a); err != nil {
		log.Fatalf("fail to server : %v", err)
	}
}

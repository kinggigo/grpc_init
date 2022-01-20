package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_test/greet/greetpb"
	"log"
)

func main() {
	fmt.Println("hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect : %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)
	//req := &greetpb.GreetRequest{
	//	Greeting: &greetpb.Greeting{
	//		FirstName: "dong",
	//		LastName:  "ha",
	//	},
	//}
	//res, err := c.Greet(context.Background(), req)
	//
	//if err != nil {
	//	log.Fatalf("fail call greet %v", err)
	//}
	//log.Printf("response from Greet , %v", res)
	//fmt.Printf("Create clent %f", c)
}

func doUnary(c greetpb.GreetServiceClient) {

	log.Println("Starting Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "dong",
			LastName:  "ha",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("fail call greet %v", err)
	}
	log.Printf("response from Greet , %v", res)
	//fmt.Printf("Create clent %f", c)
}

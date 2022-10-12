package main

import (
	"context"
	"fmt"
	"hello/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect %+v", err)
	}

	defer cc.Close()

	c := proto.NewHelloServiceClient(cc)

	helloUnary(c)
}

func helloUnary(c proto.HelloServiceClient) {
	fmt.Println("Starting unary RPC Hello")

	req := &proto.HelloRequest{
		Hello: &proto.Hello{
			FirstName: "Christian",
			Prefix:    "Joven",
		},
	}

	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error, calling HEllo RPC: %+v", err)
	}

	log.Printf("Response Hello %+v \n", res.CustomHello)
}

package main

import (
	"context"
	"fmt"
	"hello/proto"
	"io"
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

	// helloUnary(c)
	helloServerStream(c)
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
		log.Fatalf("Error, calling Hello RPC: %+v", err)
	}

	log.Printf("Response Hello %+v \n", res.CustomHello)
}

func helloServerStream(c proto.HelloServiceClient) {
	fmt.Printf("Starting server streaming RPC Hello \n")

	req := &proto.HelloManyLanguagesRequest{
		Hello: &proto.Hello{
			FirstName: "Christian",
			Prefix:    "Developer",
		},
	}

	restSream, err := c.HelloManyLanguages(context.Background(), req)
	if err != nil {
		log.Printf("Error calling the stream %+v \n", err)
	}

	for {
		msg, err := restSream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("stream error %+v \n", err)
		}

		log.Printf("Res from HML: %+v\n", msg.GetHelloLanguage())
	}
}

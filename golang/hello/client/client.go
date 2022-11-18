package main

import (
	"context"
	"fmt"
	"hello/proto"
	"io"
	"log"
	"time"

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
	// helloServerStream(c)
	// goodbyeClientStreaming(c)
	goodbyeBIDI(c)
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

func goodbyeClientStreaming(c proto.HelloServiceClient) {
	fmt.Println("Starting goodbye function")

	requests := []*proto.HelloGoodbyeRequest{
		{Hello: &proto.Hello{FirstName: "Christian", Prefix: "Sr."}},
		{Hello: &proto.Hello{FirstName: "Andrea", Prefix: "Ms."}},
		{Hello: &proto.Hello{FirstName: "Andoni", Prefix: "Sr."}},
		{Hello: &proto.Hello{FirstName: "Iker", Prefix: "Sr."}},
	}

	stream, err := c.HelloGoodbye(context.Background())

	if err != nil {
		log.Printf("Error calling goodbye %+v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %+v \n", req)

		stream.Send(req)
		time.Sleep(time.Second)
	}

	goodbye, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error goodbye receive %+v", err)
	}

	fmt.Println(goodbye)
}

func goodbyeBIDI(c proto.HelloServiceClient) {
	fmt.Println("Starting goodbye BIDI function")

	//create stream to call server
	stream, err := c.Goodbye(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream %+v", err)
	}
	requests := []*proto.GoodbyeRequest{
		{Hello: &proto.Hello{FirstName: "Christian", Prefix: "Sr."}},
		{Hello: &proto.Hello{FirstName: "Andrea", Prefix: "Ms."}},
		{Hello: &proto.Hello{FirstName: "Andoni", Prefix: "Sr."}},
		{Hello: &proto.Hello{FirstName: "Iker", Prefix: "Sr."}},
	}

	waitc := make(chan struct{})

	//send many messages to the server (go routines)
	go func() {
		for _, req := range requests {
			log.Printf("Sending message %+v \n", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}

		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("Error closing stream %+v", err)
		}
	}()

	//Receive messages from the server (go routines)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving stream %+v", err)
			}

			fmt.Printf(" got it : %+v \n", res.GetGoodbye())
			close(waitc)
		}
	}()

	//Block when everything is completed
	<-waitc
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"hello/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Printf("Hello function was called with %+v \n", req)
	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	customHello := "Welcome! " + prefix + " " + firstName

	res := &proto.HelloResponse{
		CustomHello: customHello,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello, Go server is running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	serv := grpc.NewServer()

	proto.RegisterHelloServiceServer(serv, &server{})

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %+v", err)
	}

}

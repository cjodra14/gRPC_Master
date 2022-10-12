package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

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

func (*server) HelloManyLanguages(req *proto.HelloManyLanguagesRequest, stream proto.HelloService_HelloManyLanguagesServer) error {
	fmt.Printf("Hello many times function was invoked with %+v \n", req)

	langs := []string{"Salut", "Hello", "Ni hao", "Kaixo", "Alô", "Privyét", "Schalom", "Hola", "Yassou", "Hej"}

	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	for _, greeting := range langs {
		helloLanguage := greeting + " " + prefix + " " + firstName

		res := &proto.HelloManyLanguagesResponse{
			HelloLanguage: helloLanguage,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	return nil
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

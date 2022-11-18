package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) HelloGoodbye(stream proto.HelloService_HelloGoodbyeServer) error {
	fmt.Println("Goodbye function was invoked")

	goodbye := "Goodbye guys:"

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//Once is finished the strwam we gonna send the response

			return stream.SendAndClose(&proto.HelloGoodbyeResponse{
				Goodbye: goodbye,
			})
		}

		if err != nil {
			log.Fatalf("Error reading the client stream %+v", err)
		}

		firstName := req.GetHello().GetFirstName()
		prefix := req.GetHello().GetPrefix()

		goodbye += prefix + " " + firstName + " "
	}
}

func (*server) Goodbye(stream proto.HelloService_GoodbyeServer) error{
	fmt.Printf("Goodabye bidirectional function was invoked \n")
	
	for{
		req, err := stream.Recv()
		if err != nil{
			log.Fatalf("Error reading the client stream %+v", err)
		}

		if err == io.EOF{
			return nil
		}

		firstName := req.Hello.FirstName

		prefix := req.Hello.Prefix

		goodbye := "Goodbye "+prefix+" "+firstName+" :("

		sendErr:= stream.Send(&proto.GoodbyeResponse{
			Goodbye: goodbye,
		})
		if sendErr != nil{
			log.Fatalf("Error sending to the client %+v", err)
		}
	}
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

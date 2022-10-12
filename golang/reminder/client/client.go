package main

import (
	"context"
	"fmt"
	"log"
	"reminder/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect %+v", err)
	}

	defer cc.Close()

	c := proto.NewReminderServiceClient(cc)

	setRemindUnary(c)
	getRemindUnary(c)
}

func setRemindUnary(c proto.ReminderServiceClient) {
	fmt.Println("Starting unary RPC Reminder")

	req := &proto.SetReminderRequest{
		Reminder: &proto.Reminder{
			Id:         12,
			Message:    "reminder",
			Deadline:   timestamppb.Now(),
			Tags:       []string{"gRPC", "API"},
			TeamMember: "Christian",
			Priority:   1,
		},
	}

	res, err := c.SetRemind(context.Background(), req)
	if err != nil {
		log.Fatalf("Error, calling Reminder RPC: %+v", err)
	}

	log.Printf("Response Reminder %+v \n", res.Id)
}

func getRemindUnary(c proto.ReminderServiceClient) {
	fmt.Println("Starting unary RPC Reminder")

	req := &proto.ReminderRequest{
		Id: 12,
	}

	res, err := c.GetRemindByID(context.Background(), req)
	if err != nil {
		log.Fatalf("Error, calling Reminder RPC: %+v", err)
	}

	log.Printf("Response Reminder %+v \n", res.Reminder)
}

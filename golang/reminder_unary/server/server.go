package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"reminder/proto"

	"google.golang.org/grpc"
)

type server struct {
	Reminds map[int64]proto.Reminder
}

func (s *server) GetRemindByID(ctx context.Context, req *proto.ReminderRequest) (*proto.ReminderResponse, error) {
	fmt.Printf("Reminder function was called with %+v \n", req)

	reminderID := req.GetId()

	reminder := s.Reminds[reminderID]

	res := &proto.ReminderResponse{
		Reminder: &proto.Reminder{
			Id:         reminder.Id,
			Message:    reminder.Message,
			Deadline:   reminder.Deadline,
			Tags:       reminder.Tags,
			TeamMember: reminder.TeamMember,
			Priority:   reminder.Priority,
		},
	}

	return res, nil
}

func (s *server) SetRemind(ctx context.Context, req *proto.SetReminderRequest) (*proto.SetReminderResponse, error) {
	if s.Reminds == nil {
		s.Reminds = make(map[int64]proto.Reminder)
	}
	fmt.Printf("Reminder function was called with %+v \n", req)

	reminder := req.GetReminder()

	s.Reminds[reminder.Id] = *reminder

	res := &proto.SetReminderResponse{
		Id: reminder.Id,
	}

	return res, nil
}

func main() {
	fmt.Println("Reminder, Go server is running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	serv := grpc.NewServer()

	proto.RegisterReminderServiceServer(serv, &server{})

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %+v", err)
	}

}

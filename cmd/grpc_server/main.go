package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/travacry/chat-server/pkg/chat_v1"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) RegistrationUser(_ context.Context, req *desc.RegistrationUserRequest) (*desc.RegistrationUserResponse, error) {

	fmt.Printf(color.RedString("Registration User:\n"),
		color.GreenString("info : %+v", req.GetUser()))

	return &desc.RegistrationUserResponse{Id: 10001}, nil
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	fmt.Printf(color.RedString("Create Chat:\n"),
		color.GreenString("info : %+v", req.GetUsers()))

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	fmt.Printf(color.RedString("Delete Chat:\n"),
		color.GreenString("info : %d", req.GetId()))

	return &empty.Empty{}, nil
}

func (s *server) AddUser(_ context.Context, req *desc.AddUserRequest) (*empty.Empty, error) {

	fmt.Printf(color.RedString("Add User:\n"),
		color.GreenString("info : %+v", req.GetUser()))

	return &empty.Empty{}, nil
}

func (s *server) Ban(_ context.Context, req *desc.BanRequest) (*empty.Empty, error) {

	fmt.Printf(color.RedString("Ban User:\n"),
		color.GreenString("info : %+v", req.GetUser()))

	return &empty.Empty{}, nil
}

func (s *server) Confirm(_ context.Context, req *desc.ConfirmRequest) (*empty.Empty, error) {

	fmt.Printf(color.RedString("Ban User:\n"),
		color.GreenString("info : %d", req.GetId()))

	return &empty.Empty{}, nil
}

func (s *server) Connect(_ context.Context, req *desc.ConnectRequest) (*empty.Empty, error) {

	fmt.Printf(color.RedString("Ban User:\n"),
		color.GreenString("info : %d", req.GetId()))

	return &empty.Empty{}, nil
}

func (s *server) Send(_ context.Context, req *desc.SendRequest) (*empty.Empty, error) {

	fmt.Printf(color.RedString("Send Text:\n"),
		color.GreenString("from : %s, msg : %s", req.From, req.Text))

	return &empty.Empty{}, nil
}

func (s *server) List(_ context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {

	fmt.Printf(color.RedString("Send Text:\n"),
		color.GreenString("info : %d", req.GetId()))

	return &desc.ListResponse{
		Chats: []*desc.ChatInfo{
			{Id: gofakeit.Int64(), Name: gofakeit.Name(), CreateAt: timestamppb.New(gofakeit.Date())},
			{Id: gofakeit.Int64(), Name: gofakeit.Name(), CreateAt: timestamppb.New(gofakeit.Date())},
		},
	}, nil
}

func (s *server) GetInfo(_ context.Context, req *desc.GetInfoRequest) (*desc.GetInfoResponse, error) {

	fmt.Printf(color.RedString("Get Info:\n"),
		color.GreenString("info : %d", req.GetId()))

	return &desc.GetInfoResponse{Chat: &desc.ChatInfo{
		Id:       gofakeit.Int64(),
		Name:     gofakeit.Name(),
		CreateAt: timestamppb.New(gofakeit.Date()),
	}}, nil
}

func (s *server) ListUsers(_ context.Context, req *desc.ListUsersRequest) (*desc.ListUsersResponse, error) {

	fmt.Printf(color.RedString("List Users:\n"),
		color.GreenString("info : %+d", req.GetId()))

	return &desc.ListUsersResponse{Users: []*desc.UserInfo{
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))

	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server chat listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Printf("failed to chat serve: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/travacry/chat-server/pkg/chat_v1"
)

const (
	grpcPort = 50052
)

type server struct {
	desc.UnimplementedChatV1Server
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

func (s *server) Connect(_ context.Context, req *desc.ConnectRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Ban User: "))
	fmt.Print(color.GreenString("info: %d\n", req.GetId()))

	return &empty.Empty{}, nil
}

func (s *server) Send(_ context.Context, req *desc.SendRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Send: "))
	fmt.Print(color.GreenString("from: %s, msg : %s\n", strconv.FormatInt(req.From, 10), req.Text))

	return &empty.Empty{}, nil
}

func (s *server) List(_ context.Context, _ *desc.ListRequest) (*desc.ListResponse, error) {
	fmt.Print(color.RedString("ListResponse.\n"))

	return &desc.ListResponse{
		Chats: []*desc.ChatModel{
			{Id: gofakeit.Int64(), Chat: &desc.ChatInfo{Name: gofakeit.Name(), CreateAt: timestamppb.New(gofakeit.Date())}},
			{Id: gofakeit.Int64(), Chat: &desc.ChatInfo{Name: gofakeit.Name(), CreateAt: timestamppb.New(gofakeit.Date())}},
			{Id: gofakeit.Int64(), Chat: &desc.ChatInfo{Name: gofakeit.Name(), CreateAt: timestamppb.New(gofakeit.Date())}},
			{Id: gofakeit.Int64(), Chat: &desc.ChatInfo{Name: gofakeit.Name(), CreateAt: timestamppb.New(gofakeit.Date())}},
		},
	}, nil
}

func (s *server) ListUsers(_ context.Context, req *desc.ListUsersRequest) (*desc.ListUsersResponse, error) {
	fmt.Print(color.RedString("UserInfo: "))
	fmt.Print(color.GreenString("%+d\n", req.GetId()))

	return &desc.ListUsersResponse{Users: []*desc.UserModel{
		{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}},
		{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}},
		{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}},
		{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}},
		{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}},
		{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}},
		{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}},
		{Id: gofakeit.Int64(), User: &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}},
	}}, nil
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Print(color.RedString("CreateResponse: "))
	fmt.Print(color.GreenString("%+v\n", req.GetUsers()))

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Delete Chat: "))
	fmt.Print(color.GreenString("%d\n", req.GetId()))

	return &empty.Empty{}, nil
}

func (s *server) AddUser(_ context.Context, req *desc.AddUserRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Add User: "))
	fmt.Print(color.GreenString("info : %+v\n", req.GetUser()))

	return &empty.Empty{}, nil
}

func (s *server) Ban(_ context.Context, req *desc.BanRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Ban User: "))
	fmt.Print(color.GreenString("%+d\n", req.GetId()))

	return &empty.Empty{}, nil
}

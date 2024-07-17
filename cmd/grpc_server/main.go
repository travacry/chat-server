package main

import (
	"context"
	"fmt"
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

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	fmt.Printf(color.RedString("Create Chat:\n"),
		color.GreenString("info : %+v", req.GetUsers()))

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Send(_ context.Context, req *desc.SendRequest) (*empty.Empty, error) {

	fmt.Printf(color.RedString("Send Text:\n"),
		color.GreenString("from : %s, msg : %s", req.From, req.Text))

	return &empty.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	fmt.Printf(color.RedString("Delete Chat:\n"),
		color.GreenString("info : %+v", req.GetId()))

	return &empty.Empty{}, nil
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

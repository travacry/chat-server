package main

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	desc "github.com/travacry/chat-server/pkg/chat_v1"
)

const (
	address = "localhost:50052"
	userID  = 100001
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createRequest(ctx, client)
	sendRequest(ctx, client)
	deleteRequest(ctx, client)
}

func createRequest(ctx context.Context, client desc.ChatV1Client) {

	users := []*desc.UserInfo{
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
	}

	createRequest, err := client.Create(ctx, &desc.CreateRequest{
		Users: users,
	})

	if err != nil {
		log.Fatalf("failed to create chat by id: %v", err)
	}
	log.Printf(color.RedString("Created chat :\n"),
		color.GreenString("%+d", createRequest.GetId()))
}

func deleteRequest(ctx context.Context, client desc.ChatV1Client) {

	_, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: userID,
	})

	if err != nil {
		log.Fatalf("failed to delete chat by id: %v", err)
	}
	log.Print(color.RedString("Delete chat :\n"))
}

func sendRequest(ctx context.Context, client desc.ChatV1Client) {

	_, err := client.Send(ctx, &desc.SendRequest{
		From: "name:name@mail",
		Text: "text text text",
	})

	if err != nil {
		log.Fatalf("failed send to chat : %v", err)
	}
	log.Print(color.RedString("Send to chat :\n"))
}

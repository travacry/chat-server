package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	desc "github.com/travacry/chat-server/pkg/chat_v1"
)

const (
	address = "localhost:50052"
	userID  = 100001
)

func main() {

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	_, err = createRequest(ctx, client)

	if err != nil {
		log.Print(err)
	}

	err = sendRequest(ctx, client)

	if err != nil {
		log.Print(err)
	}

	err = deleteRequest(ctx, client)

	if err != nil {
		log.Print(err)
	}
}

func createRequest(ctx context.Context, client desc.ChatV1Client) (*desc.CreateResponse, error) {

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
		return nil, fmt.Errorf("failed to create chat by id: %v", err)
	}

	fmt.Printf(color.RedString("Created chat :\n"), color.GreenString("%+d", createRequest.GetId()))
	return createRequest, nil
}

func deleteRequest(ctx context.Context, client desc.ChatV1Client) error {

	_, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: userID,
	})

	if err != nil {
		return fmt.Errorf("failed to delete chat by id: %v", err)
	}

	fmt.Print(color.RedString("Delete chat :\n"))
	return nil
}

func sendRequest(ctx context.Context, client desc.ChatV1Client) error {

	_, err := client.Send(ctx, &desc.SendRequest{
		From: "name:name@mail",
		Text: "text text text",
	})

	if err != nil {
		return fmt.Errorf("failed send to chat : %v", err)
	}

	fmt.Print(color.RedString("Send to chat :\n"))
	return nil
}

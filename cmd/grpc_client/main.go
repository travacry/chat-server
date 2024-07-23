package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
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
	chatID  = 101
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

	err = connect(ctx, client)

	if err != nil {
		log.Print(err)
	}

	err = send(ctx, client)

	if err != nil {
		log.Print(err)
	}

	_, err = create(ctx, client)

	if err != nil {
		log.Print(err)
	}

	_, err = list(ctx, client)

	if err != nil {
		log.Print(err)
	}

	_, err = listUsers(ctx, client)

	if err != nil {
		log.Print(err)
	}

	err = del(ctx, client)

	if err != nil {
		log.Print(err)
	}

	err = addUser(ctx, client)

	if err != nil {
		log.Print(err)
	}

	err = ban(ctx, client)
}

func connect(ctx context.Context, client desc.ChatV1Client) error {
	_, err := client.Connect(ctx, &desc.ConnectRequest{Id: chatID})

	if err != nil {
		return connectError(userID, chatID)
	}

	fmt.Print(color.RedString("Connect: "))
	fmt.Print(color.GreenString("%d\n", chatID))
	return nil
}
func connectError(userID int64, chatID int64) error {
	return fmt.Errorf("user : %d filed to connect by chat : %d", userID, chatID)
}

func send(ctx context.Context, client desc.ChatV1Client) error {
	uID := int64(userID)

	_, err := client.Send(ctx, &desc.SendRequest{
		From: uID,
		Text: "text text text",
	})

	if err != nil {
		return sendError(err)
	}

	fmt.Print(color.RedString("Send to client: "))
	fmt.Print(color.GreenString("%d\n", uID))

	return nil
}
func sendError(err error) error {
	return fmt.Errorf("failed send to chat : %v", err)
}

func list(ctx context.Context, client desc.ChatV1Client) (*desc.ListResponse, error) {
	listResponse, err := client.List(ctx, &desc.ListRequest{})

	if err != nil {
		return nil, listError(err)
	}

	fmt.Print(color.RedString("List: "))
	fmt.Print(color.GreenString("%v\n", listResponse.GetChats()))

	return listResponse, nil
}
func listError(err error) error {
	return fmt.Errorf("failed view list : %v", err)
}

func listUsers(ctx context.Context, client desc.ChatV1Client) (*desc.ListUsersResponse, error) {
	listUsersResponse, err := client.ListUsers(ctx, &desc.ListUsersRequest{Id: userID})

	if err != nil {
		return nil, listUsersError(err)
	}

	fmt.Print(color.RedString("List users: "))
	fmt.Print(color.GreenString("%v\n", listUsersResponse.GetUsers()))
	return listUsersResponse, nil
}
func listUsersError(err error) error {
	return fmt.Errorf("failed list users : %v", err)
}

func create(ctx context.Context, client desc.ChatV1Client) (*desc.CreateResponse, error) {
	users := []*desc.UserInfo{
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
		{Name: gofakeit.Name(), Email: gofakeit.Email()},
	}

	createResponse, err := client.Create(ctx, &desc.CreateRequest{
		Users: users,
	})

	if err != nil {
		return nil, createError(err)
	}

	fmt.Print(color.RedString("Create chat: "))
	fmt.Print(color.GreenString("%s\n", strconv.FormatInt(createResponse.GetId(), 10)))
	return createResponse, nil
}
func createError(err error) error {
	return fmt.Errorf("failed to create chat : %v", err)
}

func del(ctx context.Context, client desc.ChatV1Client) error {
	cID := int64(chatID)

	_, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: cID,
	})

	if err != nil {
		return delError(err)
	}

	fmt.Print(color.RedString("Delete chat: "))
	fmt.Print(color.GreenString(fmt.Sprintf("%d\n", cID)))
	return nil
}
func delError(err error) error {
	return fmt.Errorf("failed to delete chat : %v", err)
}

func addUser(ctx context.Context, client desc.ChatV1Client) error {
	user := &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}

	_, err := client.AddUser(ctx, &desc.AddUserRequest{
		User: user,
	})

	if err != nil {
		return addUserError(err)
	}

	fmt.Print(color.RedString("Add user: "))
	fmt.Print(color.GreenString(fmt.Sprintf("%v\n", user)))
	return nil
}
func addUserError(err error) error {
	return fmt.Errorf("failed to add user : %v", err)
}

func ban(ctx context.Context, client desc.ChatV1Client) error {
	user := &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email()}

	_, err := client.Ban(ctx, &desc.BanRequest{
		Id: userID,
	})

	if err != nil {
		return banError(err)
	}

	fmt.Print(color.RedString("Ban user: "))
	fmt.Print(color.GreenString(fmt.Sprintf("%v\n", user)))
	return nil
}
func banError(err error) error {
	return fmt.Errorf("failed to ban user : %v", err)
}

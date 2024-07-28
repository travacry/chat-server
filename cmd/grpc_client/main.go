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
	"google.golang.org/protobuf/types/known/timestamppb"

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

	_, err = createChat(ctx, client)
	if err != nil {
		log.Print(err)
	}

	_, err = listChats(ctx, client)
	if err != nil {
		log.Print(err)
	}

	err = connect(ctx, client)
	if err != nil {
		log.Print(err)
	}

	_, err = listUsers(ctx, client)
	if err != nil {
		log.Print(err)
	}

	err = send(ctx, client)
	if err != nil {
		log.Print(err)
	}

	err = delChat(ctx, client)
	if err != nil {
		log.Print(err)
	}

	err = addUser(ctx, client)
	if err != nil {
		log.Print(err)
	}

	err = banUser(ctx, client)
	if err != nil {
		log.Print(err)
	}
}

func createChat(ctx context.Context, client desc.ChatV1Client) (*desc.CreateChatResponse, error) {
	users := []*desc.UserInfo{
		{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_NOT_CONFIRM},
		{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_NOT_CONFIRM},
		{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_NOT_CONFIRM},
		{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_NOT_CONFIRM},
	}

	CreateChatResponse, err := client.CreateChat(ctx, &desc.CreateChatRequest{
		Chat: &desc.ChatInfo{
			Name:     gofakeit.Name(),
			State:    desc.ChatState_CHAT_ACTIVE,
			CreateAt: timestamppb.New(gofakeit.Date())}, Users: users,
	})
	if err != nil {
		return nil, createError(err)
	}

	fmt.Print(color.RedString("Create chat: "))
	fmt.Print(color.GreenString("%s\n", strconv.FormatInt(CreateChatResponse.GetId(), 10)))
	return CreateChatResponse, nil
}
func createError(err error) error {
	return fmt.Errorf("failed to create chat : %v", err)
}

func delChat(ctx context.Context, client desc.ChatV1Client) error {
	_, err := client.DeleteChat(ctx, &desc.DeleteChatRequest{
		Id: chatID,
	})
	if err != nil {
		return delError(err)
	}

	fmt.Print(color.RedString("Delete chat: "))
	fmt.Print(color.GreenString("%d\n", chatID))
	return nil
}
func delError(err error) error {
	return fmt.Errorf("failed to delete chat : %v", err)
}

func listChats(ctx context.Context, client desc.ChatV1Client) (*desc.ListChatsResponse, error) {
	ListChatsResponse, err := client.ListChats(ctx, &desc.ListChatsRequest{})
	if err != nil {
		return nil, listError(err)
	}

	fmt.Print(color.RedString("List chats: "))
	fmt.Print(color.GreenString("%v\n", ListChatsResponse.GetChats()))

	return ListChatsResponse, nil
}
func listError(err error) error {
	return fmt.Errorf("failed view list : %v", err)
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
	_, err := client.SendMessage(ctx, &desc.SendMessageRequest{
		Message: &desc.Message{
			From: userID,
			Text: "text text text",
		},
	})
	if err != nil {
		return sendError(err)
	}

	fmt.Print(color.RedString("Send message to client: "))
	fmt.Print(color.GreenString("%d\n", userID))

	return nil
}
func sendError(err error) error {
	return fmt.Errorf("failed send to chat : %v", err)
}

func addUser(ctx context.Context, client desc.ChatV1Client) error {
	user := &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_NOT_CONFIRM}

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

func banUser(ctx context.Context, client desc.ChatV1Client) error {
	user := &desc.UserInfo{Name: gofakeit.Name(), Email: gofakeit.Email(), State: desc.UserState_USER_ACTIVE}

	_, err := client.BanUser(ctx, &desc.BanUserRequest{
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

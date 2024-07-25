package ctr

import (
	"chat_cli/internal/app/models"
	"chat_cli/internal/app/services"
	"chat_cli/internal/app/utils/console"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

type ChatController struct {
	auth *services.AuthService
	s    *services.ChatService
}

func NewChatController(auth *services.AuthService, s *services.ChatService) *ChatController {
	return &ChatController{auth, s}
}

func (c *ChatController) Create(cCtx *cli.Context) error {
	title := cCtx.Args().Get(0)
	if title == "" {
		return errors.New("title (arg[0]) is required")
	}

	users := cCtx.Int64Slice("users")
	if len(users) == 0 {
		return errors.New("passed empty user list: %v")
	}

	id, err := c.s.Create(cCtx.Context, title, users)
	if err != nil {
		return err
	}

	fmt.Printf("chat %v: %v has been created with members: %v\n", id, title, users)
	return nil
}

func (c *ChatController) Connect(cCtx *cli.Context) error {
	author := c.auth.GetClaims().Name
	fmt.Printf("claims: %v\n", c.auth.GetClaims())
	fmt.Printf("name: %v\n", c.auth.GetClaims().Name)

	rawID := cCtx.Args().Get(0)
	chatID, err := strconv.Atoi(rawID)
	if err != nil {
		return fmt.Errorf("failed converting id (%v) to int: %v", rawID, err)
	}

	messages, err := c.s.Connect(cCtx.Context, int64(chatID))
	if err != nil {
		return err
	}

	fmt.Printf("Connected to the chat %v\n\n", chatID)
	go func() {
		for message := range messages {
			fmt.Printf("%v: %v\n", message.Author, message.Text)
		}
	}()

	for {
		text := console.ReadLine()
		fmt.Printf("\033[1A\033[K") // Removes last line

		if !strings.HasPrefix(text, "/") {
			c.s.Send(cCtx.Context, int64(chatID), models.ChatMessage{
				Author: author,
				Text:   text,
			})
		} else if text == "/q" || text == "/quit" {
			break
		} else {
			fmt.Printf("alert: unknown command: %v\n", text)
		}
	}

	return nil
}

func (c *ChatController) Delete(cCtx *cli.Context) error {
	rawID := cCtx.Args().Get(0)
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return fmt.Errorf("failed converting id (%v) to int: %v", rawID, err)
	}

	if err := c.s.Delete(cCtx.Context, int64(id)); err != nil {
		return err
	}

	fmt.Printf("chat %v has been deleted\n", id)
	return nil
}

func (c *ChatController) List(cCtx *cli.Context) error {
	chats, err := c.s.List(cCtx.Context)
	if err != nil {
		return err
	}

	if len(chats) == 0 {
		fmt.Println("Seems like no chats haven't been added yet")
		return nil
	}

	fmt.Println("Avaliable chats:")
	for _, c := range chats {
		fmt.Println(c.String())
	}

	return nil
}

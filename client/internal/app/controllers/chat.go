package ctr

import (
	"chat_cli/internal/app/gen/chat"
	"chat_cli/internal/app/models"
	"chat_cli/internal/app/services"
	"chat_cli/internal/app/utils/console"
	"chat_cli/internal/app/utils/styled"
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
		return styled.Errorf("title (arg[0]) is required")
	}

	users := cCtx.Int64Slice("users")
	if len(users) == 0 {
		return styled.Errorf("passed empty user list")
	}

	id, err := c.s.Create(cCtx.Context, title, users)
	if err != nil {
		return styled.Errorf(err.Error())
	}

	return styled.Successf("chat %v: %v has been created with members: %v\n\n", id, title, users)
}

func (c *ChatController) Connect(cCtx *cli.Context) error {
	claims := c.auth.GetClaims()

	rawID := cCtx.Args().Get(0)
	chatID, err := strconv.Atoi(rawID)
	if err != nil {
		return styled.Errorf("failed converting id (%v) to int: %v", rawID, err)
	}

	events, errch, err := c.s.Connect(cCtx.Context, int64(chatID))
	if err != nil {
		return styled.Errorf(err.Error())
	}

	styled.Successf("Connected to the chat %v\n\n", chatID)
	go func() {
		for event := range events {
			fmt.Printf("\033[0G\033[K")
			if event.Type == chat.EventType_Info {
				if event.UserID != int64(claims.UserID) {
					text := styled.New(event.Text).WithStyle(styled.Blue).Build()
					fmt.Println(text)
				}
			} else {
				authorColor := styled.ColorFromCode(event.ColorCode)
				author := styled.New(event.Author).WithStyle(authorColor).Build()
				fmt.Printf("%v: %v\n", author, event.Text)
			}
		}
	}()

	input := make(chan string, 1)
	go func() {
		for {
			text := console.ReadLine()
			fmt.Printf("\033[1A\033[K") // Removes last line
			input <- text
		}
	}()

	colorCode := styled.RndColorCode()
	for {
		select {
		case text := <-input:
			if !strings.HasPrefix(text, "/") {
				c.s.Send(cCtx.Context, int64(chatID), models.ChatEvent{
					UserID:    int64(claims.UserID),
					Author:    claims.Name,
					Text:      text,
					ColorCode: colorCode,

					Type: chat.EventType_Message,
				})
			} else if text == "/q" || text == "/quit" {
				fmt.Printf("\n\n")
				return nil
			} else {
				styled.Errorf("alert: unknown command: %v\n", text)
			}
		case err := <-errch:
			return err
		}
	}
}

func (c *ChatController) Delete(cCtx *cli.Context) error {
	rawID := cCtx.Args().Get(0)
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return styled.Errorf("failed converting id (%v) to int: %v", rawID, err)
	}

	if err := c.s.Delete(cCtx.Context, int64(id)); err != nil {
		return styled.Errorf(err.Error())
	}

	styled.Successf("chat %v has been deleted\n\n", id)
	return nil
}

func (c *ChatController) List(cCtx *cli.Context) error {
	chats, err := c.s.List(cCtx.Context)
	if err != nil {
		return styled.Errorf(err.Error())
	}

	if len(chats) == 0 {
		return styled.Infof("Seems like no chats haven't been added yet")
	}

	for _, c := range chats {
		fmt.Println(c.String())
	}
	fmt.Println()
	return nil
}

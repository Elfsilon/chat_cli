package app

import (
	ctr "chat_cli/internal/app/controllers"
	"chat_cli/internal/app/gen/auth"
	"chat_cli/internal/app/gen/chat"
	"chat_cli/internal/app/gen/user"
	intc "chat_cli/internal/app/interceptors"
	"chat_cli/internal/app/models"
	"chat_cli/internal/app/services"
	"chat_cli/internal/app/utils/console"
	"chat_cli/internal/app/utils/styled"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	config models.Config
}

func New() *App {
	return &App{}
}

func PromptString(message string) string {
	fmt.Print(message)
	return console.ReadLine()
}

func (a *App) GetClientConn(url string, opts ...grpc.DialOption) *grpc.ClientConn {
	defaultOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	c, err := grpc.NewClient(url, append(defaultOpts, opts...)...)
	if err != nil {
		styled.Fatalf(err.Error())
	}
	return c
}

func (a *App) PromptLogin(ctx context.Context, s *services.AuthService) error {
	name, password := PromptString("Name: "), PromptString("Password: ")
	return s.Login(ctx, name, password)
}

func (a *App) Authorize(ctx context.Context, s *services.AuthService) {
	if err := s.UpdateRefreshToken(ctx); err != nil {
		if errors.Is(err, services.ErrTokenExpiredOrEmpty) || errors.Is(err, services.ErrSessionNotFound) {
			for lerr := a.PromptLogin(ctx, s); lerr != nil; lerr = a.PromptLogin(ctx, s) {
				styled.Errorf(lerr.Error())
			}
		} else {
			styled.Fatalf(err.Error())
		}
	}
	if err := s.UpdateAccessToken(ctx); err != nil {
		styled.Fatalf(err.Error())
	}
	go s.RunAccessTokenUpdater(ctx)
}

func (a *App) Run() {
	a.LoadConfig()

	authClient := auth.NewAuthClient(a.GetClientConn(a.config.AuthServiceUrl))
	authService := services.NewAuthService(authClient, a.config.AccessTokenTTL, a.config.JwtSecret)

	if err := authService.Load(); err != nil {
		styled.Errorf("failed load auth state from file: %v", err)
	}
	a.Authorize(context.Background(), authService)
	authService.Save()

	tokenProvider := intc.NewTokenProvider(authService)
	secureOpts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(tokenProvider.UnaryInterceptor),
		grpc.WithStreamInterceptor(tokenProvider.StreamInterceptor),
	}
	userClient := user.NewUserServiceClient(a.GetClientConn(a.config.UserServiceUrl, secureOpts...))
	chatClient := chat.NewChatClient(a.GetClientConn(a.config.ChatServiceUrl, secureOpts...))

	chatService := services.NewChatService(chatClient)
	chatController := ctr.NewChatController(authService, chatService)

	userService := services.NewUserService(userClient)
	userController := ctr.NewUserController(userService)

	app := &cli.App{
		Name:  "chater",
		Usage: "",
		Commands: []*cli.Command{
			{
				Name:     "chat",
				Category: "Chat",
				Usage:    "",
				Subcommands: []*cli.Command{
					{
						Name:     "create",
						Aliases:  []string{"c"},
						Category: "Chat",
						Usage:    "",
						Flags: []cli.Flag{
							&cli.Int64SliceFlag{
								Name:    "users",
								Aliases: []string{"u"},
								Value:   cli.NewInt64Slice(),
								Usage:   "u",
							},
						},
						Action: chatController.Create,
					},
					{
						Name:     "delete",
						Aliases:  []string{"d"},
						Category: "Chat",
						Usage:    "",
						Action:   chatController.Delete,
					},
					{
						Name:     "connect",
						Category: "Chat",
						Usage:    "",
						Action:   chatController.Connect,
					},
					{
						Name:     "list",
						Aliases:  []string{"ls"},
						Category: "Chat",
						Usage:    "",
						Action:   chatController.List,
					},
				},
			},
			{
				Name:     "user",
				Category: "User",
				Usage:    "",
				Subcommands: []*cli.Command{
					{
						Name:     "create",
						Aliases:  []string{"c"},
						Category: "User",
						Usage:    "",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "role"},
						},
						Action: userController.Create,
					},
					{
						Name:     "get",
						Aliases:  []string{"g"},
						Category: "User",
						Usage:    "",
						Action:   userController.Get,
					},
					{
						Name:     "list",
						Aliases:  []string{"ls"},
						Category: "User",
						Usage:    "",
						Action:   userController.List,
					},
					{
						Name:     "update",
						Aliases:  []string{"u"},
						Category: "User",
						Usage:    "",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name"},
							&cli.StringFlag{Name: "email"},
						},
						Action: userController.Update,
					},
					{
						Name:     "delete",
						Aliases:  []string{"d"},
						Category: "User",
						Usage:    "",
						Action:   userController.Delete,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		styled.Fatalf(err.Error())
	}
}

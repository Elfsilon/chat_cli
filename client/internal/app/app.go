package app

import (
	"bufio"
	"chat_cli/internal/app/gen/auth"
	"chat_cli/internal/app/gen/user"
	intc "chat_cli/internal/app/interceptors"
	"chat_cli/internal/app/models"
	"chat_cli/internal/app/services"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	config models.Config
	log    *log.Logger
	auths  *services.AuthService
	authc  auth.AuthClient
	userc  user.UserServiceClient
}

func New() *App {
	return &App{}
}

func PromptString(message string) string {
	fmt.Println(message)
	text, r := "", bufio.NewReader(os.Stdin)
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		text = str
		break
	}
	return strings.ReplaceAll(text, "\n", "")
}

func (a *App) GetClientConn(url string, opts ...grpc.DialOption) *grpc.ClientConn {
	defaultOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	c, err := grpc.NewClient(url, append(defaultOpts, opts...)...)
	if err != nil {
		a.log.Fatal(err)
	}
	return c
}

func (a *App) PromptLogin(ctx context.Context) error {
	fmt.Println("-- Authorization --")
	name, password := PromptString("Name:"), PromptString("Password")
	return a.auths.Login(ctx, name, password)
}

func (a *App) Authorize(ctx context.Context) {
	if err := a.auths.UpdateRefreshToken(ctx); err != nil {
		if errors.Is(err, services.ErrTokenExpiredOrEmpty) || errors.Is(err, services.ErrSessionNotFound) {
			for err := err; err != nil; err = a.PromptLogin(ctx) {
				a.log.Println(err)
			}
		} else {
			a.log.Fatal(err)
		}
	}
	a.auths.UpdateAccessToken(ctx)
	go a.auths.RunAccessTokenUpdater(ctx)
}

func (a *App) Run1() {
	a.log = log.Default()
	a.LoadConfig()

	a.authc = auth.NewAuthClient(a.GetClientConn(a.config.AuthServiceUrl))
	tokenProvider := intc.NewTokenProvider(a.auths)

	secureOpts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(tokenProvider.UnaryInterceptor),
		grpc.WithStreamInterceptor(tokenProvider.StreamInterceptor),
	}

	a.userc = user.NewUserServiceClient(a.GetClientConn(a.config.UserServiceUrl, secureOpts...))
	// chatc = auth.NewAuthClient(a.GetClientConn(a.config.ChatServiceUrl, secureOpts...))

	auths := services.NewAuthService(a.authc, a.config.AccessTokenTTL)
	if err := auths.Load(); err != nil {
		a.log.Printf("failed load auth state from file: %v", err)
	}

	a.Authorize(context.Background())
	auths.Save()

	a.log.Println("Authorization completed")

	time.Sleep(time.Hour)
}

// ---------

func JoinChat(cCtx *cli.Context) error {
	fmt.Println("added task: ", cCtx.Args().First())
	return nil
}

func (a *App) Run() {
	app := &cli.App{
		Name:  "chacli",
		Usage: "",
		Commands: []*cli.Command{
			{
				Name:     "join",
				Aliases:  []string{"j"},
				Category: "Chat",
				Usage:    "",
				Action:   JoinChat,
			},
			{
				Name:     "create",
				Aliases:  []string{"c"},
				Category: "Chat",
				Usage:    "",
				Action:   JoinChat,
			},
			{
				Name:     "delete",
				Aliases:  []string{"d"},
				Category: "Chat",
				Usage:    "",
				Action:   JoinChat,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

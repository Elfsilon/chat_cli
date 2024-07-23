package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/internal/auth/controllers"
	"server/internal/auth/gen/ent"
	"server/internal/auth/gen/user"
	"server/internal/auth/models"
	"server/internal/auth/repos"
	"server/internal/auth/services"
	"server/internal/auth/utils"
	"server/pkg/constants"
	"server/pkg/gen/auth"
	intc "server/pkg/interceptors"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	cfg        models.Config
	log        *log.Logger
	db         *ent.Client
	userClient user.UserServiceClient
	tcpLis     *net.Listener
	grpcServer *grpc.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) ConnectDatabase() func() error {
	a.log.Println("Try to connect to the database")
	client, err := ent.Open("postgres", a.cfg.Db.ConnURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	a.log.Println("Create db schema")
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	a.db = client

	return client.Close
}

func (a *App) StartTCP() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", a.cfg.Server.Host, a.cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	a.tcpLis = &lis
}

func (a *App) ConnectUserClient() func() error {
	withoutTLS := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(constants.UserServiceAddr, withoutTLS)
	if err != nil {
		log.Fatalf("failed creating new grpc user client: %v", err)
	}

	a.userClient = user.NewUserServiceClient(conn)

	return conn.Close
}

func (a *App) SetupGrpcServer() {
	tokenManager := utils.NewTokenManager(a.cfg.Auth.AccessTokenTTL, a.cfg.Auth.RefreshTokenTTL, []byte(a.cfg.Auth.JWTSecret))
	tokenUpdater := utils.NewTokenUpdater(tokenManager, a.cfg.Auth.AccessTokenTTL)

	if err := tokenUpdater.UpdateToken(); err != nil {
		a.log.Fatal(err)
	}
	go tokenUpdater.Run()

	authorizer := utils.NewRequestAuthorizer(tokenManager)
	sessionRepo := repos.NewSessionRepo(a.db)
	authService := services.NewAuthService(a.cfg.Auth.ServiceSecret, a.userClient, sessionRepo, tokenManager, tokenUpdater, authorizer)
	authController := controllers.NewAuthController(authService)

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(intc.Log, recovery.UnaryServerInterceptor()),
	}
	server := grpc.NewServer(opts...)

	auth.RegisterAuthServer(server, authController)

	a.grpcServer = server
}

func (a *App) Run() {
	a.log = log.Default()

	a.log.Println("Loading config")
	a.LoadConfig()

	closeDB := a.ConnectDatabase()
	defer closeDB()

	a.StartTCP()
	a.ConnectUserClient()
	a.SetupGrpcServer()

	a.log.Printf("Server has been launched on %v\n", (*a.tcpLis).Addr())
	if err := a.grpcServer.Serve(*a.tcpLis); err != nil {
		log.Fatal(err)
	}
}

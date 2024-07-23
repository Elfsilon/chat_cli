package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"server/internal/chat/controllers"
	"server/internal/chat/gen/chat"
	"server/internal/chat/gen/ent"
	"server/internal/chat/models"
	"server/internal/chat/repos"
	"server/internal/chat/services"
	"server/pkg/constants"
	"server/pkg/gen/auth"
	intc "server/pkg/interceptors"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	cfg        models.Config
	log        *log.Logger
	db         *ent.Client
	nc         *nats.Conn
	authClient auth.AuthClient
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

func (a *App) ConnectNATS() func() {
	nc, err := nats.Connect(a.cfg.NatsUrl)
	if err != nil {
		log.Fatalf("failed connecting to NATS server: %v", err)
	}

	a.nc = nc

	return nc.Close
}

func (a *App) StartTCP() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", a.cfg.Server.Host, a.cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	a.tcpLis = &lis
}

func (a *App) ConnectAuthClient() func() error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(constants.AuthServiceAddr, opts...)
	if err != nil {
		log.Fatalf("failed creating new grpc auth client: %v", err)
	}

	a.authClient = auth.NewAuthClient(conn)
	return conn.Close
}

func (a *App) SetupGrpcServer() func() {
	chatRepo := repos.NewChatRepo(a.db)

	chatService := services.NewChatService(a.nc, chatRepo)
	disposeChatService, err := chatService.Init()
	if err != nil {
		log.Fatalf("failed initializing chat service: %v", err)
	}

	chatController := controllers.NewChatController(chatService)

	guard := intc.NewAuthGuard(a.authClient)

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(intc.Log, guard.UnaryInterceptor, recovery.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(guard.StreamInterceptor, recovery.StreamServerInterceptor()),
	}
	server := grpc.NewServer(opts...)
	chat.RegisterChatServer(server, chatController)

	a.grpcServer = server

	return func() {
		disposeChatService()
	}
}

func (a *App) Run() {
	a.log = log.Default()

	a.log.Println("Loading config")
	a.LoadConfig()

	closeDB := a.ConnectDatabase()
	defer closeDB()

	closeNATS := a.ConnectNATS()
	defer closeNATS()

	a.StartTCP()
	a.ConnectAuthClient()

	dispose := a.SetupGrpcServer()
	defer dispose()

	a.log.Printf("Server has been launched on %v\n", (*a.tcpLis).Addr())
	if err := a.grpcServer.Serve(*a.tcpLis); err != nil {
		log.Fatal(err)
	}
}

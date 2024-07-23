package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/internal/user/controllers"
	"server/internal/user/gen/ent"
	"server/internal/user/gen/user"
	"server/internal/user/models"
	"server/internal/user/repos"
	"server/internal/user/services"
	"server/pkg/constants"
	"server/pkg/gen/auth"
	intc "server/pkg/interceptors"
	"server/pkg/utils/role"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	cfg        models.Config
	log        *log.Logger
	db         *ent.Client
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

func (a *App) SetupGrpcServer() {
	userRepo := repos.NewUserRepo(a.db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userService.Create(context.Background(), "konnichiwa", "konnichiwa@test.gore", "konnichiwa", role.Admin)

	guard := intc.NewAuthGuard(a.authClient)

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(intc.Log, guard.UnaryInterceptor, recovery.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(guard.StreamInterceptor, recovery.StreamServerInterceptor()),
	}
	server := grpc.NewServer(opts...)
	user.RegisterUserServiceServer(server, userController)

	a.grpcServer = server
}

func (a *App) Run() {
	a.log = log.Default()

	a.log.Println("Loading config")
	a.LoadConfig()

	closeDB := a.ConnectDatabase()
	defer closeDB()

	a.StartTCP()
	a.ConnectAuthClient()
	a.SetupGrpcServer()

	a.log.Printf("Server has been launched on %v\n", (*a.tcpLis).Addr())
	if err := a.grpcServer.Serve(*a.tcpLis); err != nil {
		log.Fatal(err)
	}
}

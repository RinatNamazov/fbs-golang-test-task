package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"

	fibonacciGrpcDelivery "github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci/delivery/grpc"
	fibonacciHttpDelivery "github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci/delivery/http"
	fibonacciRedisRepository "github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci/repository/redis"
	fibonacciUsecase "github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci/usecase"
)

type Application struct {
	config     *config
	rdb        *redis.Client
	grpcServer *grpc.Server
	httpRouter *httprouter.Router
	httpServer *http.Server
}

func New(configPath string) (*Application, error) {
	config, err := readConfig(configPath)
	if err != nil {
		return nil, err
	}

	app := &Application{
		config:     config,
		grpcServer: grpc.NewServer(),
		httpRouter: httprouter.New(),
	}
	app.httpServer = &http.Server{
		Addr:    app.config.Server.HttpAddress,
		Handler: app.httpRouter,
	}

	app.connectToRedis()
	app.initHandlers()

	return app, nil
}

func (m *Application) Run() {
	go m.startGrpcServer()
	go m.startHttpServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	// Waiting.
	<-quit

	m.Shutdown()
}

func (m *Application) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	grpcStoped := make(chan struct{})
	go func() {
		m.grpcServer.GracefulStop()
		close(grpcStoped)
	}()

	if err := m.httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown: %v", err)
	}

	select {
	case <-ctx.Done():
		m.grpcServer.Stop()
		break
	case <-grpcStoped:
		break
	}
}

func (m *Application) connectToRedis() {
	m.rdb = redis.NewClient(&redis.Options{
		Addr:     m.config.Redis.Address,
		Password: m.config.Redis.Password,
		DB:       m.config.Redis.Database,
	})

	_, err := m.rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis ping error: %v", err)
	}
}

func (m *Application) initHandlers() {
	fibonacciUcase := fibonacciUsecase.NewUsecase(fibonacciRedisRepository.NewRepository(m.rdb))

	fibonacciGrpcDelivery.RegisterService(m.grpcServer, fibonacciUcase)
	fibonacciHttpDelivery.NewHandler(m.httpRouter, fibonacciUcase)
}

func (m *Application) startHttpServer() {
	if err := m.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to listen and serve http: %+v", err)
	}
}

func (m *Application) startGrpcServer() {
	ln, err := net.Listen("tcp", m.config.Server.GrpcAddress)
	if err != nil {
		log.Fatalf("Failed to listen grpc: %v", err)
	}

	if err := m.grpcServer.Serve(ln); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}
}

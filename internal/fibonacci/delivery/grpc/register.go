package grpc

import (
	"google.golang.org/grpc"

	"github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci"
	pb "github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci/delivery/grpc/proto"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/fibonacci.proto

type FibonacciService struct {
	pb.UnimplementedFibonacciServiceServer

	usecase fibonacci.UseCase
}

func RegisterService(server *grpc.Server, uc fibonacci.UseCase) {
	pb.RegisterFibonacciServiceServer(server, &FibonacciService{usecase: uc})
}

package grpc

import (
	"context"

	pb "github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci/delivery/grpc/proto"
)

func (m *FibonacciService) GetFibonacciSequence(ctx context.Context, r *pb.GetFibonacciSequenceRequest) (*pb.GetFibonacciSequenceResponse, error) {
	sequence, err := m.usecase.GetFibonacciSequence(ctx, r.From, r.To)
	if err != nil {
		return nil, err
	}

	return &pb.GetFibonacciSequenceResponse{
		Sequence: sequence,
	}, nil
}

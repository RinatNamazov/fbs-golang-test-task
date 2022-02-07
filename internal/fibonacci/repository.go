package fibonacci

import (
	"context"
)

type Repository interface {
	GetFibonacciSequenceLength(ctx context.Context) (int64, error)
	GetFibonacciSequence(ctx context.Context, from, to int64) ([]uint64, error)
	AddFibonacciSequence(ctx context.Context, sequence []uint64) error
}

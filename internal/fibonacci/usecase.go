package fibonacci

import (
	"context"
)

type UseCase interface {
	GetFibonacciSequence(ctx context.Context, from, to uint32) ([]uint64, error)
}

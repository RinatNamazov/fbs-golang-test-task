package usecase

import (
	"context"
	"sync"

	"github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci"
)

type UseCase struct {
	repo       fibonacci.Repository
	cacheMutex sync.Mutex
}

func NewUsecase(repo fibonacci.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (m *UseCase) GetFibonacciSequence(ctx context.Context, from, to uint32) ([]uint64, error) {
	if from > to {
		return nil, fibonacci.ErrBadIndex
	}

	needCount := int64(to-from) + 1
	if needCount <= 0 {
		return []uint64{}, nil
	}

	m.cacheMutex.Lock()
	defer m.cacheMutex.Unlock()

	cachedSequence, err := m.repo.GetFibonacciSequence(ctx, int64(from), int64(to))
	if err != nil {
		return nil, nil
	}

	cachedCount := len(cachedSequence)
	if int64(cachedCount) == needCount {
		return cachedSequence, nil
	}

	var a, b uint64
	var cacheLength int64
	if cachedCount > 2 {
		a, b = cachedSequence[cachedCount-2], cachedSequence[cachedCount-1]
		cacheLength = int64(cachedCount) + int64(from)
	} else {
		cacheLength, err = m.repo.GetFibonacciSequenceLength(ctx)
		if err != nil {
			return nil, err
		}

		if cacheLength >= 2 {
			lastSequence, err := m.repo.GetFibonacciSequence(ctx, -2, -1) // Last 2 values.
			if err != nil {
				return nil, nil
			}
			a, b = lastSequence[0], lastSequence[1]
		} else {
			a, b = 0, 1
		}
	}

	offset := int64(from) - cacheLength
	diffCount := uint32(needCount + offset)

	sequence := caclulateFibonacciSequence(a, b, diffCount)

	if cacheLength > 2 {
		m.repo.AddFibonacciSequence(ctx, sequence[2:])
	} else {
		m.repo.AddFibonacciSequence(ctx, sequence)
	}

	if cachedCount > 0 {
		return append(cachedSequence, sequence[2:]...), nil
	} else {
		if cacheLength > 0 {
			return sequence[offset+2:], nil
		} else {
			return sequence[offset : len(sequence)-2], nil
		}

	}
}

func caclulateFibonacciSequence(a, b uint64, count uint32) []uint64 {
	count += 2
	f := make([]uint64, count)
	f[0] = a
	f[1] = b
	for i := uint32(2); i < count; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

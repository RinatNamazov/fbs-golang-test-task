package redis

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const (
	cacheKey = "fibonacci_sequence"
)

type Repository struct {
	db *redis.Client
}

func NewRepository(db *redis.Client) *Repository {
	return &Repository{db: db}
}

func (m *Repository) GetFibonacciSequenceLength(ctx context.Context) (int64, error) {
	return m.db.LLen(ctx, cacheKey).Result()
}

func (m *Repository) GetFibonacciSequence(ctx context.Context, from, to int64) ([]uint64, error) {
	res, err := m.db.LRange(ctx, cacheKey, from, to).Result()
	if err != nil {
		return nil, err
	}
	return convertStringSliceToUint64Slice(res)
}

func (m *Repository) AddFibonacciSequence(ctx context.Context, sequence []uint64) error {
	return m.db.RPush(ctx, cacheKey, convertUint64SliceToInterfaceSlice(sequence)...).Err()
}

func convertUint64SliceToInterfaceSlice(s []uint64) []interface{} {
	ns := make([]interface{}, len(s))
	for i, v := range s {
		ns[i] = v
	}
	return ns
}

func convertStringSliceToUint64Slice(s []string) ([]uint64, error) {
	ns := make([]uint64, len(s))
	for k, v := range s {
		nv, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ns[k] = nv
	}
	return ns, nil
}

package pgsql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func intToInt32Array(in []int) []int32 {
	var out []int32
	for _, val := range in {
		out = append(out, int32(val))
	}
	return out
}

func NewStorage(conn string) (*Storage, error) {
	p, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	return &Storage{pool: p}, nil
}

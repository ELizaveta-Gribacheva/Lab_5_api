package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error)
	GetTask(ctx context.Context, id int32) (Task, error)
	ListTasks(ctx context.Context) ([]Task, error)
	UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error)
	DeleteTask(ctx context.Context, id int32) error
}

type SQLStore struct {
	*Queries
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) Store {
	return &SQLStore{
		Queries: New(pool),
		pool:    pool,
	}
}

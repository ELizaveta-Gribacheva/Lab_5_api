package server

import (
	"context"

	db "github.com/ELizaveta-Gribacheva/Lab_5_api/db/sqlc"
)

type MockStore struct{}

func (m *MockStore) CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
	return db.Task{ID: 101, Title: arg.Title, Description: arg.Description, Completed: arg.Completed}, nil
}

func (m *MockStore) GetTask(ctx context.Context, id int32) (db.Task, error) {
	return db.Task{ID: id, Title: "Sample", Description: "Testing", Completed: false}, nil
}

func (m *MockStore) ListTasks(ctx context.Context) ([]db.Task, error) {
	return []db.Task{{ID: 1, Title: "One"}, {ID: 2, Title: "Two"}}, nil
}

func (m *MockStore) UpdateTask(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error) {
	return db.Task(arg), nil
}

func (m *MockStore) DeleteTask(ctx context.Context, id int32) error {
	return nil
}

func newTestHandler() *TaskHandler {
	return NewTaskHandler(&MockStore{})
}

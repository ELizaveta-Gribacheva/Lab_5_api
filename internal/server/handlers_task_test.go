package server

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/ELizaveta-Gribacheva/Lab_5_api/db/sqlc"
	"github.com/gorilla/mux"
)

type MockStore struct{}

func (m *MockStore) CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
	return db.Task{
		ID:          101,
		Title:       arg.Title,
		Description: arg.Description,
		Completed:   arg.Completed,
	}, nil
}

func (m *MockStore) GetTask(ctx context.Context, id int32) (db.Task, error) {
	return db.Task{
		ID:          id,
		Title:       "Sample Task",
		Description: "Just testing",
		Completed:   false,
	}, nil
}

func (m *MockStore) ListTasks(ctx context.Context) ([]db.Task, error) {
	return []db.Task{
		{ID: 1, Title: "One", Description: "First", Completed: false},
		{ID: 2, Title: "Two", Description: "Second", Completed: true},
	}, nil
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

func TestCreateTask(t *testing.T) {
	h := newTestHandler()
	payload := `{"title":"Title","description":"Desc","completed":true}`

	req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(payload))
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)
	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201, got %v", rec.Code)
	}
}

func TestCreateTaskInvalid(t *testing.T) {
	h := newTestHandler()
	payload := `{"title":"","description":"","completed":true}`

	req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(payload))
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %v", rec.Code)
	}
}

func TestGetTask(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest("GET", "/tasks/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rec := httptest.NewRecorder()

	h.GetTask(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %v", rec.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	h := newTestHandler()
	payload := `{"title":"Updated","description":"New Desc","completed":false}`
	req := httptest.NewRequest("PUT", "/tasks/1", bytes.NewBufferString(payload))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rec := httptest.NewRecorder()

	h.UpdateTask(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %v", rec.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest("DELETE", "/tasks/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rec := httptest.NewRecorder()

	h.DeleteTask(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %v", rec.Code)
	}
}

func TestListTasks(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest("GET", "/tasks", nil)
	rec := httptest.NewRecorder()

	h.ListTasks(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %v", rec.Code)
	}
}

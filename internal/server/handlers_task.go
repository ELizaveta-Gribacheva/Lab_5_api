package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/ELizaveta-Gribacheva/Lab_5_api/db/sqlc"
	"github.com/gorilla/mux"
)

type TaskHandler struct {
	store db.Store
}

func NewTaskHandler(store db.Store) *TaskHandler {
	return &TaskHandler{store: store}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req db.CreateTaskParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("CreateTask error: %v\n", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := ValidateTaskInput(req.Title, req.Description); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.store.CreateTask(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task, err := h.store.GetTask(r.Context(), id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.store.ListTasks(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"tasks": tasks})
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req db.UpdateTaskParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	req.ID = id

	if err := ValidateTaskInput(req.Title, req.Description); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.store.UpdateTask(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.store.DeleteTask(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getIDParam(r *http.Request) (int32, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(id64), nil
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

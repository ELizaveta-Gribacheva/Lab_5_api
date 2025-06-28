package server

import (
	"net/http"

	db "github.com/ELizaveta-Gribacheva/Lab_5_api/db/sqlc"
	"github.com/gorilla/mux"
)

func SetupRouter(store db.Store) *mux.Router {
	router := mux.NewRouter()

	taskHandler := NewTaskHandler(store)

	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/tasks", taskHandler.ListTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods(http.MethodDelete)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	return router
}

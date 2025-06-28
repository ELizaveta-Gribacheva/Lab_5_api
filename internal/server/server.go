package server

import (
	"log"
	"net/http"

	db "github.com/ELizaveta-Gribacheva/Lab_5_api/db/sqlc"
)

type Server struct {
	router http.Handler
}

func NewServer(store db.Store) *Server {
	router := SetupRouter(store)
	return &Server{router: router}
}

func (s *Server) Run(port string) {
	log.Printf("Listening on %s...\n", port)
	err := http.ListenAndServe(port, s.router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

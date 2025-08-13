package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/log6"
)

type Server struct {
	Server http.Server
}

// Here we initialise HTTP server
func InitServer(logErr *log.Logger) *Server {
	mux := http.NewServeMux()
	server := &Server{
		Server: http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logErr,
			ReadTimeout:  (time.Second * 5),
			WriteTimeout: (time.Second * 10),
			IdleTimeout:  (time.Second * 15),
		},
	}

	// Register two handlers
	mux.HandleFunc("GET /", handlers.GetRoot)
	mux.HandleFunc("POST /upload", handlers.Upload)

	// Init server and check errors
	log6.Info.Printf("Init server complete")
	return server

}

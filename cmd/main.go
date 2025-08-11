package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/log6"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Initialize a server
	srv := server.InitServer(log6.Err)

	// Start a server
	err := srv.Server.ListenAndServe()
	if err != nil {
		// Check logs
		srv.Server.ErrorLog.Fatal("Ошибка при запуске сервера: ", err)
	}
}

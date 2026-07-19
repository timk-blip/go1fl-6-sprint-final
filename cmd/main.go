package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "[SERVER]", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	srv := server.NewServer(logger)
	srv.Logger.Println("Сервер запущен на порту :8080...")
	if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}
}

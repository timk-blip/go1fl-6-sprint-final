package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	r := chi.NewRouter()
	r.Get("/", handlers.HandleIndex)
	r.Post("/upload", handlers.HandleUpload)
	server := &http.Server{
		Addr:         ":8080",          // Порт 8080
		Handler:      r,                // Ваш http-роутер
		ErrorLog:     logger,           // Ваш логгер для системных ошибок сервера
		ReadTimeout:  5 * time.Second,  // Таймаут для чтения: 5 секунд
		WriteTimeout: 10 * time.Second, // Таймаут для записи: 10 секунд
		IdleTimeout:  15 * time.Second, // Таймаут ожидания следующего запроса: 15 секунд
	}
	return &Server{logger, server}
}

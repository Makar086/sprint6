package server

import (
	"log"
	"net/http"
	"time"
)


type Server struct {
	logger *log.Logger
	server *http.Server
}


func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	registerHandlers(mux, logger)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: httpServer,
	}
}


func registerHandlers(mux *http.ServeMux, logger *log.Logger) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Получен новый запрос")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	})
	

}

func (s *Server) Start() error {
	s.logger.Printf("Сервер запускается на порту %s...", s.server.Addr)
	return s.server.ListenAndServe()
}
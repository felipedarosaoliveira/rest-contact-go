package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//Server struct for application
type Server struct {
	server *http.Server
}

//New init a new Server
func New(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
	}
}

//ListenAndServe run the server into a new gorotine
func (s *Server) ListenAndServe() {
	go func() {
		fmt.Println("Server up!")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("erro: %s\n", err)
		}
	}()
}

//Shudown finish the server
func (s *Server) Shudown() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	s.server.Shutdown(ctx)
	fmt.Println("Server down!")
}

package server

import (
	"net/http"
	"time"
)

type HttpServer struct {
	server  *http.Server
	port    string
	handler http.Handler
}

func NewHttpServer(port string, handler http.Handler) *HttpServer {
	return &HttpServer{
		port:    port,
		handler: handler,
	}
}

func (s *HttpServer) Run() error {
	s.server = &http.Server{
		Addr:           ":" + s.port,
		Handler:        s.handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.server.ListenAndServe()
}

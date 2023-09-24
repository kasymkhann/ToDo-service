package todoprojectgo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpSV *http.Server
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.httpSV = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 0, // by default 0
		ReadTimeout:    12 * time.Second,
		WriteTimeout:   12 * time.Second,
	}
	return s.httpSV.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.ShutDown(ctx)
}

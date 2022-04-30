// структура сервера, которую будем использовать для запуска http сервера
package internalhttp

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Server struct {
	httpServer *http.Server
}

type Application interface { // TODO
}

func NewServer(logger Logger, app Application) *Server {
	return &Server{}
}

// Start http server
func (s *Server) Start(ctx context.Context, port string, handler http.Handler) error {
	s.log.Debug("HTTP server started")
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return errors.Wrap(err, "No connection established.")
	}
	return nil
}

/* func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
} */

// Stop http server
func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Debug("HTTP server stopped")
	return s.httpServer.Shutdown(ctx)
}

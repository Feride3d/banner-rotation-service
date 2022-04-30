package internalhttp

import (
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO
	})
}

/* type Handler struct {
	services *service.Service // dependency injunction
}

func NewHandler(services *service.Service) *Handler { // dependency injunction
	return &Handler{services: services}
}

func initHandlers(pool *pgxpool.Pool) http.Handler { // connection pool for concurrency
	r := mux.NewRouter() */

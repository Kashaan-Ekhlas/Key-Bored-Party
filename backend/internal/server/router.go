package server

import (
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/auth"
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/health"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Frontend Serving Routes
	mux.HandleFunc("GET /{$}", Root)

	// Auth Routes
	mux.HandleFunc("POST /auth/login", auth.Login)
	mux.HandleFunc("POST /auth/register", auth.Register)

	// Health Route(s?)
	mux.HandleFunc("GET /health", health.HealthCheck)

	return mux
}

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Key Bored Party\n"))
}

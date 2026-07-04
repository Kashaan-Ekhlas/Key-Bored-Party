package main

import (
	"log"
	"net/http"

	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/config"
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mux := server.NewRouter()

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}

	log.Println("Listening on port 5005")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

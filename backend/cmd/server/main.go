package main

import (
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/config"
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/database"
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/server"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.NewPool(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

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

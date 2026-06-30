package main

import (
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/server"
	"log"
	"net/http"
)

func main() {
	mux := server.NewRouter()

	server := &http.Server{
		Addr:    ":5005",
		Handler: mux,
	}

	log.Println("Listening on port 5005")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

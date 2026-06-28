package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

type Response struct {
  Status string `json:"status"`
	RandomNumber int `json:"randomNumber"`
}

func rootController(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to HooperParty"))
}

func jsonController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	Response := Response{
		Status: "okay",
		RandomNumber: rand.IntN(100), 
	}
	json.NewEncoder(w).Encode(Response)
}

func statusCodeController(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
}

func main(){
  mux := http.NewServeMux()

	mux.HandleFunc("/", rootController)
	mux.HandleFunc("/json", jsonController)
	mux.HandleFunc("/status", statusCodeController)

	server:= &http.Server{
		Addr: ":5005",
		Handler: mux,
	}

	fmt.Println("Listening on port 5005")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

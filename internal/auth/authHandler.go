package auth

import (
	"encoding/json"
	"net/http"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginPayload AuthPayload
	err := json.NewDecoder(r.Body).Decode(&loginPayload)
	if err != nil || loginPayload.Email == "" || loginPayload.Password == "" {
		http.Error(w, "Invalid JSON Body", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var registerPayload AuthPayload
	err := json.NewDecoder(r.Body).Decode(&registerPayload)
	if err != nil || registerPayload.Email == "" || registerPayload.Password == "" {
		http.Error(w, "Invalid JSON Body", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

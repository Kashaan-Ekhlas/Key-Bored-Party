package auth

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func hashPassword(password string) (string, error) {
	return "dummy string", nil
}

func emailFormatCheck(email string) error {
	if strings.Contains(email, " ") {
		return errors.New("invalid email format")
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email format")
	}

	if addr.Address != email {
		return errors.New("invalid email format")
	}
	return nil
}

func passwordFormatCheck(password string) error {
	return nil
}

func validatePayload(body io.ReadCloser) (AuthPayload, error) {
	var authPayload AuthPayload

	err := json.NewDecoder(body).Decode(&authPayload)
	if err != nil {
		return authPayload, err
	}

	authPayload.Email = strings.TrimSpace(authPayload.Email)
	authPayload.Password = strings.TrimSpace(authPayload.Password)

	if authPayload.Email == "" || authPayload.Password == "" {
		return authPayload, errors.New("email or password empty")
	}

	if len(authPayload.Email) > 254 {
		return authPayload, errors.New("email too large")
	}

	if len(authPayload.Password) > 128 {
		return authPayload, errors.New("password too large")
	}

	if err := emailFormatCheck(authPayload.Email); err != nil {
		return authPayload, err
	}

	return authPayload, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 8<<10)
	if r.ContentLength > 8<<10 {
		http.Error(w, "request too large", http.StatusRequestEntityTooLarge)
		return
	}

	loginPayload, err := validatePayload(r.Body)
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}

	loginPayload.Password, err = hashPassword(loginPayload.Password)
	if err != nil {
		log.Println("hash function failure")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// CompareHash()

	w.WriteHeader(http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 8<<10)
	if r.ContentLength > 8<<10 {
		http.Error(w, "request too large", http.StatusRequestEntityTooLarge)
		return
	}

	registerPayload, err := validatePayload(r.Body)
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}

	if err := passwordFormatCheck(registerPayload.Password); err != nil {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}

	registerPayload.Password, err = hashPassword(registerPayload.Password)
	if err != nil {
		log.Println("hash function failure")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// createAccount()

	w.WriteHeader(http.StatusCreated)
}

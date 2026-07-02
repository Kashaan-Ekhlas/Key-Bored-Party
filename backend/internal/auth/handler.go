package auth

import (
	"encoding/json"
	"errors"
	"github.com/alexedwards/argon2id"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
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

func passwordStrengthCheck(registerPayload AuthPayload) error {
	if len(registerPayload.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	strength := zxcvbn.PasswordStrength(registerPayload.Password, []string{registerPayload.Email})
	if strength.Score < 3 {
		return errors.New("weak password")
	}
	return nil
}

func validatePayload(body io.ReadCloser) (AuthPayload, error) {
	var authPayload AuthPayload

	decoder := json.NewDecoder(body)

	if err := decoder.Decode(&authPayload); err != nil {
		return authPayload, err
	}

	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return authPayload, errors.New("invalid json")
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

	// to do: check email existence
	// to do: retrieve hashedpassword for comparison if email exists, use dummy hash if not

	mockHash := "$argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG"
	match, err := argon2id.ComparePasswordAndHash(loginPayload.Password, mockHash)
	if err != nil {
		log.Println("hash function failure")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !match {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}

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

	if err := passwordStrengthCheck(registerPayload); err != nil {
		http.Error(w, "weak password", http.StatusBadRequest)
		return
	}

	registerPayload.Password, err = argon2id.CreateHash(registerPayload.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println("hash function failure")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// todo: createAccount()

	w.WriteHeader(http.StatusCreated)
}

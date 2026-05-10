package json

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Read(r *http.Request, dest any) error {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dest); err != nil {
		return err
	}

	if decoder.More() {
		return errors.New("request body must contain a single JSON object")
	}

	return nil
}

func Write(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if payload == nil {
		return
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, status int, message string) {
	Write(w, status, map[string]string{"error": message})
}

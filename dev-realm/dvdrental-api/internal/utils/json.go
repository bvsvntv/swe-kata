package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"dvdrental-api/internal/types"
)

func RespondWithError(w http.ResponseWriter, status int, msg string) {
	if status > http.StatusInternalServerError {
		log.Printf("Responding with %d error: %s", status, msg)
	}

	RespondWithJSON(w, status, types.ErrorResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, status int, payload any) {
	res, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

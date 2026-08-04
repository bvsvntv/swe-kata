package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetUrlID(r *http.Request, param string) (int32, error) {
	value := chi.URLParam(r, param)

	id, err := strconv.Atoi(value)
	if err != nil || id <= 0 {
		return 0, err
	}

	return int32(id), nil
}

func GetQueryInt(r *http.Request, key string, defaultValue int) int {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultValue
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return result
}

func DecodeJSON(r *http.Request, payload any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(payload)
}

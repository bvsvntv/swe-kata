package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"dvdrental-api/config"
	"dvdrental-api/models"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	var films []models.Film

	limit := 50
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err != nil {
			limit = parsed
		}
	}

	if err := config.DB.Limit(limit).Find(&films).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(films)
}

package films

import (
	"log"
	"net/http"

	"dvdrental-api/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) GetFilms(w http.ResponseWriter, r *http.Request) {
	films, err := h.service.GetFilms(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, films)
}

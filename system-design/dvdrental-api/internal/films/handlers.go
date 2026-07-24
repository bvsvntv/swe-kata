package films

import (
	"log"
	"net/http"
	"strconv"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/json"
	"dvdrental-api/internal/types"
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
	// Extract page, limit from query parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Calculate offset
	offset := (page - 1) * limit

	arg := repo.GetFilmsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	films, total, err := h.service.GetFilms(r.Context(), arg)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := FilmsResponse{
		Films: films,
		Pagination: types.Pagination{
			Page:        page,
			Limit:       limit,
			Total:       total,
			TotalPages:  totalPages,
			HasNextPage: page < totalPages,
			HasPrevPage: page > 1,
		},
	}

	json.Write(w, http.StatusOK, resp)
}

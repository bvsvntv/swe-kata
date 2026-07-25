package films

import (
	"net/http"
	"strconv"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
	"dvdrental-api/internal/utils"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) FetchFilms(w http.ResponseWriter, r *http.Request) {
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

	arg := repo.FetchFilmsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	films, total, err := h.service.FetchFilms(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
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

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

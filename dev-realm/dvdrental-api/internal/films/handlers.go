package films

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
	"dvdrental-api/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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
		PaginatedResponse: types.PaginatedResponse{
			Pagination: types.Pagination{
				Page:        page,
				Limit:       limit,
				Total:       total,
				TotalPages:  totalPages,
				HasNextPage: page < totalPages,
				HasPrevPage: page > 1,
			},
		},
		MessageResponse: types.MessageResponse{
			Message: "Films has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetFilm(w http.ResponseWriter, r *http.Request) {
	filmIDString := chi.URLParam(r, "filmID")
	filmID, err := strconv.Atoi(filmIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse film id: %v", err))
		return
	}

	film, err := h.service.GetFilm(r.Context(), int32(filmID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Film not found")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, FilmResponse{
		Film: film,
		MessageResponse: types.MessageResponse{
			Message: "Film detail has been fetched successfully.",
		},
	})
}

func (h *handler) FetchFilmActors(w http.ResponseWriter, r *http.Request) {
	actorIDString := chi.URLParam(r, "filmID")
	filmID, err := strconv.Atoi(actorIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse id: %v", err))
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	arg := repo.FetchFilmActorsParams{
		FilmID: int16(filmID),
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	actors, err := h.service.FetchFilmActors(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := FilmActorsResponse{
		Actors: actors,
		MessageResponse: types.MessageResponse{
			Message: "Actors has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) FetchFilmCategories(w http.ResponseWriter, r *http.Request) {
	actorIDString := chi.URLParam(r, "filmID")
	filmID, err := strconv.Atoi(actorIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse id: %v", err))
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	arg := repo.FetchFilmCategoriesParams{
		FilmID: int16(filmID),
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	categories, err := h.service.FetchFilmCategories(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := FilmCategoriesResponse{
		Categoryies: categories,
		MessageResponse: types.MessageResponse{
			Message: "Categories has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

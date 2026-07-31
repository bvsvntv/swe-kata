package languages

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

func (h *handler) FetchLanguages(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	arg := repo.FetchLanguagesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	languages, total, err := h.service.FetchLanguages(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := LanguagesResponse{
		Languages: languages,
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
			Message: "Languages has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetLanguage(w http.ResponseWriter, r *http.Request) {
	languageIDString := chi.URLParam(r, "languageID")
	actorID, err := strconv.Atoi(languageIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse language id: %v", err))
		return
	}

	language, err := h.service.GetLanguage(r.Context(), int32(actorID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Language not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, LanguageResponse{
		Language: language,
		MessageResponse: types.MessageResponse{
			Message: "Language detail has been fetched successfully.",
		},
	})
}

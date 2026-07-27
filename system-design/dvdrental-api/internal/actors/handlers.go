package actors

import (
	"encoding/json"
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

func (h *handler) FetchActors(w http.ResponseWriter, r *http.Request) {
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

	arg := repo.FetchActorsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	actors, total, err := h.service.FetchActors(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := ActorsResponse{
		Actors: actors,
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

func (h *handler) GetActor(w http.ResponseWriter, r *http.Request) {
	actorIDString := chi.URLParam(r, "actorID")
	actorID, err := strconv.Atoi(actorIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse actor id: %v", err))
		return
	}

	actor, err := h.service.GetActor(r.Context(), int32(actorID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Actor not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, ActorResponse{
		Actor: actor,
	})
}

func (h *handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	args := CreateActorRequest{}

	err := decoder.Decode(&args)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	actor, err := h.service.CreateActor(r.Context(), repo.CreateActorParams{
		FirstName: args.FirstName,
		LastName:  args.LastName,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create actor.\nERROR: %v", err))
	}

	utils.RespondWithJSON(w, http.StatusCreated, ActorResponse{
		Actor: actor,
	})
}

func (h *handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	actorIDString := chi.URLParam(r, "actorID")
	actorID, err := strconv.Atoi(actorIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse actor id: %v", err))
		return
	}

	err = h.service.DeleteActor(r.Context(), int32(actorID))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete actor.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}

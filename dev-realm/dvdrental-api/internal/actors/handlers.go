package actors

import (
	"errors"
	"fmt"
	"net/http"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
	"dvdrental-api/internal/utils"

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
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

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
			Message: "Actors has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetActor(w http.ResponseWriter, r *http.Request) {
	actorID, err := utils.GetUrlID(r, "actorID")
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
		MessageResponse: types.MessageResponse{
			Message: "Actor detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	req := ActorRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreateActorParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	actor, err := h.service.CreateActor(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create actor.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, ActorResponse{
		Actor: actor,
		MessageResponse: types.MessageResponse{
			Message: "Actor has been created successfully.",
		},
	})
}

func (h *handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	actorID, err := utils.GetUrlID(r, "actorID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse actor id: %v", err))
		return
	}

	err = h.service.DeleteActor(r.Context(), int32(actorID))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete actor.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Actor has been deleted successfully.",
	})
}

func (h *handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	actorID, err := utils.GetUrlID(r, "actorID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse actor id: %v", err))
		return
	}

	req := ActorRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateActorParams{
		ActorID:   int32(actorID),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	actor, err := h.service.UpdateActor(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update actor.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, ActorResponse{
		Actor: actor,
		MessageResponse: types.MessageResponse{
			Message: "Actor has been updated successfully.",
		},
	})
}

func (h *handler) UpdateActorPartial(w http.ResponseWriter, r *http.Request) {
	actorID, err := utils.GetUrlID(r, "actorID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse actor id: %v", err))
		return
	}

	req := UpdateActorPartialRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateActorPartialParams{
		ActorID:   int32(actorID),
		FirstName: utils.ToText(req.FirstName),
		LastName:  utils.ToText(req.LastName),
	}

	actor, err := h.service.UpdateActorPartial(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update actor.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, ActorResponse{
		Actor: actor,
		MessageResponse: types.MessageResponse{
			Message: "Actor has been updated successfully.",
		},
	})
}

func (h *handler) FetchActorFilms(w http.ResponseWriter, r *http.Request) {
	actorID, err := utils.GetUrlID(r, "actorID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse actor id: %v", err))
		return
	}

	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchActorFilmsParams{
		ActorID: int16(actorID),
		Limit:   int32(limit),
		Offset:  int32(offset),
	}

	films, err := h.service.FetchActorFilms(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := ActorFilmsResponse{
		Films: films,
		MessageResponse: types.MessageResponse{
			Message: "Films has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

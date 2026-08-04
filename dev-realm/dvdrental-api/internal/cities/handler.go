package cities

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

func (h *handler) FetchCities(w http.ResponseWriter, r *http.Request) {
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchCitiesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	cities, total, err := h.service.FetchCities(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := CitiesResponse{
		Cities: cities,
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
			Message: "Cities has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetCity(w http.ResponseWriter, r *http.Request) {
	cityID, err := utils.GetUrlID(r, "cityID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse city id: %v", err))
		return
	}

	city, err := h.service.GetCity(r.Context(), cityID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "City not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CityResponse{
		City: city,
		MessageResponse: types.MessageResponse{
			Message: "City detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreateCity(w http.ResponseWriter, r *http.Request) {
	req := CityRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreateCityParams{
		CountryID: int16(req.CountryID),
		City:      req.City,
	}

	city, err := h.service.CreateCity(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create city.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, CityResponse{
		City: city,
		MessageResponse: types.MessageResponse{
			Message: "City has been created successfully.",
		},
	})
}

func (h *handler) DeleteCity(w http.ResponseWriter, r *http.Request) {
	cityID, err := utils.GetUrlID(r, "cityID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse city id: %v", err))
		return
	}

	err = h.service.DeleteCity(r.Context(), cityID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete city.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "City has been deleted successfully.",
	})
}

func (h *handler) UpdateCity(w http.ResponseWriter, r *http.Request) {
	cityID, err := utils.GetUrlID(r, "cityID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse city id: %v", err))
		return
	}

	req := CityRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateCityParams{
		CityID:    cityID,
		CountryID: int16(req.CountryID),
		City:      req.City,
	}

	city, err := h.service.UpdateCity(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update city.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CityResponse{
		City: city,
		MessageResponse: types.MessageResponse{
			Message: "City has been updated successfully.",
		},
	})
}

func (h *handler) UpdateCityPartial(w http.ResponseWriter, r *http.Request) {
	cityID, err := utils.GetUrlID(r, "cityID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse city id: %v", err))
		return
	}

	req := UpdateCityPartialRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateCityPartialParams{
		CityID:    cityID,
		CountryID: utils.ToInt2(req.CountryID),
		City:      utils.ToText(req.City),
	}

	city, err := h.service.UpdateCityPartial(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update city.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CityResponse{
		City: city,
		MessageResponse: types.MessageResponse{
			Message: "City has been updated successfully.",
		},
	})
}

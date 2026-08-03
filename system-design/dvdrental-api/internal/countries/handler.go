package countries

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

func (h *handler) FetchCountries(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	arg := repo.FetchCountriesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	countries, total, err := h.service.FetchCountries(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := CountriesResponse{
		Countries: countries,
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
			Message: "Countries has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetCountry(w http.ResponseWriter, r *http.Request) {
	countryIDString := chi.URLParam(r, "countryID")
	countryID, err := strconv.Atoi(countryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse country id: %v", err))
		return
	}

	country, err := h.service.GetCountry(r.Context(), int32(countryID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Country not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CountryResponse{
		Country: country,
		MessageResponse: types.MessageResponse{
			Message: "Country has been updated successfully.",
		},
	})
}

func (h *handler) CreateCountry(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	args := CreateCountryRequest{}

	err := decoder.Decode(&args)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	country, err := h.service.CreateCountry(r.Context(), args.Country)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create country.\nERROR: %v", err))
	}

	utils.RespondWithJSON(w, http.StatusCreated, CountryResponse{
		Country: country,
		MessageResponse: types.MessageResponse{
			Message: "Country has been created successfully.",
		},
	})
}

func (h *handler) DeleteCountry(w http.ResponseWriter, r *http.Request) {
	countryIDString := chi.URLParam(r, "countryID")
	countryID, err := strconv.Atoi(countryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse country id: %v", err))
		return
	}

	err = h.service.DeleteCountry(r.Context(), int32(countryID))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete country.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}

func (h *handler) UpdateCountry(w http.ResponseWriter, r *http.Request) {
	countryIDString := chi.URLParam(r, "countryID")
	countryID, err := strconv.Atoi(countryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse country id: %v", err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	args := UpdateCountryRequest{}

	err = decoder.Decode(&args)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	country, err := h.service.UpdateCountry(r.Context(), repo.UpdateCountryParams{
		CountryID: int32(countryID),
		Country:   args.Country,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update country.\nERROR: %v", err))
	}

	utils.RespondWithJSON(w, http.StatusOK, CountryResponse{
		Country: country,
		MessageResponse: types.MessageResponse{
			Message: "Country has been updated successfully.",
		},
	})
}

func (h *handler) FetchCountryCities(w http.ResponseWriter, r *http.Request) {
	countryIDString := chi.URLParam(r, "countryID")
	countryID, err := strconv.Atoi(countryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse country id: %v", err))
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

	arg := repo.FetchCountryCitiesParams{
		CountryID: int16(countryID),
		Limit:     int32(limit),
		Offset:    int32(offset),
	}

	cities, err := h.service.FetchCountryCities(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return

	}

	utils.RespondWithJSON(w, http.StatusOK, CountryCitiesResponse{
		Cities: cities,
		MessageResponse: types.MessageResponse{
			Message: "Cities has been fetched successfully.",
		},
	})
}

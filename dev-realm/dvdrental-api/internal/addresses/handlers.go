package addresses

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

func (h *handler) FetchAddresses(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	arg := repo.FetchAddressesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	addresses, total, err := h.service.FetchAddresses(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := AddressesResponse{
		Addresses: addresses,
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
			Message: "Addresses have been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetAddress(w http.ResponseWriter, r *http.Request) {
	addressIDString := chi.URLParam(r, "addressID")
	addressID, err := strconv.Atoi(addressIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse address id: %v", err))
		return
	}

	address, err := h.service.GetAddress(r.Context(), int32(addressID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Address not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, AddressResponse{
		Address: address,
		MessageResponse: types.MessageResponse{
			Message: "Address detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := AddressRequest{}

	err := decoder.Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	createParams := repo.CreateAddressParams{
		Address:    req.Address,
		District:   req.District,
		CityID:     int16(req.CityID),
		Phone:      req.Phone,
		Address2:   utils.ToText(req.Address2),
		PostalCode: utils.ToText(req.PostalCode),
	}

	address, err := h.service.CreateAddress(r.Context(), createParams)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create address.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, AddressResponse{
		Address: address,
		MessageResponse: types.MessageResponse{
			Message: "Address has been created successfully.",
		},
	})
}

func (h *handler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	addressIDString := chi.URLParam(r, "addressID")
	addressID, err := strconv.Atoi(addressIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse address id: %v", err))
		return
	}

	err = h.service.DeleteAddress(r.Context(), int32(addressID))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete address.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Address has been deleted successfully.",
	})
}

func (h *handler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	addressIDString := chi.URLParam(r, "addressID")
	addressID, err := strconv.Atoi(addressIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse address id: %v", err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	req := AddressRequest{}

	err = decoder.Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	args := repo.UpdateAddressParams{
		AddressID:  int32(addressID),
		Address:    req.Address,
		District:   req.District,
		CityID:     int16(req.CityID),
		Phone:      req.Phone,
		Address2:   utils.ToText(req.Address2),
		PostalCode: utils.ToText(req.PostalCode),
	}

	address, err := h.service.UpdateAddress(r.Context(), args)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update address.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, AddressResponse{
		Address: address,
		MessageResponse: types.MessageResponse{
			Message: "Address has been updated successfully.",
		},
	})
}

func (h *handler) UpdateAddressPartial(w http.ResponseWriter, r *http.Request) {
	addressIDString := chi.URLParam(r, "addressID")
	addressID, err := strconv.Atoi(addressIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse address id: %v", err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	req := UpdateAddressPartialRequest{}

	err = decoder.Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	args := repo.UpdateAddressPartialParams{
		AddressID:  int32(addressID),
		Address:    utils.ToText(req.Address),
		Address2:   utils.ToText(req.Address2),
		District:   utils.ToText(req.District),
		PostalCode: utils.ToText(req.PostalCode),
		Phone:      utils.ToText(req.Phone),
		CityID:     utils.ToInt2(req.CityID),
	}

	address, err := h.service.UpdateAddressPartial(r.Context(), args)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update address.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, AddressResponse{
		Address: address,
		MessageResponse: types.MessageResponse{
			Message: "Address has been updated successfully.",
		},
	})
}

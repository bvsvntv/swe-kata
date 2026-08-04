package addresses

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

func (h *handler) FetchAddresses(w http.ResponseWriter, r *http.Request) {
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

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
	addressID, err := utils.GetUrlID(r, "addressID")
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
	req := AddressRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreateAddressParams{
		Address:    req.Address,
		District:   req.District,
		CityID:     req.CityID,
		Phone:      req.Phone,
		Address2:   utils.ToText(req.Address2),
		PostalCode: utils.ToText(req.PostalCode),
	}

	address, err := h.service.CreateAddress(r.Context(), arg)
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
	addressID, err := utils.GetUrlID(r, "addressID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse address id: %v", err))
		return
	}

	err = h.service.DeleteAddress(r.Context(), addressID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete address.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Address has been deleted successfully.",
	})
}

func (h *handler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	addressID, err := utils.GetUrlID(r, "addressID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse address id: %v", err))
		return
	}

	req := AddressRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateAddressParams{
		AddressID:  int32(addressID),
		Address:    req.Address,
		District:   req.District,
		CityID:     int16(req.CityID),
		Phone:      req.Phone,
		Address2:   utils.ToText(req.Address2),
		PostalCode: utils.ToText(req.PostalCode),
	}

	address, err := h.service.UpdateAddress(r.Context(), arg)
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
	addressID, err := utils.GetUrlID(r, "addressID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse address id: %v", err))
		return
	}

	req := UpdateAddressPartialRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateAddressPartialParams{
		AddressID:  int32(addressID),
		Address:    utils.ToText(req.Address),
		Address2:   utils.ToText(req.Address2),
		District:   utils.ToText(req.District),
		PostalCode: utils.ToText(req.PostalCode),
		Phone:      utils.ToText(req.Phone),
		CityID:     utils.ToInt2(req.CityID),
	}

	address, err := h.service.UpdateAddressPartial(r.Context(), arg)
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

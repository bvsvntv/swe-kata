package staffs

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

func (h *handler) FetchStaffs(w http.ResponseWriter, r *http.Request) {
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchStaffsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	staffs, total, err := h.service.FetchStaffs(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := StaffsResponse{
		Staffs: staffs,
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
			Message: "Staffs has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetStaff(w http.ResponseWriter, r *http.Request) {
	staffID, err := utils.GetUrlID(r, "staffID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse staff id: %v", err))
		return
	}

	staff, err := h.service.GetStaff(r.Context(), staffID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Staff not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, StaffResponse{
		Staff: staff,
		MessageResponse: types.MessageResponse{
			Message: "Staff detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreateStaff(w http.ResponseWriter, r *http.Request) {
	req := StaffRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreateStaffParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Active:    req.Active,
		AddressID: req.AddressID,
		StoreID:   req.StoreID,
		Email:     utils.ToText(req.Email),
		Password:  utils.ToText(req.Password),
	}

	staff, err := h.service.CreateStaff(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create staff.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, StaffResponse{
		Staff: staff,
		MessageResponse: types.MessageResponse{
			Message: "Staff has been created successfully.",
		},
	})
}

func (h *handler) DeleteStaff(w http.ResponseWriter, r *http.Request) {
	staffID, err := utils.GetUrlID(r, "staffID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse staff id: %v", err))
		return
	}

	err = h.service.DeleteStaff(r.Context(), staffID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete staff.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Staff has been deleted successfully.",
	})
}

func (h *handler) UpdateStaff(w http.ResponseWriter, r *http.Request) {
	staffID, err := utils.GetUrlID(r, "staffID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse staff id: %v", err))
		return
	}

	req := StaffRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateStaffParams{
		StaffID:   staffID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Active:    req.Active,
		AddressID: req.AddressID,
		StoreID:   req.StoreID,
		Email:     utils.ToText(req.Email),
		Password:  utils.ToText(req.Password),
	}

	staff, err := h.service.UpdateStaff(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update staff.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, StaffResponse{
		Staff: staff,
		MessageResponse: types.MessageResponse{
			Message: "Staff has been updated successfully.",
		},
	})
}

func (h *handler) UpdateStaffPartial(w http.ResponseWriter, r *http.Request) {
	staffID, err := utils.GetUrlID(r, "staffID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse staff id: %v", err))
		return
	}

	req := UpdateStaffPartialRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateStaffPartialParams{
		StaffID:   staffID,
		FirstName: utils.ToText(req.FirstName),
		LastName:  utils.ToText(req.LastName),
		Username:  utils.ToText(req.Username),
		Email:     utils.ToText(req.Email),
		Password:  utils.ToText(req.Password),
		AddressID: utils.ToInt2(req.AddressID),
		StoreID:   utils.ToInt2(req.StoreID),
		Active:    utils.ToBool(req.Active),
	}

	staff, err := h.service.UpdateStaffPartial(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update staff.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, StaffResponse{
		Staff: staff,
		MessageResponse: types.MessageResponse{
			Message: "Staff has been updated successfully.",
		},
	})
}

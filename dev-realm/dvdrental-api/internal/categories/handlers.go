package categories

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

func (h *handler) FetchCategories(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	arg := repo.FetchCategoriesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	categories, total, err := h.service.FetchCategories(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := CategoriesResponse{
		Categories: categories,
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
			Message: "Categories has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDString := chi.URLParam(r, "categoryID")
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse category id: %v", err))
		return
	}

	category, err := h.service.GetCategory(r.Context(), int32(categoryID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Category not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CategoryResponse{
		Category: category,
		MessageResponse: types.MessageResponse{
			Message: "Category has been updated successfully.",
		},
	})
}

func (h *handler) FetchCategoryFilms(w http.ResponseWriter, r *http.Request) {
	categoryIDString := chi.URLParam(r, "categoryID")
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse category id: %v", err))
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

	arg := repo.FetchCategoryFilmsParams{
		CategoryID: int16(categoryID),
		Limit:      int32(limit),
		Offset:     int32(offset),
	}

	films, err := h.service.FetchCategoryFilms(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := CategoryFilmsResponse{
		Films: films,
		MessageResponse: types.MessageResponse{
			Message: "Films has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	req := CategoryRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	category, err := h.service.CreateCategory(r.Context(), req.Name)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create category.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, CategoryResponse{
		Category: category,
		MessageResponse: types.MessageResponse{
			Message: "Category has been created successfully.",
		},
	})
}

func (h *handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDString := chi.URLParam(r, "categoryID")
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse category id: %v", err))
		return
	}

	err = h.service.DeleteCategory(r.Context(), int32(categoryID))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete category.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Category has been deleted successfully.",
	})
}

func (h *handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDString := chi.URLParam(r, "categoryID")
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse category id: %v", err))
		return
	}

	req := CategoryRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	category, err := h.service.UpdateCategory(r.Context(), repo.UpdateCategoryParams{
		CategoryID: int32(categoryID),
		Name:       req.Name,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update category.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CategoryResponse{
		Category: category,
		MessageResponse: types.MessageResponse{
			Message: "Category has been updated successfully.",
		},
	})
}

func (h *handler) UpdateCategoryPartial(w http.ResponseWriter, r *http.Request) {
	categoryIDString := chi.URLParam(r, "categoryID")
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse category id: %v", err))
		return
	}

	req := UpdateCategoryPartialRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	category, err := h.service.UpdateCategoryPartial(r.Context(), repo.UpdateCategoryPartialParams{
		CategoryID: int32(categoryID),
		Name:       utils.ToText(req.Name),
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update category.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CategoryResponse{
		Category: category,
		MessageResponse: types.MessageResponse{
			Message: "Category has been updated successfully.",
		},
	})
}

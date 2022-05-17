package services

import (
	"promo-tables/database/models"
	"promo-tables/utils"
	"promo-tables/utils/mock"
)

func (h *Handler) GetCategoriesByName(name string) []models.Category {
	categories, err := h.db.GetCategoriesByName(name)
	utils.CheckError(err)
	return categories
}

func (h *Handler) GetCategories() []models.Category {
	categories, err := h.db.GetCategories()
	utils.CheckError(err)
	return categories
}

func (h *Handler) AddCategories() error {
	categories := GenerateCategories(true)
	return h.db.AddCategories(categories)
}

func GenerateCategories(enabled bool) []models.Category {
	categories := []models.Category{}
	for _, categoryName := range mock.Categories {
		categories = append(categories, models.Category{
			Name:    categoryName,
			Enabled: enabled,
		})
	}
	return categories
}

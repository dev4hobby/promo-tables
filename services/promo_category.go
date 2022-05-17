package services

import (
	"math/rand"
	"promo-tables/database/models"
	"promo-tables/utils"
	"promo-tables/utils/mock"
)

func (h *Handler) GetAllPromoCategories() []models.PromoCategory {
	promoCategories, err := h.db.GetAllPromoCategories()
	utils.CheckError(err)
	return promoCategories
}

func (h *Handler) AddPromoCategories(howMany int) error {
	promoCategories := GeneratePromoCategory(true, howMany)
	return h.db.AddPromoCategories(promoCategories)
}

func GeneratePromoCategory(enabled bool, count int) []models.PromoCategory {
	promoCategories := []models.PromoCategory{}
	if len(mock.PromoCategory) < count {
		count = len(mock.PromoCategory)
	}

	loopCount := 0
	picked := []string{}
	for {
		promoCategory := mock.PromoCategory[rand.Intn(len(mock.PromoCategory))]
		if !utils.Contains(picked, promoCategory) {
			picked = append(picked, promoCategory)
			loopCount += 1
		}
		if len(picked) >= count {
			break
		}
	}

	for _, promoCategory := range picked {
		promoCategories = append(promoCategories, models.PromoCategory{
			Name:    promoCategory,
			Enabled: enabled,
		})
	}
	return promoCategories
}

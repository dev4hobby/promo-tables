package services

import (
	"promo-tables/database/models"
	"promo-tables/utils"
)

func (h *Handler) InitMockTable() {
	h.CreateEmptyTable()
	h.AddUsers(utils.GenerateRandomNumber(20, 200))
	if !h.db.RowExists(&models.Category{}) {
		h.AddCategories()
		h.AddPromoCategories(utils.GenerateRandomNumber(1, 5))
		h.AddPromoControllerByPromoCategories(h.GetAllPromoCategories(), true)
		h.AddProducts()
		h.AddUserStoreSubscriptions()
	}
}

func (h *Handler) FlushAllProgress() {
	h.db.FlushAll()
}

func (h *Handler) CreateEmptyTable() {
	// User
	h.db.InitTable(&models.User{})

	// store
	h.db.InitTable(&models.Order{})
	h.db.InitTable(&models.Product{})
	h.db.InitTable(&models.Category{})
	h.db.InitTable(&models.UserStoreSubscription{})

	// promotion
	h.db.InitTable(&models.PromoCategory{})
	h.db.InitTable(&models.PromoController{})
	h.db.InitTable(&models.PromoPurchasedLog{})
}

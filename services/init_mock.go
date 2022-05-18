package services

import (
	"promo-tables/database/models"
	"promo-tables/utils"
)

var tableInterfaceList = []interface{}{
	&models.User{},
	&models.Order{},
	&models.Product{},
	&models.Category{},
	&models.PromoCategory{},
	&models.PromoController{},
	&models.PromoPurchasedLog{},
	&models.UserStoreSubscription{},
}

func (h *Handler) GetMockTableInterfaceList() []interface{} {
	return tableInterfaceList
}

func (h *Handler) InitMockTable() {
	env := utils.GetEnv()
	h.CreateEmptyTable(tableInterfaceList)
	if h.db.RowExists(&models.User{}, 1) < 1 {
		h.AddUsers(utils.JsonNumberToInt(env.MockUserCount))
		h.AddCategories()
		h.AddPromoCategories(utils.JsonNumberToInt(env.MockPromoCategoryCount))
		h.AddPromoControllerByPromoCategories(h.GetAllPromoCategories(), true)
		h.AddProducts()
		h.AddUserStoreSubscriptions()
	}
}

func (h *Handler) FlushAllProgress() {
	h.db.FlushAll()
}

func (h *Handler) CreateEmptyTable(tableInterfaceList []interface{}) {
	for _, tableInterface := range tableInterfaceList {
		h.db.InitTable(tableInterface)
	}
}

func (h *Handler) CheckTableExists(tableInterfaceList []interface{}) bool {
	for _, tableInterface := range tableInterfaceList {
		h.db.RowExists(tableInterface, 1)
	}
	return true
}

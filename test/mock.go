package test

import (
	"promo-tables/services"
	"promo-tables/utils"
	"promo-tables/utils/mock"
	"testing"
)

func InitMockTable(handler services.HandlerInterface) {
	handler.InitMockTable()
}

func CheckTableExists(handler services.HandlerInterface) bool {
	tableInterfaceList := handler.GetMockTableInterfaceList()
	return handler.CheckTableExists(tableInterfaceList)
}

func CheckUserTable(handler services.HandlerInterface, t *testing.T) {
	env := utils.GetEnv()
	users := handler.GetUsers()

	if len(users) != utils.JsonNumberToInt(env.MockUserCount) {
		t.Errorf("[FAIL] user table has %v users, not %v", len(users), utils.JsonNumberToInt(env.MockUserCount))
	}

	for _, user := range users {
		if !ValidateUUID(user.UUID) {
			t.Errorf("[FAIL] user.UUID is not a valid uuid")
		}
		if !ValidateEmail(user.Email) {
			t.Errorf("[FAIL] user.Email is not a valid email")
		}
		if !ValidateSHA256Password(user.Password) {
			t.Errorf("[FAIL] user.Password is not a valid password")
		}
	}
}

func CheckCategoryTable(handler services.HandlerInterface, t *testing.T) {
	categories := handler.GetCategories()

	categoryCount := len(mock.Categories)
	if len(categories) != categoryCount {
		t.Errorf("[FAIL] category table has %v categories, not %v", len(categories), categoryCount)
	}
}

func CheckPromoCategoryTable(handler services.HandlerInterface, t *testing.T) {
	env := utils.GetEnv()
	promoCategories := handler.GetAllPromoCategories()
	promoCategoryCount := utils.JsonNumberToInt(env.MockPromoCategoryCount)
	if len(promoCategories) != promoCategoryCount {
		t.Errorf("[FAIL] promoCategory table has %v promoCategories, not %v", len(promoCategories), promoCategoryCount)
	}
}

func CheckPromoControllerTable(handler services.HandlerInterface, t *testing.T) {
	env := utils.GetEnv()
	promoCategoryCount := utils.JsonNumberToInt(env.MockPromoCategoryCount)
	promoControllers := handler.GetAllPromoControllers()
	promoControllerCount := promoCategoryCount
	if len(promoControllers) != promoControllerCount {
		t.Errorf("[FAIL] promoController table has %v promoControllers, not %v", len(promoControllers), promoControllerCount)
	}
}

func CheckProductTable(handler services.HandlerInterface, t *testing.T) {
	products := handler.GetAllProducts()

	productCount := len(mock.Products)
	if len(products) != productCount {
		t.Errorf("[FAIL] product table has %v products, not %v", len(products), productCount)
	}
}

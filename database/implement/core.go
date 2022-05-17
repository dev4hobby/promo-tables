package implement

import "promo-tables/database/models"

type DBLayer interface {
	// Category
	AddCategories([]models.Category) error
	GetCategories() ([]models.Category, error)
	GetCategoriesByName(string) ([]models.Category, error)

	// Promo
	AddPromoCategories([]models.PromoCategory) error
	AddPromoControllers([]models.PromoController) error
	GetAllPromoCategories() ([]models.PromoCategory, error)
	GetAllPromoControllers() ([]models.PromoController, error)

	// User
	GetUsersFilteredBySubscribeCondition(models.PromoControllerSubscribeCondition) []models.User
	AddUsers([]models.User) error
	AddUserStoreSubscriptions([]models.UserStoreSubscription) error

	// Product
	AddProducts([]models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(int) (models.Product, error)

	// Misc
	RowExists(interface{}) bool
	InitTable(interface{})
	FlushAll()
}

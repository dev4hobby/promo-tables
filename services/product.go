package services

import (
	"promo-tables/database/models"
	"promo-tables/utils"
	"promo-tables/utils/mock"
	"strings"
)

func (h *Handler) GetProductByID(id int) models.Product {
	product, err := h.db.GetProductByID(id)
	utils.CheckError(err)
	return product
}

func (h *Handler) GetAllProducts() []models.Product {
	products, err := h.db.GetAllProducts()
	utils.CheckError(err)
	return products
}

func (h *Handler) AddProducts() error {
	categories := h.GetCategories()
	promoCategories := h.GetAllPromoCategories()
	products := BindProductWithCategoryInfo(
		GenerateProducts(true),
		categories,
		promoCategories,
	)
	return h.db.AddProducts(products)
}

func GenerateProducts(enabled bool) []models.Product {
	products := []models.Product{}
	status := "NOT_ON_SALE"
	if enabled {
		status = "ON_SALE"
	}

	for productName, price := range mock.Products {
		products = append(products, models.Product{
			Name:   productName,
			Price:  utils.IntegerToString(price),
			Status: status,
		})
	}
	return products
}

func BindProductWithCategoryInfo(products []models.Product, categories []models.Category, promoCategories []models.PromoCategory) []models.Product {
	categoryIDMap := map[string]int{}
	promoCategoryIDs := []int{}

	for _, category := range categories {
		categoryIDMap[category.Name] = category.ID
	}

	for _, promoCategory := range promoCategories {
		promoCategoryIDs = append(promoCategoryIDs, promoCategory.ID)
	}

	for index, product := range products {
		categoryIDList := []int{}
		for categoryName, categoryID := range categoryIDMap {
			if strings.Contains(product.Name, categoryName) {
				categoryIDList = append(categoryIDList, categoryID)
			}
		}
		products[index].CategoryIDList = utils.MustMarshal(categoryIDList)
		if utils.GenerateRandomNumber(1, 10) > 7 {
			products[index].PromoCategoryID = promoCategoryIDs[utils.PickRandomIndex(promoCategoryIDs)]
		}
	}
	return products
}

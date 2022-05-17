package implement

import (
	"promo-tables/database/models"
)

func (db *ORM) GetAllPromoCategories() (promoCategories []models.PromoCategory, err error) {
	return promoCategories, db.Find(&promoCategories).Error
}

func (db *ORM) AddPromoCategories(promoCategories []models.PromoCategory) error {
	return db.Create(&promoCategories).Error
}

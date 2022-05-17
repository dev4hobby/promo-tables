package implement

import "promo-tables/database/models"

func (db *ORM) AddPromoControllers(promoControllers []models.PromoController) error {
	return db.Create(&promoControllers).Error
}

func (db *ORM) GetAllPromoControllers() (promoControllers []models.PromoController, err error) {
	return promoControllers, db.Find(&promoControllers).Error
}

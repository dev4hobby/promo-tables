package implement

import "promo-tables/database/models"

func (db *ORM) AddUserStoreSubscriptions(subscription []models.UserStoreSubscription) error {
	return db.Create(&subscription).Error
}

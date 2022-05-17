package implement

import (
	"promo-tables/database/models"
)

func (db *ORM) AddUsers(user []models.User) error {
	return db.Create(&user).Error
}

func (db *ORM) GetUsersFilteredBySubscribeCondition(subscription models.PromoControllerSubscribeCondition) []models.User {
	var users []models.User
	switch subscription.Birth.Operator {
	case ">=":
		db.Where("sex = ? AND birth >= ?", subscription.Sex, subscription.Birth.Timestamp).Find(&users)
	default:
		db.Where("sex = ? AND birth <= ?", subscription.Sex, subscription.Birth.Timestamp).Find(&users)
	}
	return users
}

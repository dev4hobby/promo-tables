package implement

import (
	"promo-tables/database/models"
)

func (db *ORM) GetCategories() (categories []models.Category, err error) {
	return categories, db.Find(&categories).Error
}

func (db *ORM) GetCategoriesByName(name string) (categories []models.Category, err error) {
	return categories, db.Where("name like ?", "%"+name+"%").Find(&categories).Error
}

func (db *ORM) AddCategories(categories []models.Category) error {
	return db.Create(&categories).Error
}

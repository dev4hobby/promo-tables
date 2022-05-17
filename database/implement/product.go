package implement

import (
	"promo-tables/database/models"
)

func (db *ORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *ORM) GetProductByID(id int) (product models.Product, err error) {
	return product, db.Find(&product, id).Error
}

func (db *ORM) AddProducts(products []models.Product) error {
	return db.Create(&products).Error
}

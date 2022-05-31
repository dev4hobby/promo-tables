package implement

import (
	"promo-tables/database/models"
)

func (db *ORM) JoinTableAsQueryBuilder() (err error) {
	orders := []models.PromoController{}
	return db.Table("promo_controller pcon").Select("*").Joins("join promo_category pcat on pcon.promo_category_id=pcat.id").Scan(orders).Error
}

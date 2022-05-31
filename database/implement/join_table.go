package implement

import (
	"fmt"
	"promo-tables/database/models"
)

func (db *ORM) JoinTable() (err error) {
	var promoController []*models.PromoController
	db.Unscoped().Preload("promo_categories").Scan(&promoController)
	fmt.Println(promoController)
	return nil
}

package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                 int             `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Name               string          `json:"name"`
	Price              string          `json:"price"`
	ImageUrlList       json.RawMessage `json:"image_url_list" gorm:"type:json"`
	DescImageUrlList   json.RawMessage `json:"desc_image_url_list" gorm:"type:json"`
	OptionImageUrlList json.RawMessage `json:"option_image_url_list" gorm:"type:json"`
	CategoryIDList     json.RawMessage `json:"category_id_list" gorm:"type:json;"`
	PromoCategoryID    int             `json:"promo_category_id" gorm:"default:NULL"`
	Status             string          `json:"status" gorm:"type:varchar(20); default:'NOT_ON_SALE'"`
}

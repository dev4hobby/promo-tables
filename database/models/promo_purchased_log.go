package models

import (
	"time"
)

type PromoPurchasedLog struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	PromoCategoryID   int       `json:"promo_category_id"`
	PromoControllerID int       `json:"promo_controller_id"`
	Count             int       `json:"count"`
	IsCancelled       bool      `json:"is_cancelled"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

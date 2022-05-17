package models

import (
	"time"
)

type PromoCategory struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Ordering  int       `json:"ordering"`
	Enabled   bool      `json:"enabled" gorm:"default:false"`
}

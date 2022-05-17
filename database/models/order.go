package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        int       `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	UUID      string    `json:"uuid" gorm:"type:varchar(36);unique_index"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

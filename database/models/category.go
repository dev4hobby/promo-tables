package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        int       `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Name      string    `json:"name" gorm:"type:varchar(100);unique_index"`
	Enabled   bool      `json:"enabled" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

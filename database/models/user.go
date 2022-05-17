package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              int       `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	UUID            string    `json:"uuid" gorm:"type:varchar(36);unique_index"`
	Name            string    `json:"name" gorm:"type:varchar(120)"`
	Nickname        string    `json:"nickname"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Sex             string    `json:"sex"`
	Birth           time.Time `json:"birth"`
	Phone           string    `json:"phone"`
	ProfileImageUrl string    `json:"profile_image_url"`
	IsDeleted       bool      `json:"is_deleted"`
	IsBanned        bool      `json:"is_banned"`
	// ...
}

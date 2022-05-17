package models

import (
	"encoding/json"
)

// Disable ID column
type UserStoreSubscription struct {
	UserID                int             `json:"user_id" gorm:"primary_key; not null; index:idx_user_id"`
	PromoControllerIDList json.RawMessage `json:"promo_controller_id_list" gorm:"type:json"`
}

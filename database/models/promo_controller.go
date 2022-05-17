package models

import (
	"encoding/json"
	"time"
)

type PromoController struct {
	ID                 int             `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	PromoCategoryID    int             `json:"promo_category_id"`
	Description        string          `json:"description" gorm:"type:varchar(255)"`
	BuyLimitCount      int             `json:"buy_limit_count" gorm:"default:1"`
	StartAt            time.Time       `json:"start_at"`
	EndAt              time.Time       `json:"end_at"`
	Enabled            bool            `json:"enabled" gorm:"default:false"`
	SubscribeCondition json.RawMessage `json:"subscribe_condition" gorm:"type:json"`
}

type PromoControllerSubscribeCondition struct {
	Sex   string                  `json:"sex"`
	Birth SubscribeConditionBirth `json:"birth"`
}

type SubscribeConditionBirth struct {
	Operator  string    `json:"operator"`
	Timestamp time.Time `json:"timestamp"`
}

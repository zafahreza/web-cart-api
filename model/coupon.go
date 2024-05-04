package model

import "time"

type Coupon struct {
	Id                 int       `json:"id"`
	UserId             int       `json:"user_id"`
	Description        string    `json:"description"`
	DiscountPercentage int       `json:"discount_percentage"`
	MinimumPayment     int       `json:"minimum_payment"`
	MaximumDiscount    int       `json:"maximum_discount"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

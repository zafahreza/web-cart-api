package model

import "time"

type Cart struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	ProductId    int       `json:"product_id"`
	Product      Product   `json:"product"`
	ProductCount int       `json:"product_count"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

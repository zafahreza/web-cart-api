package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Id                    int         `json:"id"`
	UserId                int         `json:"user_id"`
	CouponId              int         `json:"coupon_id"`
	TotalBeforeDiscount   int         `json:"total_before_discount"`
	Discount              int         `json:"discount"`
	TotalAfterDiscount    int         `json:"total_after_discount"`
	Status                string      `gorm:"-" json:"status"`
	StatusConfirmation    string      `json:"status_confirmation"`
	OrderCouponReceived   int         `json:"order_coupon_received"`
	ProductCouponReceived int         `json:"product_coupon_received"`
	OrderItems            []OrderItem `json:"order_items"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
}

type OrderItem struct {
	Id           int       `json:"id"`
	OrderId      int       `json:"order_id"`
	ProductId    int       `json:"product_id"`
	Product      Product   `json:"product"`
	ProductCount int       `json:"product_count"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

const (
	OrderStatusOpen                  = "Open"
	OrderStatusClose                 = "Close"
	OrderStatusConfirmationPending   = "Pending"
	OrderStatusConfirmationConfirmed = "Confirmed"
)

func (o *Order) AfterFind(tx *gorm.DB) (err error) {

	dif := time.Now().Sub(o.CreatedAt)
	if dif > (3 * time.Hour) {
		o.Status = OrderStatusClose
	} else {
		o.Status = OrderStatusOpen
	}

	return
}

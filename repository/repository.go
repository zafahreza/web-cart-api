package repository

import (
	"gorm.io/gorm"
	"web-cart-api/model"
)

type Repository interface {
	CreateUser(tx *gorm.DB, user model.User) (err error)
	FindUserById(tx *gorm.DB, id int) (user model.User, err error)
	FindUserByUsername(tx *gorm.DB, username string) (user model.User, err error)
	FindProducts(tx *gorm.DB, name string, pagination model.QueryPagination) (count int64, products []model.Product)
	FindProductById(tx *gorm.DB, id int) (product model.Product, err error)
	FindCartByUserId(tx *gorm.DB, userId int) (count int64, carts []model.Cart)
	CreateProductToCart(tx *gorm.DB, cart model.Cart) (model.Cart, error)
	DeleteCart(tx *gorm.DB, cartId int) error
	CreateOrder(tx *gorm.DB, order model.Order) (model.Order, error)
	FindOrderById(tx *gorm.DB, id int) (order model.Order, err error)
	FindOrders(tx *gorm.DB, userId int) (count int64, orders []model.Order)
	UpdateOrder(tx *gorm.DB, id int, property map[string]interface{}) (order model.Order, err error)
	FindCartById(tx *gorm.DB, id int, userId int) (cart model.Cart, err error)
	FindCouponById(tx *gorm.DB, id int, userId int) (coupon model.Coupon, err error)
	CreateCoupon(tx *gorm.DB, coupons []model.Coupon) (err error)
	DeleteCarts(tx *gorm.DB, cartsId []int) error
	DeleteCoupon(tx *gorm.DB, couponId int) error
	FindCoupons(tx *gorm.DB, userId int) (count int64, coupons []model.Coupon)
}

package service

import (
	"web-cart-api/model"
)

type Service interface {
	CreateUser(req model.RequestSignup) (user model.User, err error)
	GetUserById(req model.RequestGetUserById) (user model.User, err error)
	Login(req model.RequestLogin) (response model.ResponseLogin, err error)
	GetProducts(req model.RequestGetProducts) (response model.ResponseProductList)
	GetProductById(req model.RequestGetProductById) (product model.Product, err error)
	GetCartByUserId(req model.RequestGetCartByUserId) (response model.ResponseCartList)
	AddProductToCart(req model.RequestAddProductToCart) error
	DeleteCart(req model.RequestDeleteCart) error
	CreateOrder(req model.RequestOrder) (order model.Order, err error)
	GetOrderById(req model.RequestGetOrderById) (order model.Order, err error)
	GetOrders(req model.RequestGetOrders) (response model.ResponseOrderList)
	UpdateStatusOrder(req model.RequestUpdateOrder) (err error)
	GetCoupons(req model.RequestGetCoupons) (response model.ResponseCouponList)
}

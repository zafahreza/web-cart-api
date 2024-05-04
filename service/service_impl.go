package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"web-cart-api/helper"
	"web-cart-api/model"
	"web-cart-api/repository"
)

type service struct {
	Repository repository.Repository
	Db         *gorm.DB
}

func NewService(repository repository.Repository, db *gorm.DB) Service {
	return &service{Repository: repository, Db: db}
}

func (s *service) GetUserById(req model.RequestGetUserById) (user model.User, err error) {
	user, err = s.Repository.FindUserById(s.Db, req.Id)
	return
}

func (s *service) CreateUser(req model.RequestSignup) (user model.User, err error) {

	if len(req.Name) < 1 && len(req.Username) < 1 && len(req.Password) < 1 {
		err = errors.New("invalid data requested")
		return
	}

	user, err = s.Repository.FindUserByUsername(s.Db, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if user.Id != 0 {
		err = errors.New("username already used")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return
	}

	user = model.User{
		Name:     req.Name,
		Username: req.Username,
		Password: string(passwordHash),
	}

	err = s.Repository.CreateUser(s.Db, user)
	if err != nil {
		return
	}

	return
}

func (s *service) Login(req model.RequestLogin) (response model.ResponseLogin, err error) {
	user, err := s.Repository.FindUserByUsername(s.Db, req.Username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return
	}

	response.User = user
	response.Token = token

	return
}

func (s *service) GetProducts(req model.RequestGetProducts) (response model.ResponseProductList) {
	pagination := helper.SetPaginationFromQuery(req.Limit, req.Page)

	count, products := s.Repository.FindProducts(s.Db, req.Name, pagination)

	response.Products = products
	response.Count = count
	return
}

func (s *service) GetProductById(req model.RequestGetProductById) (product model.Product, err error) {
	product, err = s.Repository.FindProductById(s.Db, req.Id)
	return
}

func (s *service) GetCartByUserId(req model.RequestGetCartByUserId) (response model.ResponseCartList) {
	count, carts := s.Repository.FindCartByUserId(s.Db, req.UserId)

	response.Carts = carts
	response.Count = count
	return
}

func (s *service) AddProductToCart(req model.RequestAddProductToCart) error {
	product, err := s.Repository.FindProductById(s.Db, req.ProductId)
	if err != nil {
		return err
	}

	cart := model.Cart{
		UserId:       req.UserId,
		ProductId:    product.Id,
		ProductCount: req.ProductCount,
		Note:         req.Note,
	}

	cart, err = s.Repository.CreateProductToCart(s.Db, cart)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteCart(req model.RequestDeleteCart) error {
	_, err := s.Repository.FindCartById(s.Db, req.Id, req.UserId)
	if err != nil {
		return err
	}

	err = s.Repository.DeleteCart(s.Db, req.Id)
	return err
}

func (s *service) CreateOrder(req model.RequestOrder) (order model.Order, err error) {
	var orderItems []model.OrderItem
	var totalPrice int
	var couponOrder, couponProduct int
	var cart model.Cart
	var coupon model.Coupon

	if len(req.CartIds) == 0 {
		err = errors.New("no product included")
		return
	}

	tx := s.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}

	}()

	for _, cartId := range req.CartIds {
		cart, err = s.Repository.FindCartById(tx, cartId, req.UserId)
		if err != nil {
			return
		}

		totalPrice = totalPrice + (cart.ProductCount * cart.Product.Price)

		orderItem := model.OrderItem{
			ProductId:    cart.ProductId,
			ProductCount: cart.ProductCount,
			Product:      cart.Product,
			Note:         cart.Note,
		}

		orderItems = append(orderItems, orderItem)

		//check get coupon if product price > 50000
		if cart.Product.Price > 50000 {
			couponProduct++
		}
	}

	//check get 1 coupon for each 100k total price
	couponOrder = totalPrice / 100000

	afterDiscount := totalPrice
	discount := 0

	if req.CouponId != 0 {
		coupon, err = s.Repository.FindCouponById(tx, req.CouponId, req.UserId)
		if err != nil {
			return
		}

		discount = totalPrice * coupon.DiscountPercentage / 100
		if discount > coupon.MaximumDiscount {
			discount = coupon.MaximumDiscount
		}

		afterDiscount = totalPrice - discount
	}

	order = model.Order{
		UserId:                req.UserId,
		CouponId:              req.CouponId,
		TotalBeforeDiscount:   totalPrice,
		TotalAfterDiscount:    afterDiscount,
		Discount:              discount,
		OrderItems:            orderItems,
		Status:                model.OrderStatusOpen,
		StatusConfirmation:    model.OrderStatusConfirmationPending,
		OrderCouponReceived:   couponOrder,
		ProductCouponReceived: couponProduct,
	}

	order, err = s.Repository.CreateOrder(tx, order)
	if err != nil {
		return
	}

	err = s.Repository.DeleteCarts(tx, req.CartIds)
	if err != nil {
		return
	}

	err = s.Repository.DeleteCoupon(tx, req.CouponId)
	if err != nil {
		return
	}

	tx.Commit()
	return
}

func (s *service) GetOrderById(req model.RequestGetOrderById) (order model.Order, err error) {
	order, err = s.Repository.FindOrderById(s.Db, req.Id)
	if err != nil {
		return
	}
	return
}

func (s *service) GetOrders(req model.RequestGetOrders) (response model.ResponseOrderList) {
	count, orders := s.Repository.FindOrders(s.Db, req.UserId)

	response.Count = count
	response.Orders = orders

	return
}

func (s *service) UpdateStatusOrder(req model.RequestUpdateOrder) (err error) {
	var order model.Order
	tx := s.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}

	}()

	order, err = s.Repository.FindOrderById(s.Db, req.Id)
	if err != nil {
		return
	}

	if order.UserId != req.UserId {
		err = errors.New("not found")
		return
	}

	switch req.Status {
	case model.OrderStatusConfirmationPending, model.OrderStatusConfirmationConfirmed:
	default:
		err = errors.New("invalid status")
		return
	}

	property := map[string]interface{}{
		"status_confirmation": req.Status,
	}

	totalCoupon := order.ProductCouponReceived + order.OrderCouponReceived

	if totalCoupon > 0 {
		var coupons []model.Coupon

		for i := 0; i < totalCoupon; i++ {
			coupon := model.Coupon{
				UserId:             order.UserId,
				Description:        fmt.Sprintf("discount %d percent maksimum %d after order %d", 10, 10000, order.Id),
				DiscountPercentage: 10,
				MinimumPayment:     50000,
				MaximumDiscount:    10000,
			}

			coupons = append(coupons, coupon)
		}

		err = s.Repository.CreateCoupon(tx, coupons)
		if err != nil {
			return
		}
	}

	_, err = s.Repository.UpdateOrder(tx, req.Id, property)
	if err != nil {
		return
	}

	tx.Commit()

	return
}

func (s *service) GetCoupons(req model.RequestGetCoupons) (response model.ResponseCouponList) {
	count, coupons := s.Repository.FindCoupons(s.Db, req.UserId)

	response.Count = count
	response.Coupons = coupons

	return
}

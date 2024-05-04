package repository

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"web-cart-api/model"
)

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateUser(tx *gorm.DB, user model.User) (err error) {
	err = tx.Create(&user).Error
	return
}

func (r *repository) FindUserById(tx *gorm.DB, id int) (user model.User, err error) {
	err = tx.Where("id = ?", id).First(&user).Error
	return
}

func (r *repository) FindUserByUsername(tx *gorm.DB, username string) (user model.User, err error) {
	err = tx.Where("username = ?", username).First(&user).Error
	return
}

func (r *repository) FindProducts(tx *gorm.DB, name string, pagination model.QueryPagination) (count int64, products []model.Product) {
	tx = tx.Model(&model.Product{})

	if name != "" {
		name = "%" + name + "%"
		tx = tx.Where("LOWER(name) LIKE ?", name)
	}

	tx = tx.Count(&count).Limit(pagination.Limit).Offset(pagination.Offset)

	tx.Find(&products)
	fmt.Println(name)
	return
}

func (r *repository) FindProductById(tx *gorm.DB, id int) (product model.Product, err error) {
	err = tx.Where("id = ?", id).First(&product).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) FindCartByUserId(tx *gorm.DB, userId int) (count int64, carts []model.Cart) {
	tx.Model(&model.Cart{}).Where("user_id = ?", userId).Preload(clause.Associations).Count(&count).Find(&carts)
	return
}

func (r *repository) CreateProductToCart(tx *gorm.DB, cart model.Cart) (model.Cart, error) {
	err := tx.Create(&cart).Error
	return cart, err
}

func (r *repository) DeleteCart(tx *gorm.DB, cartId int) error {
	err := tx.Where("id = ?", cartId).Delete(&model.Cart{}).Error
	return err
}

func (r *repository) CreateOrder(tx *gorm.DB, order model.Order) (model.Order, error) {
	err := tx.Create(&order).Error
	return order, err
}

func (r *repository) FindOrderById(tx *gorm.DB, id int) (order model.Order, err error) {
	err = tx.Where("id = ?", id).Preload("OrderItems.Product").First(&order).Error
	return
}

func (r *repository) FindOrders(tx *gorm.DB, userId int) (count int64, orders []model.Order) {
	tx.Model(&model.Order{}).Preload("OrderItems.Product").Where("user_id = ?", userId).Order("id desc").Count(&count).Find(&orders)
	return
}

func (r *repository) UpdateOrder(tx *gorm.DB, id int, property map[string]interface{}) (order model.Order, err error) {
	err = tx.Model(&order).Clauses(clause.Returning{}).Where("id = ?", id).Updates(property).Error
	return
}

func (r *repository) FindCartById(tx *gorm.DB, id int, userId int) (cart model.Cart, err error) {
	err = tx.Where("id = ? AND user_id = ?", id, userId).Preload(clause.Associations).First(&cart).Error
	return
}

func (r *repository) FindCouponById(tx *gorm.DB, id int, userId int) (coupon model.Coupon, err error) {
	err = tx.Where("id = ? AND user_id = ?", id, userId).First(&coupon).Error
	return
}

func (r *repository) CreateCoupon(tx *gorm.DB, coupons []model.Coupon) (err error) {
	err = tx.Create(&coupons).Error
	return
}

func (r *repository) DeleteCarts(tx *gorm.DB, cartsId []int) error {
	err := tx.Where("id IN ?", cartsId).Delete(&model.Cart{}).Error
	return err
}

func (r *repository) DeleteCoupon(tx *gorm.DB, couponId int) error {
	err := tx.Where("id = ?", couponId).Delete(&model.Coupon{}).Error
	return err
}

func (r *repository) FindCoupons(tx *gorm.DB, userId int) (count int64, coupons []model.Coupon) {
	tx.Model(&model.Coupon{}).Where("user_id = ?", userId).Count(&count).Find(&coupons)
	return
}

package model

type RequestGetUserById struct {
	Id int `json:"id" uri:"id"`
}

type RequestSignup struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestGetProducts struct {
	Name string `json:"q"`
	RequestPagination
}

type RequestGetProductById struct {
	Id int `json:"id" uri:"id"`
}

type RequestGetCartByUserId struct {
	UserId int `json:"user_id" uri:"user_id"`
}

type QueryPagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type RequestPagination struct {
	Limit string `json:"limit"`
	Page  string `json:"page"`
}

type RequestAddProductToCart struct {
	UserId       int    `json:"user_id"`
	ProductId    int    `json:"product_id"`
	ProductCount int    `json:"product_count"`
	Note         string `json:"note"`
}

type RequestDeleteCart struct {
	Id     int `json:"id" uri:"id"`
	UserId int `json:"userId"`
}

type RequestOrder struct {
	UserId   int   `json:"user_id"`
	CartIds  []int `json:"cart_ids"`
	CouponId int   `json:"coupon_id"`
}

type RequestGetOrderById struct {
	Id int `json:"id" uri:"id"`
}

type RequestUpdateOrder struct {
	Id     int `json:"id" uri:"id"`
	UserId int
	Status string `json:"status"`
}

type RequestGetOrders struct {
	UserId int `json:"user_id" uri:"user_id"`
}

type RequestGetCoupons struct {
	UserId int `json:"user_id" uri:"user_id"`
}

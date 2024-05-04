package model

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseLogin struct {
	User
	Token string `json:"token"`
}

type ResponseProductList struct {
	Products []Product `json:"products"`
	Count    int64     `json:"count"`
}

type ResponseCartList struct {
	Carts []Cart `json:"carts"`
	Count int64  `json:"count"`
}

type ResponseOrderList struct {
	Orders []Order `json:"orders"`
	Count  int64   `json:"count"`
}

type ResponseCouponList struct {
	Coupons []Coupon `json:"coupons"`
	Count   int64    `json:"count"`
}

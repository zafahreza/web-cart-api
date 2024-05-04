package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-cart-api/helper"
	"web-cart-api/model"
	"web-cart-api/service"
)

type Handler struct {
	Service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) FetchUser(c *gin.Context) {

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	helper.ResponseSuccess(c, curretUser)
}

func (h *Handler) CreateUser(c *gin.Context) {

	var request model.RequestSignup

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.Service.CreateUser(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusInternalServerError, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	helper.ResponseSuccess(c, user)
}

func (h *Handler) GetUserById(c *gin.Context) {

	var request model.RequestGetUserById

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.Service.GetUserById(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusNotFound, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	helper.ResponseSuccess(c, user)
}

func (h *Handler) Login(c *gin.Context) {

	var request model.RequestLogin

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.Service.Login(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusNotFound, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	helper.ResponseSuccess(c, user)
}

func (h *Handler) GetProducts(c *gin.Context) {

	var request model.RequestGetProducts

	request.Limit = c.Query("limit")
	request.Page = c.Query("page")
	request.Name = c.Query("q")

	products := h.Service.GetProducts(request)

	helper.ResponseSuccess(c, products)
}

func (h *Handler) GetProductById(c *gin.Context) {

	var request model.RequestGetProductById

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	product, err := h.Service.GetProductById(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusNotFound, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	helper.ResponseSuccess(c, product)
}

func (h *Handler) GetCarts(c *gin.Context) {

	var request model.RequestGetCartByUserId

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	request.UserId = curretUser.Id

	carts := h.Service.GetCartByUserId(request)

	helper.ResponseSuccess(c, carts)
}

func (h *Handler) AddProductToCart(c *gin.Context) {

	var request model.RequestAddProductToCart

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	request.UserId = curretUser.Id

	err = h.Service.AddProductToCart(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusInternalServerError, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	helper.ResponseSuccess(c, nil)
}

func (h *Handler) DeleteCart(c *gin.Context) {

	var request model.RequestDeleteCart

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	request.UserId = curretUser.Id

	err = h.Service.DeleteCart(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusNotFound, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	helper.ResponseSuccess(c, nil)
}

func (h *Handler) CreateOrder(c *gin.Context) {

	var request model.RequestOrder

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	request.UserId = curretUser.Id

	order, err := h.Service.CreateOrder(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusInternalServerError, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	helper.ResponseSuccess(c, order)
}

func (h *Handler) GetOrderById(c *gin.Context) {

	var request model.RequestGetOrderById

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	order, err := h.Service.GetOrderById(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusNotFound, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	helper.ResponseSuccess(c, order)
}

func (h *Handler) GetOrders(c *gin.Context) {

	var request model.RequestGetOrders

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	request.UserId = curretUser.Id

	orders := h.Service.GetOrders(request)

	helper.ResponseSuccess(c, orders)
}

func (h *Handler) UpdateStatusOrder(c *gin.Context) {

	var request model.RequestUpdateOrder

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusUnprocessableEntity, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	request.UserId = curretUser.Id

	err = h.Service.UpdateStatusOrder(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ResponseFormater(http.StatusNotFound, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	helper.ResponseSuccess(c, nil)
}

func (h *Handler) GetCoupon(c *gin.Context) {

	var request model.RequestGetCoupons

	curretUser := c.MustGet("current_user").(model.User)
	if curretUser.Id < 1 {
		errorMessage := gin.H{"errors": "not identifier included"}

		response := helper.ResponseFormater(http.StatusUnauthorized, "error", errorMessage)

		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	request.UserId = curretUser.Id

	coupons := h.Service.GetCoupons(request)

	helper.ResponseSuccess(c, coupons)
}

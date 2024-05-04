package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-cart-api/model"
)

func ResponseFormater(statusCode int, message string, data interface{}) model.Response {
	return model.Response{
		Code:    statusCode,
		Status:  http.StatusText(statusCode),
		Message: message,
		Data:    data,
	}
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "success",
		Data:    data,
	})
}

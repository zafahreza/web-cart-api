package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"web-cart-api/handler"
	"web-cart-api/middleware"
	"web-cart-api/model"
	"web-cart-api/repository"
	"web-cart-api/service"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:        time.Second, // Slow SQL threshold
			LogLevel:             logger.Info, // Log level
			ParameterizedQueries: true,        // Don't include params in the SQL log
			Colorful:             true,        // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Product{}, &model.Coupon{}, &model.Order{}, &model.OrderItem{}, &model.Cart{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository()
	srvc := service.NewService(repo, db)
	han := handler.NewHandler(srvc)
	mid := middleware.NewAuthMiddleware()

	auth := mid.ValidateToken(srvc)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"authorization", "content-type"}
	router.Use(cors.New(corsConfig))

	router.POST("/users", han.CreateUser)
	router.GET("/users", auth, han.GetUserById)
	router.GET("/users/fetch", auth, han.FetchUser)
	router.POST("/login", han.Login)
	router.GET("/products", auth, han.GetProducts)
	router.GET("/products/:id", auth, han.GetProductById)
	router.GET("/carts", auth, han.GetCarts)
	router.POST("/carts", auth, han.AddProductToCart)
	router.DELETE("/carts/:id", auth, han.DeleteCart)
	router.POST("/orders", auth, han.CreateOrder)
	router.GET("/orders", auth, han.GetOrders)
	router.GET("/orders/:id", auth, han.GetOrderById)
	router.PATCH("/orders/:id", auth, han.UpdateStatusOrder)
	router.GET("/coupons", auth, han.GetCoupon)

	router.Run(":3000")
	return
}

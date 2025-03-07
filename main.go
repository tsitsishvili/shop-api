package main

import (
	"os"

	"github.com/tsitsishvili/shop-api/config"
	"github.com/tsitsishvili/shop-api/internal/handlers"
	"github.com/tsitsishvili/shop-api/internal/migrations"
	"github.com/tsitsishvili/shop-api/internal/repository"
	"github.com/tsitsishvili/shop-api/internal/service"
	"github.com/tsitsishvili/shop-api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := config.ConnectDB()
	migrations.Migrate(db)

	shopRepo := repository.NewShopRepository(db)
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	r := gin.Default()

	r.GET("/products", middlewares.APIKeyMiddleware(shopRepo), productHandler.FindAll)
	r.POST("/products", middlewares.APIKeyMiddleware(shopRepo), productHandler.Create)
	r.PUT("/products/:id", middlewares.APIKeyMiddleware(shopRepo), productHandler.Update)
	r.Run(":" + os.Getenv("PORT"))
}

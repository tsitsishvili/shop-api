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

	protected := r.Group("/")
	protected.Use(middlewares.APIKeyMiddleware(shopRepo))
	protected.GET("/products", productHandler.FindAll)
	protected.POST("/products", productHandler.Create)
	protected.PUT("/products/:id", productHandler.Update)

	r.Run(":" + os.Getenv("PORT"))
}

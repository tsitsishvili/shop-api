package repository

import (
	"github.com/tsitsishvili/shop-api/internal/models"

	"gorm.io/gorm"
)

type ShopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepository {
	return &ShopRepository{db}
}

func (r *ShopRepository) FindByAPIKey(apiKey string) (*models.Shop, error) {
	var shop models.Shop
	err := r.db.Where("api_key = ?", apiKey).First(&shop).Error
	return &shop, err
}

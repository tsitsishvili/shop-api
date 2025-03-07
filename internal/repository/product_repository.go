package repository

import (
	"github.com/tsitsishvili/shop-api/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) FindByIDAndShop(id uint, shopID uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ? AND shop_id = ?", id, shopID).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) FindAll(page, limit int, shopID uint) ([]models.Product, error) {
	var products []models.Product
	offset := (page - 1) * limit
	err := r.db.Where("shop_id = ?", shopID).Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}

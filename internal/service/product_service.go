package service

import (
	"errors"
	"os"

	"github.com/tsitsishvili/shop-api/internal/models"
	"github.com/tsitsishvili/shop-api/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) Update(product *models.Product, shopID uint) (*models.Product, error) {
	existing, err := s.repo.FindByIDAndShop(product.ID, shopID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if existing.Image != "" && product.Image != "" {
		os.Remove("./uploads/products/" + existing.Image)
	}

	if product.Image != "" {
		existing.Image = product.Image
	}

	existing.Title = product.Title
	existing.Price = product.Price
	existing.Quantity = product.Quantity

	err = s.repo.Update(existing)
	if err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *ProductService) FindAll(page, limit int, shopID uint) ([]models.Product, error) {
	return s.repo.FindAll(page, limit, shopID)
}

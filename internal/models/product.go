package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"` // We don't want to show DeletedAt in response
	Title     string         `json:"title" form:"title" validate:"required,min=3"`
	Barcode   string         `json:"barcode" form:"barcode" gorm:"unique" validate:"required,len=10"`
	Price     float64        `json:"price" form:"price" validate:"required,gt=0"`
	Quantity  int            `json:"quantity" form:"quantity" validate:"required,gt=0"`
	Image     string         `json:"image"`
	ShopID    uint           `json:"shop_id"`
}

func (p *Product) Validate() map[string]string {
	errors := make(map[string]string)
	v := validator.New()

	if err := v.Struct(p); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = "Invalid " + e.Tag()
		}
	}

	return errors
}

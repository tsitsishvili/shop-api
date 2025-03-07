package migrations

import (
	"fmt"

	"github.com/tsitsishvili/shop-api/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Shop{}, &models.Product{})
	fmt.Println("✅ Database Migrated")
}

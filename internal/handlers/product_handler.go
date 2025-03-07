package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tsitsishvili/shop-api/internal/models"
	"github.com/tsitsishvili/shop-api/internal/service"
	"github.com/tsitsishvili/shop-api/pkg/upload"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{s}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	file, _ := c.FormFile("image")
	if file != nil {
		imagePath, err := upload.UploadImage(file, "products")
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		product.Image = imagePath
	}

	product.ShopID = c.GetUint("shopID")

	errors := product.Validate()
	if len(errors) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
		return
	}

	err := h.service.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product Created",
		"product": product,
	})
}

func (h *ProductHandler) Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Product ID"})
		return
	}

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}

	file, _ := c.FormFile("image")
	if file != nil {
		imagePath, err := upload.UploadImage(file, "products")
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		product.Image = imagePath
	}

	product.ID = uint(productID)
	errors := product.Validate()
	if len(errors) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
		return
	}

	updatedProduct, err := h.service.Update(&product, c.GetUint("shopID"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Updated",
		"product": updatedProduct,
	})
}

func (h *ProductHandler) FindAll(c *gin.Context) {
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	products, err := h.service.FindAll(int(page), int(limit), c.GetUint("shopID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

package controller

import (
	"tokobelanja-kelompok7/service"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *productController {
	return &productController{productService}
}

func (h *productController) CreateProduct(c *gin.Context) {
	// TODO
}

func (h *productController) GetAllProducts(c *gin.Context) {
	// TODO
}

func (h *productController) UpdateProduct(c *gin.Context) {
	// TODO
}

func (h *productController) DeleteProduct(c *gin.Context) {
	// TODO
}

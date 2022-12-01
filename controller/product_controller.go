package controller

import (
	"net/http"
	"tokobelanja-kelompok7/helper"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/model/response"
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
	var inputBody input.ProductCreateInput

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindJSON(&inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	productData, err := h.productService.CreateProduct(role_user, inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	productResponse := response.ProductCreateResponse{
		ID:         productData.ID,
		Title:      productData.Title,
		Price:      productData.Price,
		Stock:      productData.Stock,
		CategoryID: productData.CategoryID,
		CreatedAt:  productData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			productResponse,
		),
	)
}

func (h *productController) GetAllProducts(c *gin.Context) {
	var allProducts []response.ProductGetResponse

	_ = c.MustGet("currentUser").(int)

	productsData, err := h.productService.GetAllProducts()
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	for _, product := range productsData {
		allProductsTmp := response.ProductGetResponse{
			ID:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt,
		}
		allProducts = append(allProducts, allProductsTmp)
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			allProducts,
		),
	)
}

func (h *productController) UpdateProduct(c *gin.Context) {
	// TODO
}

func (h *productController) DeleteProduct(c *gin.Context) {
	// TODO
}

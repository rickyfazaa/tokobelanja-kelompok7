package controller

import (
	"net/http"
	"strconv"
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
	var (
		inputBody input.ProductUpdateInput
		inputUri  input.ProductIdUri
	)

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

	err = c.ShouldBindUri(&inputUri)
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

	productData, err := h.productService.UpdateProduct(role_user, inputUri.ID, inputBody)
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

	productResponse := response.ProductUpdateResponse{
		Product: response.ProductUpdate{
			ID:         productData.ID,
			Title:      productData.Title,
			Price:      ToRupiah(productData.Price),
			Stock:      productData.Stock,
			CategoryID: productData.CategoryID,
			CreatedAt:  productData.CreatedAt,
			UpdatedAt:  productData.UpdatedAt,
		},
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			productResponse,
		),
	)
}

func (h *productController) DeleteProduct(c *gin.Context) {
	var inputUri input.ProductIdUri

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindUri(&inputUri)
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

	err = h.productService.DeleteProduct(role_user, inputUri.ID)
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

	productResponse := response.ProductDeleteResponse{
		Message: "Product has been successfully deleted",
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			productResponse,
		),
	)
}

func ToRupiah(price int) string {
	var rupiah string
	var num = []rune(strconv.Itoa(price))
	qty := len(num)
	for i := 0; i < qty; i++ {
		if i%3 == 0 && i != 0 {
			rupiah = string(num[qty-i-1]) + "." + rupiah
		} else {
			rupiah = string(num[qty-i-1]) + rupiah
		}
	}
	return "Rp. " + rupiah
}

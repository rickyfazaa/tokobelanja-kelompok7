package controller

import (
	"net/http"
	"tokobelanja-kelompok7/helper"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/model/response"
	"tokobelanja-kelompok7/service"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *categoryController {
	return &categoryController{categoryService}
}

func (h *categoryController) CreateCategory(c *gin.Context) {
	var input input.CategoryCreateInput

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindJSON(&input)
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

	categoryData, err := h.categoryService.CreateCategory(role_user, input)
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

	categoryResponse := response.CategoryCreateResponse{
		ID:                categoryData.ID,
		Type:              categoryData.Type,
		SoldProductAmount: categoryData.SoldProductAmount,
		CreatedAt:         categoryData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			categoryResponse,
		),
	)
}

func (h *categoryController) GetAllCategories(c *gin.Context) {
	var (
		allProducts   []response.CategoryProduct
		allCategories []response.CategoryGetResponse
	)

	_ = c.MustGet("currentUser").(int)

	categoryData, err := h.categoryService.GetAllCategories()
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

	for _, dataCategory := range categoryData {
		productsData, err := h.categoryService.GetProductsByCategoryID(dataCategory.ID)
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
		for _, dataProduct := range productsData {
			allProductsTmp := response.CategoryProduct{
				ID:        dataProduct.ID,
				Title:     dataProduct.Title,
				Price:     dataProduct.Price,
				Stock:     dataProduct.Stock,
				CreatedAt: dataProduct.CreatedAt,
				UpdatedAt: dataProduct.UpdatedAt,
			}
			allProducts = append(allProducts, allProductsTmp)
		}
		allCategoriesTmp := response.CategoryGetResponse{
			ID:                dataCategory.ID,
			Type:              dataCategory.Type,
			SoldProductAmount: dataCategory.SoldProductAmount,
			CreatedAt:         dataCategory.CreatedAt,
			UpdatedAt:         dataCategory.UpdatedAt,
			Product:           allProducts,
		}

		allCategories = append(allCategories, allCategoriesTmp)
		allProducts = []response.CategoryProduct{}
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			allCategories,
		),
	)
}

func (h *categoryController) PatchCategory(c *gin.Context) {
	var (
		inputBody input.CategoryPatchInput
		uri       input.CategoryIdUri
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

	err = c.ShouldBindUri(&uri)
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

	categoryData, err := h.categoryService.PatchCategory(role_user, uri.ID, inputBody)
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

	categoryResponse := response.CategoryPatchResponse{
		ID:                categoryData.ID,
		Type:              categoryData.Type,
		SoldProductAmount: categoryData.SoldProductAmount,
		UpdatedAt:         categoryData.UpdatedAt,
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			categoryResponse,
		),
	)
}

func (h *categoryController) DeleteCategory(c *gin.Context) {
	var uri input.CategoryIdUri

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindUri(&uri)
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

	err = h.categoryService.DeleteCategory(role_user, uri.ID)
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

	categoryResponse := response.CategoryDeleteResponse{
		Message: "Category has been successfully deleted",
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			categoryResponse,
		),
	)
}

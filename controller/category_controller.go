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
	// TODO
}

func (h *categoryController) PatchCategory(c *gin.Context) {
	// TODO
}

func (h *categoryController) DeleteCategory(c *gin.Context) {
	// TODO
}

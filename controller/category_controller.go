package controller

import (
	"project4/service"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *categoryController {
	return &categoryController{categoryService}
}

func (h *categoryController) CreateCategory(c *gin.Context) {
	// TODO
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

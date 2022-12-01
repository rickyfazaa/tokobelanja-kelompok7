package controller

import (
	"net/http"
	"tokobelanja-kelompok7/helper"
	"tokobelanja-kelompok7/model/input"
	"tokobelanja-kelompok7/model/response"
	"tokobelanja-kelompok7/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.UserRegisterInput

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

	userData, err := h.userService.RegisterUser(input)
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

	userResponse := response.UserRegisterResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Password:  userData.Password,
		Balance:   userData.Balance,
		CreatedAt: userData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			userResponse,
		),
	)
}

func (h *userController) LoginUser(c *gin.Context) {
	// TODO
}

func (h *userController) PatchTopUpUser(c *gin.Context) {
	// TODO
}

func (h *userController) RegisterAdmin(c *gin.Context) {
	// TODO
}

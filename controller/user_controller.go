package controller

import (
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
	// TODO
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

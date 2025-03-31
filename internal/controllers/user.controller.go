package controllers

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/sondoannam/go-ecommerce-backend-api/internal/services"
	"github.com/sondoannam/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 201, []string{"ok"})
}

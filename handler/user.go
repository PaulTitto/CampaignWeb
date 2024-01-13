package handler

import (
	"campaignweb/helper"
	"campaignweb/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// Catch Input from user
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidatorError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account has been Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Here you can continue processing the input, for example, by passing it to your service.
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse(
			"Register Account has been Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// For now, just sending an OK response.
	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse(
		"Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

package handler

import (
	"campaignweb/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUSers()
	if err != nil {
		//	letter
	}
	c.HTML(http.StatusOK, "user_index.html", gin.H{
		"users": users,
	})
}

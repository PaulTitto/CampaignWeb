package handler

import (
	"campaignweb/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler {
	return &sessionHandler{userService}
}

func (h *sessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}

func (h *sessionHandler) Create(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := h.userService.Login(input)
	if err != nil || user.Role != "admin" {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.Id)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/users")
}

func (h *sessionHandler) Destroy(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}

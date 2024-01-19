package main

import (
	"campaignweb/auth"
	"campaignweb/handler"
	"campaignweb/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/campaignweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authSevice := auth.NewService()
	authService.ValidateToken()
	userHandler := handler.NewUserHandler(userService, authSevice)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEamilAvailablity)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()

}

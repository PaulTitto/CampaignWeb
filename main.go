package main

import (
	"campaignweb/auth"
	"campaignweb/campaign"
	"campaignweb/handler"
	"campaignweb/helper"
	"campaignweb/payment"
	"campaignweb/transaction"
	"campaignweb/user"
	"gorm.io/driver/mysql"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	webHandler "campaignweb/web/handler"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigFile("config.env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	//databasePort := viper.GetString("DATABASE_PORT")
	//errEnv := godotenv.Load(".env")
	//if errEnv != nil {
	//	log.Fatal("Error loading .env file")
	//}
	//
	//conn := os.Getenv("POSTGRES_URL")
	//if conn == "" {
	//	log.Fatal("POSTGRES_URL is not set in the environment variables")
	//}
	//db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	//if err != nil {
	//	log.Fatalf("Error connecting to database: %s", err.Error())
	//}
	databasePort := viper.GetString("DATABASE_PORT")
	dsn := "root:@tcp(" + databasePort + ")/campaignweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	userWebHandler := webHandler.NewUserHandler(userService)
	campaignsWebHandler := webHandler.NewCampaignHandler(campaignService, userService)

	router := gin.Default()

	router.HTMLRender = loadTemplates("web/templates")

	allowedOrigins := viper.GetStringSlice("CORS_ALLOW_ORIGINS")
	configCORS := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(configCORS))

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; connect-src 'self' http://localhost:8080")
		c.Next()
	})

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEamilAvailablity)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdatedCampaign) // Corrected function name
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransactions) // Corrected function name

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	router.GET("/users", userWebHandler.Index)
	router.GET("/users/new", userWebHandler.New)
	router.GET("/users/avatar/:id", userWebHandler.NewAvatar)
	router.POST("/users/avatar/:id", userWebHandler.CreateAvatar)
	router.POST("/users", userWebHandler.Create)
	router.GET("/users/edit/:id", userWebHandler.Edit)
	router.POST("/users/update/:id", userWebHandler.Update)

	router.GET("/campaigns", campaignsWebHandler.Index)
	router.GET("/campaigns/new", campaignsWebHandler.New)
	router.POST("/campaigns", campaignsWebHandler.Create)
	router.GET("/campaigns/image/:id", campaignsWebHandler.NewImage)
	router.POST("/campaigns/image/:id", campaignsWebHandler.CreateImage)
	router.GET("/campaigns/edit/:id", campaignsWebHandler.Edit)
	router.POST("/campaigns/update/:id", campaignsWebHandler.Update)
	router.GET("/campaigns/show/:id", campaignsWebHandler.Show)

	serverPort := viper.GetString("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080" // Default port if not set
	}
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		log.Println("Loading template:", include)
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

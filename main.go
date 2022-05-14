package main

import (
	"golang-gin-poc/controllers"
	"golang-gin-poc/middlewares"
	"golang-gin-poc/repositories"
	"golang-gin-poc/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoRepository repositories.VideoRepository = repositories.NewVideoRepository()
	videoService    services.VideoService        = services.NewVideoService(videoRepository)
	loginService    services.LoginService        = services.NewLoginService()
	videoController controllers.VideoController  = controllers.New(videoService)
	jwtService      services.JWTService          = services.NewJWTService()
	loginController controllers.LoginController  = controllers.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()
	server := gin.New()

	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	server.Use(gin.Recovery(), gin.Logger())

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusCreated, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Create(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{"message": "Video created"})
			}
		})
		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusNoContent, gin.H{"message": "Video updated"})
			}
		})
		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusNoContent, gin.H{"message": "Video deleted"})
			}
		})

	}

	server.Run(":8080")
}

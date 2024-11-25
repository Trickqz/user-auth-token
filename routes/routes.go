package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trickqz/user-auth-token/controllers"
	"github.com/trickqz/user-auth-token/middlewares"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)

	r.GET("/profile", middlewares.AuthMiddleware(), controllers.GetUserProfile)
	r.PUT("/profile", middlewares.AuthMiddleware(), controllers.UpdateUserProfile)
}

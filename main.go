package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	database "github.com/trickqz/user-auth-token/config"
	"github.com/trickqz/user-auth-token/controllers"
	"github.com/trickqz/user-auth-token/routes"
)

func main() {
	database.ConnectBD()
	controllers.SetDB(database.DB)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Err loading .env")
	}

	r := gin.Default()
	routes.UserRoutes(r)
	r.Run()
}

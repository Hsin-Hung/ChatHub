package main

import (
	"api-server/controllers"
	"api-server/db"
	"api-server/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()
	r := gin.New()
	r.Use(middlewares.CorsMiddleware())
	r.GET("/", controllers.SignIn)
	r.POST("/signup", controllers.Register)
	r.POST("/signin", controllers.SignIn)
	r.GET("/user", controllers.CurrentUser)
	log.Fatal(r.Run("localhost:8080"))
}

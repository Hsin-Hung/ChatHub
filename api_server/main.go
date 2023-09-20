package main

import (
	"api-server/controllers"
	"api-server/db"
	"api-server/middlewares"
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB() // connect to database
	r := gin.New()
	r.Use(middlewares.CorsMiddleware())
	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	r.GET("/", controllers.SignIn)
	r.POST("/signup", controllers.Register)
	r.POST("/signin", controllers.SignIn)
	r.GET("/user", controllers.CurrentUser)
	log.Fatal(r.Run("0.0.0.0:8080"))
}

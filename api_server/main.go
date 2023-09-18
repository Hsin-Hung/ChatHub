package main

import (

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"api-server/db"

	"api-server/controllers"
)

func main() {
	db.Connect()
	r := gin.New()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/", controllers.SignIn)
	r.POST("/signup", controllers.Register)
	r.POST("/signin", controllers.SignIn)
	r.Run("localhost:8080") 
}

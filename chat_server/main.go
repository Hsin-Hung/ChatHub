// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"chat-server/chat"
	"chat-server/middlewares"
	"log"
	"net/http"

	"chat-server/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB() // connect to database
	db.CreateIndexes()
	db.ConnectRedis() // connect to redis pub/sub
	hub := chat.NewHub()
	go hub.Run()                                // run chat room
	go db.SubscribeMessage(hub.GetSubChannel()) // run redis message subscrib routine

	r := gin.Default()
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.JwtAuthMiddleware())
	r.GET("/ws", func(c *gin.Context) {
		chat.ServeWs(hub, c)
	})
	// get chat room info: client count
	r.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"online": hub.GetClientCount()})
	})
	log.Fatal(r.Run("localhost:8081"))

}

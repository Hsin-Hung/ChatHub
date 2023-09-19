// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"

	"chat-server/chat"
	"chat-server/middlewares"

	"chat-server/db"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8082", "http service address")

func main() {
	db.ConnectDB()
	db.ConnectRedis()
	flag.Parse()
	hub := chat.NewHub()
	go hub.Run()
	go db.SubscribeMessage(hub.GetSubChannel())

	r := gin.Default()
	r.Use(middlewares.JwtAuthMiddleware())
	r.GET("/ws", func(c *gin.Context) {
		chat.ServeWs(hub, c)
	})
	log.Fatal(r.Run(*addr))

}

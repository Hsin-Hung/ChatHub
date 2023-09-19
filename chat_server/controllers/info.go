package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChatRoomInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"username": username})
}

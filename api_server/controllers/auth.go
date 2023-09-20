package controllers

import (
	"api-server/db"
	"api-server/token"
	"api-server/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Get the hash of a password
func getHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Register a user
func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "registration failed"})
		return
	}

	username := input.Username
	password := input.Password

	if err := utils.UserInputValidator(username, password); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pass_hash, err := getHash([]byte(password))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "registration failed"})
		return
	}
	err = db.CreateNewUser(username, pass_hash)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

// Sign in a user
func SignIn(c *gin.Context) {

	var input SignInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect username or password"})
		return
	}

	username := input.Username
	password := input.Password

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect username or password"})
		return
	}

	dbPass, err := db.ValidateUser(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect username or password"})
		return
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(password))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect username or password"})
		return
	}

	// generate JWT token
	token, err := token.GenerateToken(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "sign in failed"})
		return
	}

	ws_uri, ok := os.LookupEnv("WS_URI")
	if !ok {
		ws_uri = "localhost:8081"
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "ws_uri": ws_uri})
}

func CurrentUser(c *gin.Context) {

	username, err := token.ExtractTokenID(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	username_sent := c.Query("username")
	if username == username_sent {
		c.JSON(http.StatusOK, gin.H{"message": "authorized"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

}

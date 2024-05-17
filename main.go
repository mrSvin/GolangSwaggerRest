package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	_ "swaggerRest/docs"
)

// @title User API
// @version 1.0
// @description This is a simple API to get user information
// @host localhost:8080
// @BasePath /

type User struct {
	ID   string `json:"id" example:"1"`
	Name string `json:"name" example:"John Doe"`
	Age  int    `json:"age" example:"30"`
}

var users = map[string]User{
	"1": {ID: "1", Name: "John Doe", Age: 30},
	"2": {ID: "2", Name: "Jane Smith", Age: 25},
}

// GetUserHandler handles the GET /users/{id} request
// @Summary Get user by ID
// @Description Returns a single user
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func main() {
	r := gin.Default()

	r.GET("/users/:id", GetUserHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

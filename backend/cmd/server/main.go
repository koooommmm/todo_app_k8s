package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo represents a todo item
type Todo struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{}

func setupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/todos", getTodos)
	}
	return router
}

func main() {
	router := setupRouter()
	router.Run()
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

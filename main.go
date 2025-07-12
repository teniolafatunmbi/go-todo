package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teniolafatunmbi/go-todo/internal/handlers"
	"net/http"
)

func main() {
	r := gin.Default();

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/todos", handlers.GetTodos);
	r.POST("/todos", handlers.AddTodo);
	r.PUT("/todos/:id", handlers.UpdateTodo);
	r.DELETE("/todos/:id", handlers.DeleteTodo);

	r.Run(":4000")
}

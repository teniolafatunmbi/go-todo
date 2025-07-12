package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Todo struct {
	ID int "json:id"
	Title string "json:title"
	IsCompleted bool "json:isCompleted"
}

type CreateTodo struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoStruct struct {
	Title       *string `json:"title,omitempty"`
	IsCompleted *bool   `json:"isCompleted,omitempty"`
}

var todos = []Todo{}


func getTodoById(todoId int) *Todo {
	for i, todo := range todos {
		if todo.ID == todoId {
			return &todos[i]
		}
	}
	return nil
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos);
}


func AddTodo(c *gin.Context) {
	var payload CreateTodo;

	if err := c.ShouldBindJSON((&payload)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo := Todo {
		ID: len(todos) + 1,
		Title: payload.Title,
		IsCompleted: false,
	}
	todos = append(todos, newTodo);
	
	response := map[string]any{
		"message": "Todo added successfully",
		"data":    newTodo,
	}
	c.JSON(http.StatusOK, response);
}

func UpdateTodo(c *gin.Context) {
	var payload UpdateTodoStruct
	todoId, err := strconv.Atoi(c.Param("id"));

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A valid Todo ID is required"});
		return
	}

	todo := getTodoById(todoId);

	if todo == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo with ID " + strconv.Itoa(todoId) + " does not exist"});
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}

	var updatedTodo Todo;

	for index, todo := range todos {
		if todo.ID == todoId {
			if payload.Title != nil && *payload.Title != "" {
				todos[index].Title = *payload.Title
			}

			if payload.IsCompleted != nil {
				todos[index].IsCompleted = *payload.IsCompleted
			}

			updatedTodo = todos[index];
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "data": updatedTodo})
}

func DeleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"));

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A valid Todo ID is required"});
		return;
	}

	todo := getTodoById(todoId);

	if todo == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo with ID " + strconv.Itoa(todoId) + " does not exist"});
		return
	}

	newTodos := []Todo{}
	for _, todo := range todos {
		if todo.ID != todoId {
			newTodos = append(newTodos, todo)
		}
	}

	todos = newTodos;

	c.JSON(http.StatusOK, gin.H{"message": "Todo with ID " + strconv.Itoa(todoId) + " deleted successfully"})
}

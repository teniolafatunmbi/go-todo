package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teniolafatunmbi/go-todo/internal/database"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	IsCompleted bool `json:"is_completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateTodo struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoStruct struct {
	Title       *string `json:"title,omitempty"`
	IsCompleted *bool   `json:"is_completed,omitempty"`
}

var ErrTodoNotFound = errors.New("todo not found");

func getTodoByID(id int) (*Todo, error) {
	const query = `
		SELECT id, title, is_completed, created_at, updated_at
		FROM todos
		WHERE id = $1
	`

	todo := &Todo{}
	err := database.Db.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.IsCompleted,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTodoNotFound
		}

		return nil, fmt.Errorf("getTodoByID query failed: %w", err)
	}

	return todo, nil
}


func GetTodos(c *gin.Context) {
	rows, err := database.Db.Query(`SELECT * FROM todos`);

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todos := []Todo{};
	for rows.Next() {
		var todo Todo
		var updatedAt sql.NullTime

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.IsCompleted, &todo.CreatedAt, &updatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if updatedAt.Valid {
			todo.UpdatedAt = &updatedAt.Time
		} else {
			todo.UpdatedAt = nil
		}
		
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, todos);
}


func AddTodo(c *gin.Context) {
	var payload CreateTodo;

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newTodo Todo;

	err := database.Db.QueryRow(
		`INSERT INTO todos (title, is_completed) VALUES ($1, $2) RETURNING id, title, is_completed, created_at, updated_at`, payload.Title, false).Scan(
			&newTodo.ID, &newTodo.Title, &newTodo.IsCompleted, &newTodo.CreatedAt, &newTodo.UpdatedAt);
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		
	
	response := map[string]any{
		"message": "Todo added successfully",
		"data": newTodo,
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

	_, getTodoErr := getTodoByID(todoId);

	if errors.Is(getTodoErr, ErrTodoNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}

	var updatedTodo Todo;

	setClauses := []string{}

	if payload.IsCompleted != nil {
		updateIsCompletedStmt := fmt.Sprintf("is_completed = %t", *payload.IsCompleted);
		setClauses = append(setClauses, updateIsCompletedStmt);
	}

	if payload.Title != nil {
		updateTitleStmt := fmt.Sprintf("title = '%s'", *payload.Title);
		setClauses = append(setClauses, updateTitleStmt);
	}

	updateTodoStmt := fmt.Sprintf("UPDATE todos SET %s, updated_at = NOW() WHERE id = $1 RETURNING id, title, is_completed, created_at, updated_at", strings.Join(setClauses, ", "));

	fmt.Println(updateTodoStmt);

	updateTodoErr := database.Db.QueryRow(updateTodoStmt, todoId).Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.IsCompleted, &updatedTodo.CreatedAt, &updatedTodo.UpdatedAt);

	if updateTodoErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateTodoErr.Error()});
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "data": updatedTodo})
}

func DeleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"));

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A valid Todo ID is required"});
		return;
	}

	_, getTodoErr := getTodoByID(todoId);

	if errors.Is(getTodoErr, ErrTodoNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	_, deleteTodoErr := database.Db.Exec("DELETE FROM todos WHERE id = $1", todoId);

	if deleteTodoErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": deleteTodoErr.Error()});
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo with ID " + strconv.Itoa(todoId) + " deleted successfully"})
}

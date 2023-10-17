package controllers

import (
	"net/http"
	"todo-app/pkg/models"
	"todo-app/pkg/services"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	TodoService *services.TodoService
}

func NewTodoController(ts *services.TodoService) *TodoController {
	return &TodoController{TodoService: ts}
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	todos, err := c.TodoService.GetTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) CreateATodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if err := c.TodoService.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	ctx.JSON(http.StatusCreated, todo)
}

// Implement other controller functions here

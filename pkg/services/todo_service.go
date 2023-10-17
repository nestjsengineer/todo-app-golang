package services

import (
	"todo-app/pkg/models"

	"gorm.io/gorm"
)

type TodoService struct {
	DB *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{DB: db}
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	result := s.DB.Create(todo)
	return result.Error
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := s.DB.Find(&todos)
	return todos, result.Error
}

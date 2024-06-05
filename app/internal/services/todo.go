package services

import (
	"ToDoAppGo/app/internal/models"
	"ToDoAppGo/app/internal/repositories/postgres"
	"gorm.io/gorm"
)

type TodoService struct {
	todoRepository *postgres.TodoRepository
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{
		todoRepository: postgres.NewTodoRepository(db),
	}
}

func (s TodoService) GetList() ([]models.Todo, error) {
	return s.todoRepository.GetAll()
}

func (s TodoService) GetById(id uint64) (*models.Todo, error) {
	return s.todoRepository.GetByID(id)
}

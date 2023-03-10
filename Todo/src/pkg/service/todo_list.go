package service

import (
	todoapp "gin-todo-prc/src/models"
	"gin-todo-prc/src/pkg/repo"
)

type TodoListService struct {
	repo repo.TodoList
}

func NewTodoListService(repo repo.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list todoapp.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]todoapp.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, id int) (todoapp.TodoList, error) {
	return s.repo.GetByID(userID, id)
}

func (s *TodoListService) Update(userID, id int, input *todoapp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, id, input)
}

func (s *TodoListService) Delete(userID, id int) error {
	return s.repo.Delete(userID, id)
}

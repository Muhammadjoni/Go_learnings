package service

import (
	todoapp "gin-todo-prc/src/models"
	"gin-todo-prc/src/pkg/repo"
)

type TodoItemService struct {
	repo     repo.TodoItem
	listRepo repo.TodoList
}

func NewTodoItemService(repo repo.TodoItem, listRepo repo.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userID, listID int, input todoapp.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil { // list doesn't exist or not belong to user
		return 0, nil
	}

	return s.repo.Create(userID, listID, input)
}

func (s *TodoItemService) GetAll(userID, listID int) ([]todoapp.TodoItem, error) {
	return s.repo.GetAll(userID, listID)
}

func (s *TodoItemService) GetByID(userID, itemID int) (todoapp.TodoItem, error) {
	return s.repo.GetByID(userID, itemID)
}

func (s *TodoItemService) Update(userID, itemID int, input todoapp.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userID, itemID, input)
}

func (s *TodoItemService) Delete(userID, itemID int) error {
	return s.repo.Delete(userID, itemID)
}

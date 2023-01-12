package service

import (
	restapi "REST_API"
	"REST_API/pkg/repository"
)

type TodoListService struct {
	repo repository.Todolist
}

func NewListService(repo repository.Todolist) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list restapi.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]restapi.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (restapi.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

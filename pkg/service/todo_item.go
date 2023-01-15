package service

import (
	restapi "REST_API"
	"REST_API/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.Todolist
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.Todolist) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item restapi.TodoItem) (int, error) {
	if _, err := s.listRepo.GetById(userId, listId); err != nil {
		// list repo does not exist or does not belong to user
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]restapi.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (restapi.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, item restapi.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, item)
}

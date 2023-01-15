package service

import (
	restapi "REST_API"
	"REST_API/pkg/repository"
)

type Authorization interface {
	CreateUser(user restapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Todolist interface {
	Create(userId int, list restapi.TodoList) (int, error)
	GetAll(userId int) ([]restapi.TodoList, error)
	GetById(userId, listId int) (restapi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input restapi.UpdateListInput) error
}

type TodoItem interface {
	Create(UserId, listId int, input restapi.TodoItem) (int, error)
	GetAll(UserId, listId int) ([]restapi.TodoItem, error)
	GetById(userId, itemId int) (restapi.TodoItem, error)
	Delete(userId, ItemId int) error
	Update(userId, listId int, input restapi.UpdateItemInput) error
}

type Service struct {
	Authorization
	Todolist
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Todolist:      NewListService(repos.Todolist),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.Todolist),
	}
}

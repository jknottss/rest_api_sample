package service

import (
	restapi "REST_API"
	"REST_API/pkg/repository"
)

type Authorization interface {
	CreateUser(user restapi.User) (int, error)
}

type Todolist interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	Todolist
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

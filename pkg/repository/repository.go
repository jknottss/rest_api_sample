package repository

import (
	restapi "REST_API"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user restapi.User) (int, error)
	GetUser(username, password string) (restapi.User, error)
}

type Todolist interface {
	Create(userId int, list restapi.TodoList) (int, error)
	GetAll(userId int) ([]restapi.TodoList, error)
	GetById(userId, listId int) (restapi.TodoList, error)
	Delete(userId, list int) error
	Update(userId, listId int, input restapi.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item restapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]restapi.TodoItem, error)
	GetById(userId, itemId int) (restapi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, item restapi.UpdateItemInput) error
}

type Repository struct {
	Authorization
	Todolist
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Todolist:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}

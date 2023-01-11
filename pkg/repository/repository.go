package repository

import (
	restapi "REST_API"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user restapi.User) (int, error)
}

type Todolist interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	Todolist
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

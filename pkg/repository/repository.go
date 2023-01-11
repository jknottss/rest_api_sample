package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}

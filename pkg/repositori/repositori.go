package repositori

import "database/sql"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repositori struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepositori(db *sql.DB) *Repositori {
	return &Repositori{}
}

package repositori

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

func NewRepositori() *Repositori {
	return &Repositori{}
}

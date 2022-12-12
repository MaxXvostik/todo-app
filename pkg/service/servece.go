package service

import "github.com/maXvostik/todo-app/pkg/repositori"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repositori.Repositori) *Service {
	return &Service{}
}

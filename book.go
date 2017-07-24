package main

type Book struct {
	Id string
	Name string
	Author string
}

func NewBook(id, name, author string) *Book {
	return &Book{id, name, author}
}

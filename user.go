package main

import (
	"errors"
)

const MAX_BOOK_ALLOTED = 2

type User struct {
	Id string
	Name string
	Books map[string]*Book
}

func NewUser(id, name string) *User {
	u := &User{Id:id, Name:name}
	u.Books = make(map[string]*Book)

	return u
}

func (u *User) AllotBook(b *Book) error {
	if len(u.Books) == MAX_BOOK_ALLOTED {
		return errors.New("Book Limit Reached")
	}

	u.Books[b.Id] = b
	return nil
}

func (u *User) ReturnBook(b *Book) {
	delete(u.Books, b.Id)
}

package main

import (
	"strings"
	"errors"
	"crypto/md5"
	"encoding/hex"
)

type Library struct {
	Name string
	Books []*Book
	Users []*User

	BookUserMap map[string]string

	BookSearch map[string]*Book
	UserSearch map[string]*User
}

func NewLibrary(name string) *Library {
	l :=  &Library{Name: name}

	l.Books = make([]*Book, 0, 10)
	l.Users = make([]*User, 0, 10)	

	l.BookUserMap = make(map[string]string)

	return l
}

func (l *Library) AddBook(b *Book) {
	l.Books = append(l.Books, b)

	names := strings.Fields(b.Name)
	for _, n := range(names) {
		hash := GetMD5Hash(n)
		l.BookSearch[hash] = b
	}

	authors := strings.Fields(b.Author)
        for _, a := range(authors) {
                hash := GetMD5Hash(a)
                l.BookSearch[hash] = b
        }
}

func (l *Library) AddUser(u *User) {
	l.Users = append(l.Users, u)
	
	names := strings.Fields(u.Name)
        for _, n := range(names) {
                hash := GetMD5Hash(n)
                l.UserSearch[hash] = u
        }
}

func (l *Library) Lend(b *Book, u *User) error {
	if _, ok := l.BookUserMap[b.Id]; ok {
		return errors.New("Book already lent")
	}
	
	err := u.AllotBook(b)
	if err != nil {
		return err
	}

	l.BookUserMap[b.Id] = u.Id
	return nil
}

func (l *Library) Return(b *Book, u *User) {
	delete(l.BookUserMap,b.Id)
}

func (l *Library) SearchBooks(query string) *Book {
	hash := GetMD5Hash(query)
	b, found := l.BookSearch[hash]
	if !found {
		return nil
	}
	return b
}

func (l *Library) SearchUser(query string) *User {
	hash := GetMD5Hash(query)
        u, found := l.UserSearch[hash]
        if !found {
                return nil
        }
        return u
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

package main

import (
	"fmt"
	"go-ws-try-me/internal/server"
	"time"
)

type Book struct {
	title    string
	author   string
	numPages int32

	isSaved bool
	savedAt time.Time
}

func (book *Book) SaveBook() {
	book.isSaved = true
	book.savedAt = time.Now()
}

func SaveBook(book *Book) {
	book.isSaved = true
	book.savedAt = time.Now()
}

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

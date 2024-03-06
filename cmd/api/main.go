package main

import (
	"fmt"
	"go-ws-try-me/internal/server"
	"math/rand"
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

type FootballPlayer struct {
	stamina int
	power   int
}

type Player interface {
	KickBall()
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("Kicking the ball ", shot)
}

func main() {

	team := make([]Player, 11)
	for i := 0; i < len(team); i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
		team[i].KickBall()
	}

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

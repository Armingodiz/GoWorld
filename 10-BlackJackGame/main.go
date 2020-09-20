package main

import (
	"fmt"
	"github.com/Armingodiz/GoWebDewTraining/10-BlackJackGame/deck"
)

type Dealer struct {
	Cards []deck.Card
}
type Player struct {
	NickName string
	Cards []deck.Card
}
type User interface {
	userTurn()
}
var cards []deck.Card

func main() {
	cards = deck.NewDeck()
	dealer := Dealer{[]deck.Card{}}
	player := Player{"armin",[]deck.Card{}}
	users:= []User{}
	users = append(users, &dealer)
	users = append(users, &player)
	Start(users)
}

func Start(users []User) {
	for _,user := range users{
		user.userTurn()
	}
}
func (dealer *Dealer) userTurn() {
	fmt.Println("DEALER TURN : ")
}
func (player *Player) userTurn() {
	fmt.Println("PLAYER "+ player.NickName+" TURN : ")
}

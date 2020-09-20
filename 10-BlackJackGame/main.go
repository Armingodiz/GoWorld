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
	Cards    []deck.Card
}
type User interface {
	userTurn()
	play()
}

var cards []deck.Card
var holder = 0
var dealer Dealer

func main() {
	cards = deck.NewDeck(deck.MultipleDeck(2), deck.Shuffle)
	fmt.Println(cards)
	dealer = Dealer{[]deck.Card{}}
	player := Player{"armin", []deck.Card{}}
	users := []User{}
	users = append(users, &dealer)
	users = append(users, &player)
	Start(users)
}

func Start(users []User) {
	for _, user := range users {
		user.userTurn()
	}
}
func (dealer *Dealer) userTurn() {
	fmt.Println("DEALER TURN : ")
	dealer.Cards = append(dealer.Cards, cards[holder], cards[holder+1])
	holder += 2
	fmt.Println(dealer)
}
func (player *Player) userTurn() {
	fmt.Println("PLAYER " + player.NickName + " TURN : ")
	player.Cards = append(player.Cards, cards[holder], cards[holder+1])
	holder += 2
	fmt.Println(player)
	fmt.Println("DEALER CARD IS : " + dealer.Cards[0].Rank.String() + " OF " + dealer.Cards[0].Suit.String() + "s")
	player.play()
}
func (player *Player) play() {
	fmt.Println("1 ) HIT \n2 ) STAND ")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		player.Cards = append(player.Cards, cards[holder], cards[holder+1])
		holder += 1
		player.play()
	case 2:
		break
	}
}
func (dealer *Dealer) play() {
	fmt.Println("1 ) HIT \n2 ) STAND ")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
	case 2:
		break
	}
}
func scoring(cards []deck.Card) int {
	var score = 0
	return score
}

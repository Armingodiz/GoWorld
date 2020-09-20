package main

import (
	"fmt"
	"github.com/Armingodiz/GoWebDewTraining/10-BlackJackGame/deck"
)

type Dealer struct {
	Cards  []deck.Card
	Status int // 0 for stand , 1 for win , -1  for lose
}
type Player struct {
	NickName string
	Status   int // 0 for stand , 1 for win , -1  for lose
	Cards    []deck.Card
}
type User interface {
	userTurn() int
	play() int
}

var cards []deck.Card
var holder = 0
var dealer Dealer

func main() {
	cards = deck.NewDeck(deck.MultipleDeck(2), deck.Shuffle)
	fmt.Println(cards)
	dealer = Dealer{[]deck.Card{}, 0}
	dealer.Cards = append(dealer.Cards, cards[holder], cards[holder+1])
	holder += 2
	player := Player{"armin", 0, []deck.Card{}}
	users := []User{}
	users = append(users, &player)
	Start(users)
}

func Start(users []User) {
	var status = 0
	for _, user := range users {
		status += user.userTurn()
		if status == 1 {
			break
		}
	}
	if status == 0 || (status < 0 && status*-1 != len(users)) {
		dealer.userTurn()
	}
}
func scoring(cards []deck.Card) (int, string) {
	var score = 0
	var Type = "normal"
	return score, Type
}

///////////////////////////////////////////////////////////////////////////////////////////////////// DEALER PART :
func (dealer *Dealer) userTurn() int {
	fmt.Println("DEALER TURN : ")
	fmt.Println(dealer)
	return dealer.play()
}
func (dealer *Dealer) play() int {
	fmt.Println("1 ) HIT \n2 ) STAND ")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		score, Type := scoring(dealer.Cards)
		if score < 16 || (score == 17 && Type == "soft") {
			dealer.Cards = append(dealer.Cards, cards[holder], cards[holder+1])
			holder += 1
			score, _ := scoring(dealer.Cards)
			if score == 21 {
				fmt.Println("DEALER WON !")
				return 1
			} else if score > 21 {
				fmt.Println("DEALER LOST !")
				return -1
			} else {
				dealer.play()
			}
		} else {
			fmt.Println("YOU CANT HIT !")
			break
		}
	case 2:
		break
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////// DEALER PART :
func (player *Player) userTurn() int {
	fmt.Println("PLAYER " + player.NickName + " TURN : ")
	player.Cards = append(player.Cards, cards[holder], cards[holder+1])
	holder += 2
	fmt.Println(player)
	fmt.Println("DEALER CARD IS : " + dealer.Cards[0].Rank.String() + " OF " + dealer.Cards[0].Suit.String() + "s")
	return player.play()
}
func (player *Player) play() int { // 0 for stand , 1 for win , -1  for lose
	fmt.Println("1 ) HIT \n2 ) STAND ")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		player.Cards = append(player.Cards, cards[holder], cards[holder+1])
		holder += 1
		score, _ := scoring(player.Cards)
		if score == 21 {
			fmt.Println("PLAYER " + player.NickName + " WON !")
			return 1
		} else if score > 21 {
			fmt.Println("PLAYER " + player.NickName + " LOST !")
			return -1
		} else {
			player.play()
		}
	case 2:
		break
	default:
		fmt.Println("INVALID INPUT !")
		player.play()
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

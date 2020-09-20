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
	getScore() int
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
	fmt.Println(status)
	if status == 0 || (status < 0 && status*-1 != len(users)) {
		status = dealer.userTurn()
		if status == 0 {
			max, _ := scoring(dealer.Cards)
			maxIndex := -1
			for i, user := range users {
				if user.getScore() > max {
					max = user.getScore()
					maxIndex = i
				}
			}
			if maxIndex == -1 {
				fmt.Println("DEALER WON !")
			} else {
				fmt.Println("WINNER IS :")
				fmt.Println(users[maxIndex])
			}
		}
	}
}
func (player *Player) getScore() int {
	score, _ := scoring(player.Cards)
	return score
}
func (dealer *Dealer) getScore() int {
	score, _ := scoring(dealer.Cards)
	return score
}
func scoring(deckCard []deck.Card) (int, string) {
	var score = 0
	var Type = "normal"
	for _, card := range deckCard {
		if card.Rank == deck.King || card.Rank == deck.Quin || card.Rank == deck.Jack {
			score += 10
		} else if card.Rank == deck.Ace {
			score += 1
			Type = "ace"
		} else {
			score += int(card.Rank)
		}
	}
	if Type == "ace" {
		if score == 7 {
			score = 17
			Type = "soft"
		} else if score <= 11 {
			score += 10
		}
	}
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
	var code = 0
	fmt.Scan(&input)
	switch input {
	case 1:
		score, Type := scoring(dealer.Cards)
		if score < 16 || (score == 17 && Type == "soft") {
			dealer.Cards = append(dealer.Cards, cards[holder])
			fmt.Println("NEW CARD : " + cards[holder].Rank.String() + " OF " + cards[holder].Suit.String())
			holder += 1
			score, _ := scoring(dealer.Cards)
			if score == 21 {
				fmt.Println("DEALER WON !")
				return 1
			} else if score > 21 {
				fmt.Println("DEALER LOST !")
				return -1
			} else {
				code = dealer.play()
			}
		} else {
			fmt.Println("YOU CANT HIT !")
			break
		}
	case 2:
		break
	}
	return code
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
	var code = 0
	fmt.Println("1 ) HIT \n2 ) STAND ")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		player.Cards = append(player.Cards, cards[holder])
		fmt.Println("NEW CARD : " + cards[holder].Rank.String() + " OF " + cards[holder].Suit.String())
		holder += 1
		score, _ := scoring(player.Cards)
		if score == 21 {
			fmt.Println("PLAYER " + player.NickName + " WON !")
			return 1
		} else if score > 21 {
			fmt.Println("PLAYER " + player.NickName + " LOST !")
			return -1
		} else {
			code = player.play()
		}
	case 2:
		break
	default:
		fmt.Println("INVALID INPUT !")
		player.play()
	}
	return code
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

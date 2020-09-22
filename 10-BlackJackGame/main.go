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
	resetCards()
}

var cards []deck.Card
var holder = 0
var dealer Dealer
var users = []User{}

func main() {
	cards = deck.NewDeck(deck.MultipleDeck(3), deck.Shuffle)
	for {
		fmt.Println("1 ) NEW GAME \n2 ) PLAY AGAIN ")
		var opt int
		fmt.Scan(&opt)
		switch opt {
		case 1:
			start()
		case 2:
			if len(users) != 0 {
				playAgain()
			} else {
				fmt.Println("YOU DONT HAVE A OPEN GAME !")
			}
		default:
			fmt.Println("INVALID INPUT !")
		}
	}
}
func start() {
	dealer = Dealer{[]deck.Card{}, 0}
	dealer.Cards = append(dealer.Cards, cards[holder], cards[holder+1])
	holder += 2
	fmt.Println("ENTER NUMBER OF PLAYERS : ")
	users := []User{}
	var n int
	var name string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println("ENTER NICK NAME FOR PLAYER : ")
		fmt.Scan(&name)
		player := Player{name, 0, []deck.Card{}}
		users = append(users, &player)
	}
	game(users)
}
func playAgain() {
	for _, user := range users {
		user.resetCards()
	}
	game(users)
}
func game(users []User) {
	var status = 0
	for _, user := range users {
		status += user.userTurn()
		if status == 1 {
			break
		}
	}
	if status == 0 || (status < 0 && status*-1 != len(users)) {
		status = dealer.userTurn()
		max, _ := scoring(dealer.Cards)
		maxIndex := -1
		for i, user := range users {
			score := user.getScore()
			if score > max && score <= 21 {
				max = score
				maxIndex = i
			}
		}
		if maxIndex == -1 {
			fmt.Println("DEALER WON !")
			fmt.Print("CARDS : ")
			for _, card := range dealer.Cards {
				fmt.Print("***  " + card.Rank.String() + " OF " + card.Suit.String() + "s" + "  ***")
			}
		} else {
			fmt.Println("WINNER IS :")
			fmt.Println(users[maxIndex])
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
	return dealer.play()
}
func (dealer *Dealer) play() int {
	var input int
	var code = 0
	score, Type := scoring(dealer.Cards)
	input = dealerAI(score, Type)
	switch input {
	case 1:
		dealer.Cards = append(dealer.Cards, cards[holder])
		fmt.Println("NEW CARD : " + cards[holder].Rank.String() + " OF " + cards[holder].Suit.String())
		holder += 1
		score, _ := scoring(dealer.Cards)
		if score == 21 {
			fmt.Println("DEALER WON !")
			fmt.Print("CARDS : ")
			for _, card := range dealer.Cards {
				fmt.Print("***  " + card.Rank.String() + " OF " + card.Suit.String() + "s" + "  ***")
			}
			return 1
		} else if score > 21 {
			fmt.Println("DEALER LOST !")
			fmt.Print("CARDS : ")
			for _, card := range dealer.Cards {
				fmt.Print("***  " + card.Rank.String() + " OF " + card.Suit.String() + "s" + "  ***")
			}
			return -1
		} else {
			code = dealer.play()
		}
	case 2:
		fmt.Println("DEALER STANDS !")
		break
	}
	return code
}
func dealerAI(score int, Type string) int {
	opt := 0
	if score < 16 || (score == 17 && Type == "soft") {
		opt = 1
	} else {
		opt = 2
	}
	return opt
}
func (dealer *Dealer) resetCards() {
	dealer.Cards = []deck.Card{}
	dealer.Cards = append(dealer.Cards, cards[holder], cards[holder+1])
	holder += 2
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////// DEALER PART :
func (player *Player) userTurn() int {
	fmt.Println("PLAYER " + player.NickName + " TURN : ")
	player.Cards = append(player.Cards, cards[holder], cards[holder+1])
	holder += 2
	fmt.Println("DEALER CARD IS : " + dealer.Cards[0].Rank.String() + " OF " + dealer.Cards[0].Suit.String() + "s")
	return player.play()
}
func (player *Player) play() int { // 0 for stand , 1 for win , -1  for lose
	fmt.Print("YOUR CARDS : ")
	for _, card := range player.Cards {
		fmt.Print("***  " + card.Rank.String() + " OF " + card.Suit.String() + "s" + "  ***")
	}
	fmt.Println()
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
func (player *Player) resetCards() {
	player.Cards = []deck.Card{}
	player.Cards = append(player.Cards, cards[holder], cards[holder+1])
	holder += 2
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

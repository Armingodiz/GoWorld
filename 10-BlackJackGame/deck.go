//go:generate stringer -type=Suit,Rank
// typing "go generate" in terminal after typing above line
// and stringer do the rest
package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Suit int8

const (
	// iota get the first value 0 and 1 to second value and ...
	Spade Suit = iota
	Club
	Diamond
	Heart
	Joker
)

var suites = [...]Suit{Spade, Club, Diamond, Heart}

type Rank int

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Quin
	King
)
const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (card Card) string() string {
	return card.Rank.String() + " of " + card.Suit.String() + "s"
}

func NewDeck(opts ...func(cards []Card) []Card) []Card {
	var deck []Card
	for _, suit := range suites {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{suit, rank})
		}
	}
	for _, opt := range opts {
		deck = opt(deck)
	}
	return deck
}
func Shuffle(cards []Card) []Card {
	shuffled := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	parm := r.Perm(len(cards))
	for i, j := range parm {
		shuffled[i] = cards[j]
	}
	return shuffled
}
func Sort(Less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, Less(cards))
		return cards
	}
}
func Filter(condition func(card Card) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var filtered []Card
		for _, card := range cards {
			if !condition(card) {
				filtered = append(filtered, card)
			}
		}
		return filtered
	}
}

func MultipleDeck(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var deck []Card
		for i := 0; i < n; i++ {
			deck = append(deck, cards...)
		}
		return deck
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return totalRank(cards[i]) < totalRank(cards[j])
	}
}
func More(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return totalRank(cards[i]) > totalRank(cards[j])
	}
}
func AddJokers(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}
func totalRank(card Card) int {
	return int(card.Suit)*int(maxRank) + int(card.Rank)
}

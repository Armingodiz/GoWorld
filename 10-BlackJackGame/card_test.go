package deck

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var card1 = Card{Spade, Ace}
	var card2 = Card{Heart, King}
	fmt.Println(card1.string())
	fmt.Println(card2.string())
}
func TestNew(t *testing.T) {
	deck := NewDeck()
	fmt.Println(deck)
	fmt.Println("############################################")
}
func TestLess(t *testing.T) {
	deck := NewDeck(Sort(Less))
	fmt.Println(deck)
	fmt.Println("############################################")
}
func TestMore(t *testing.T) {
	deck := NewDeck(Sort(More))
	fmt.Println(deck)
	fmt.Println("############################################")
}
func TestShuffle(t *testing.T) {
	deck := NewDeck(Shuffle)
	fmt.Println(deck)
	fmt.Println("############################################")
}
func TestAddJokers(t *testing.T) {
	deck := NewDeck(AddJokers(5))
	fmt.Println(deck)
	fmt.Println("############################################")
}

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
	if len(deck)!= 13*4{
		t.Errorf("WRONG NUMBER OF CARDS !")
	}
}
func TestLess(t *testing.T) {
	deck := NewDeck(Sort(Less))
	first := Card{
		Suit: Spade,
		Rank: minRank,
	}
	last := Card{
		Suit: Heart,
		Rank: maxRank,
	}
	if deck[0]!=first || deck[51]!=last{
		t.Errorf("WRONG Order !")
	}
}
func TestMore(t *testing.T) {
	deck := NewDeck(Sort(Less))
	first := Card{
		Suit: Spade,
		Rank: minRank,
	}
	last := Card{
		Suit: Heart,
		Rank: maxRank,
	}
	if deck[0]!=last || deck[51]!=first{
		t.Errorf("WRONG Order !")
	}
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
func TestFilter(t *testing.T) {
	condition := func(card Card) bool {
		return card.Rank == minRank
	}
	deck := NewDeck(Filter(condition))
	fmt.Println(deck)
	fmt.Println("############################################")
}
func TestMultipleDeck(t *testing.T) {
	deck := NewDeck(MultipleDeck(2))
	fmt.Println(deck)
	fmt.Println("############################################")
}

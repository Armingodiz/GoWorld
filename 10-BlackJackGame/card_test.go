package deck

import (
	"testing"
)

func TestString(t *testing.T) {
	var card1 = Card{Spade, Ace}
	var card2 = Card{Heart, King}
	if card1.string()!="Ace of Spades" || card2.string()!="King of Hearts"{
		t.Errorf("STRING TEST FAILED ! ")
	}

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
		t.Errorf("WRONG Order LESS!")
	}
}
func TestMore(t *testing.T) {
	deck := NewDeck(Sort(More))
	first := Card{
		Suit: Spade,
		Rank: minRank,
	}
	last := Card{
		Suit: Heart,
		Rank: maxRank,
	}
	if deck[0]!=last || deck[51]!=first{
		t.Errorf("WRONG Order MORE  !")
	}
}
func TestShuffle(t *testing.T) {
	deck := NewDeck(Shuffle)
	reg := NewDeck()
	if reg[0]==deck[0] && reg[1]==deck[1]&& reg[51]==deck[51]{
		t.Errorf("NOT SHUFFLED ! ")
	}
}
func TestAddJokers(t *testing.T) {
	deck := NewDeck(AddJokers(5))
	counter := 0
	for _,card := range deck{
		if card.Suit==Joker{
			counter+=1
		}
	}
	if counter!=5{
		t.Errorf("ADD JOKER TEST FAILED !")
	}
}
func TestFilter(t *testing.T) {
	condition := func(card Card) bool {
		return card.Rank == minRank
	}
	deck := NewDeck(Filter(condition))
	for _,card := range deck{
		if card.Rank==minRank{
			t.Errorf("FILTER TEST FAILED ! ")
		}
	}

}
func TestMultipleDeck(t *testing.T) {
	deck := NewDeck(MultipleDeck(2))
	if len(deck)!=104{
		t.Errorf("MultipleDeck TEST FAILED ! ")
	}
}

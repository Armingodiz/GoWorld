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

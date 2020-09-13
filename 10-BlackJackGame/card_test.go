package deck

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var card = Card{Spade, Ace}
	fmt.Println(card.string())
}

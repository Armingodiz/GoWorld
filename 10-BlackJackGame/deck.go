//go:generate stringer -type=Suit,Rank
// typing "go generate" in terminal after typing above line
// and stringer do the rest
package deck

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
func NewDeck() []Card {
	var deck []Card
	for _, suit := range suites {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{suit, rank})
		}
	}
	return deck
}

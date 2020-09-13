//go:generate stringer -type=Suit,Rank
package deck

type Suit int8

const (
	Spade Suit = iota
	Club
	Diamond
	Heart
)

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

type Card struct {
	Suit
	Rank
}

func (card Card) string() string {
	return card.Rank.String() + " OF " + card.Suit.String() + "s"
}

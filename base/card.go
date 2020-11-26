package base

// Rank is the rank of a card, for human usage
type Rank int

const (
	// Two is two
	Two Rank = iota + 1
	// Three is three
	Three
	// Four is four
	Four
	// Five is five
	Five
	// Six is six
	Six
	// Seven is seven
	Seven
	// Eight is eight
	Eight
	// Nine is nine
	Nine
	// Ten is ten
	Ten
	// Jack is jack
	Jack
	// Queen is queen
	Queen
	// King is king
	King
	// Ace is ace
	Ace
)

// Suit is the suit of a card, for human usage
type Suit int

const (
	// Club is club
	Club Suit = iota + 1
	// Diamond is diamond
	Diamond
	// Heart is heart
	Heart
	// Spade is spade
	Spade
)

// Card is a card
// This is a 64 bit int because that's what we're going to use everywhere, but a card is guaranteed to use
// only the lowest 8 bits
// Guaranteed not to be 0
type Card uint64

// NewCard returns a new card
func NewCard(rank Rank, suit Suit) Card {
	// Lowest 4 bits are for the suit, next 4 bits are for the rank
	return Card((uint64(rank) << 4) | uint64(suit))

}

// GetRank returns the card's rank
func (c Card) GetRank() Rank {
	return Rank(c >> 4)
}

// GetSuit returns the card's suit
func (c Card) GetSuit() Suit {
	return Suit(c & 0xf)
}

func (c Card) String() string {
	rank := c.GetRank()
	suit := c.GetSuit()

	var r, s string
	switch rank {
	case Two:
		r = "2"
	case Three:
		r = "3"
	case Four:
		r = "4"
	case Five:
		r = "5"
	case Six:
		r = "6"
	case Seven:
		r = "7"
	case Eight:
		r = "8"
	case Nine:
		r = "9"
	case Ten:
		r = "T"
	case Jack:
		r = "J"
	case Queen:
		r = "Q"
	case King:
		r = "K"
	case Ace:
		r = "A"
	default:
		return "(invalid)"
	}
	switch suit {
	case Club:
		s = "C"
	case Diamond:
		s = "D"
	case Heart:
		s = "H"
	case Spade:
		s = "S"
	default:
		return "(invalid)"
	}
	return "(" + r + "," + s + ")"
}

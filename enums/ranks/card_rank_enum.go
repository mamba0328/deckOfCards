//go:generate stringer -type=CardRank

package ranks

type CardRank byte

const (
	Ace CardRank = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	Jack
	Queen
	King
	Joker
)

//go:generate stringer -type=CardSuit

package suits

type CardSuit byte

const (
	Spade CardSuit = iota + 1
	Diamond
	Heart
	Club
	Joker
)

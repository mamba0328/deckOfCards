package card

import (
	"fmt"

	ranks "deck.com/enums/ranks"
	suits "deck.com/enums/suits"
)

type Card struct {
	Rank ranks.CardRank
	Suit suits.CardSuit
}

func (c Card) String() string {
	if c.Rank == ranks.Joker {
		return ranks.Joker.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

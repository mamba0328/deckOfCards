package card

import (
	"fmt"

	"deck.com/enums/ranks"
	"deck.com/enums/suits"
)

func ExampleCard() { //https://go.dev/blog/examples
	fmt.Println(Card{Suit: suits.Club, Rank: ranks.Ace})
	fmt.Println(Card{Suit: suits.Joker, Rank: ranks.Joker})

	// Output:
	// Ace of Clubs
	// Joker
}

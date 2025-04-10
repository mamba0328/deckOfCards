package card

import (
	"fmt"

	"github.com/mamba0328/deckOfCards/enums/ranks"
	"github.com/mamba0328/deckOfCards/enums/suits"
)

func ExampleCard() { //https://go.dev/blog/examples
	fmt.Println(Card{Suit: suits.Club, Rank: ranks.Ace})
	fmt.Println(Card{Suit: suits.Joker, Rank: ranks.Joker})

	// Output:
	// Ace of Clubs
	// Joker
}

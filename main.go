package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"

	"github.com/mamba0328/deckOfCards/card"

	"github.com/mamba0328/deckOfCards/enums/ranks"
	"github.com/mamba0328/deckOfCards/enums/suits"
)

// [...] instead of []: ensures you get a (fixed size) array instead of a slice. So the values aren't fixed but the size is.
var rankOrder = [...]ranks.CardRank{ranks.Ace, ranks.TWO, ranks.THREE, ranks.FOUR, ranks.FIVE, ranks.SIX, ranks.SEVEN, ranks.EIGHT, ranks.NINE, ranks.TEN, ranks.Jack, ranks.Queen, ranks.King}
var suitsOrder = [...]suits.CardSuit{suits.Spade, suits.Diamond, suits.Heart, suits.Club}

type FunctionalOption func(*[]card.Card)

func main() {
	deck := New(WithShuffle(), WithSort(DefaultCompare))

	for _, v := range deck {
		fmt.Println(v)
	}
}

func New(options ...FunctionalOption) []card.Card {
	var deck []card.Card

	for _, cardSuit := range suitsOrder {
		for _, cardRank := range rankOrder {
			v := card.Card{
				Rank: cardRank,
				Suit: cardSuit,
			}

			deck = append(deck, v)
		}
	}

	for _, option := range options {
		option(&deck)
	}

	return deck
}

func WithSort(compare func(i card.Card, j card.Card) int) FunctionalOption {
	return func(deck *[]card.Card) {
		slices.SortFunc(*deck, compare)
	}
}

func DefaultCompare(i card.Card, j card.Card) int {
	if i.Suit != j.Suit {
		return int(i.Suit) - (int(j.Suit))
	}

	if i.Rank != j.Rank {
		return int(i.Rank) - (int(j.Rank))
	}

	return 0
}

func WithShuffle() FunctionalOption {
	return func(deck *[]card.Card) {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))

		random.Shuffle(len(*deck), swap(deck))
	}
}

func swap(deck *[]card.Card) func(i int, j int) {
	return func(i int, j int) {
		firstCard := (*deck)[i]
		secondCard := (*deck)[j]

		(*deck)[i] = secondCard
		(*deck)[j] = firstCard
	}
}

func WithJokers(jokersCount int) FunctionalOption {
	jokersSlice := make([]card.Card, jokersCount)

	for index := range jokersSlice {
		joker := card.Card{
			Suit: suits.Joker,
			Rank: ranks.Joker,
		}

		jokersSlice[index] = joker
	}

	return func(deck *[]card.Card) {
		newDeck := append(*deck, jokersSlice...)
		*deck = newDeck
	}
}

func WithFilter(filterTestFunc func(card card.Card) bool) FunctionalOption {
	return func(deck *[]card.Card) {
		filtered := make([]card.Card, 0)

		for _, v := range *deck {
			if filterTestFunc(v) {
				filtered = append(filtered, v)
			}
		}

		*deck = filtered
	}
}

func FilterOutCards(sliceOfCards []card.Card) func(card card.Card) bool {
	return func(card card.Card) bool {
		for _, v := range sliceOfCards {
			if v.Rank == card.Rank && v.Suit == card.Suit {
				return false
			}
		}

		return true
	}
}

func WithAdditionalDecks(additionalDecksCount int) FunctionalOption {
	return func(deck *[]card.Card) {
		newDeck := make([]card.Card, 0)

		for i := 0; i <= additionalDecksCount; i++ {
			newDeck = append(newDeck, *deck...)
		}

		*deck = newDeck
	}
}

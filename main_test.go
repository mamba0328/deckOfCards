package main

import (
	"github.com/mamba0328/deckOfCards/card"
	"github.com/mamba0328/deckOfCards/enums/ranks"
	"github.com/mamba0328/deckOfCards/enums/suits"
	"testing"
)

func TestNew(t *testing.T) {
	deck := New()

	if len(deck) != 52 {
		t.Error("Expected 52 cards, got ", len(deck))
	}
}

func TestDefaultSort(t *testing.T) {
	deck := New(WithShuffle(), WithSort(DefaultCompare))

	firstCard := deck[0]

	if firstCard.Suit != suits.Spade && firstCard.Rank != ranks.Ace {
		t.Error("Expected first card to be Ace of Spade, got ", firstCard)
	}

	lastCard := deck[len(deck)-1]

	if lastCard.Suit != suits.Club && lastCard.Rank != ranks.King {
		t.Error("Expected last card to be King of Club, got ", firstCard)
	}
}

func TestJokers(t *testing.T) {
	deck := New(WithJokers(3))

	count := 0

	for _, card := range deck {
		if card.Suit == suits.Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Expected 3 Jokers, got ", count)
	}
}

func NoSpades(card card.Card) bool {
	return card.Suit != suits.Spade
}

func TestFilter(t *testing.T) {
	deck := New(WithFilter(NoSpades))

	count := 0

	for _, card := range deck {
		if card.Suit == suits.Spade {
			count++
		}
	}

	if count != 0 {
		t.Error("Expected 0 Spades cards, got ", count)
	}
}

func TestWithAdditionalDecks(t *testing.T) {
	deck := New(WithFilter(NoSpades), WithJokers(2), WithAdditionalDecks(2))

	spadesCount := 0
	jokersCount := 0

	for _, card := range deck {
		if card.Suit == suits.Spade {
			spadesCount++
		}

		if card.Rank == ranks.Joker {
			jokersCount++
		}
	}

	if spadesCount != 0 {
		t.Error("Expected 0 Spades cards, got ", spadesCount)
	}

	if jokersCount != 6 {
		t.Error("Expected 6 Jokers cards, got ", jokersCount)
	}

	if len(deck) != 123 {
		t.Error("Expected length 123 got", len(deck))
	}
}

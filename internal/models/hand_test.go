package models

import (
	"testing"
)

func TestHand_Score(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected int
	}{
		{
			name: "Simple hand",
			cards: []Card{
				{Suit: Spades, Rank: Two},
				{Suit: Hearts, Rank: Three},
			},
			expected: 5,
		},
		{
			name: "Face cards",
			cards: []Card{
				{Suit: Spades, Rank: Ten},
				{Suit: Hearts, Rank: Jack},
			},
			expected: 20,
		},
		{
			name: "Ace as 11",
			cards: []Card{
				{Suit: Spades, Rank: Ace},
				{Suit: Hearts, Rank: Nine},
			},
			expected: 20,
		},
		{
			name: "Ace as 1",
			cards: []Card{
				{Suit: Spades, Rank: Ace},
				{Suit: Hearts, Rank: Nine},
				{Suit: Diamonds, Rank: Three},
			},
			expected: 13,
		},
		{
			name: "Multiple Aces",
			cards: []Card{
				{Suit: Spades, Rank: Ace},
				{Suit: Hearts, Rank: Ace},
			},
			expected: 12,
		},
		{
			name: "Multiple Aces with other cards",
			cards: []Card{
				{Suit: Spades, Rank: Ace},
				{Suit: Hearts, Rank: Ace},
				{Suit: Diamonds, Rank: Nine},
			},
			expected: 21,
		},
		{
			name: "Multiple Aces bust prevention",
			cards: []Card{
				{Suit: Spades, Rank: Ace},
				{Suit: Hearts, Rank: Ace},
				{Suit: Diamonds, Rank: Ace},
				{Suit: Clubs, Rank: Ace},
				{Suit: Spades, Rank: Ten},
			},
			expected: 14,
		},
		{
			name: "Blackjack",
			cards: []Card{
				{Suit: Spades, Rank: Ace},
				{Suit: Hearts, Rank: King},
			},
			expected: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{Cards: tt.cards}
			if got := h.Score(); got != tt.expected {
				t.Errorf("Hand.Score() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestHand_IsBust(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected bool
	}{
		{
			name: "Not bust",
			cards: []Card{
				{Suit: Spades, Rank: Ten},
				{Suit: Hearts, Rank: King},
			},
			expected: false,
		},
		{
			name: "Bust",
			cards: []Card{
				{Suit: Spades, Rank: Ten},
				{Suit: Hearts, Rank: King},
				{Suit: Diamonds, Rank: Two},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{Cards: tt.cards}
			if got := h.IsBust(); got != tt.expected {
				t.Errorf("Hand.IsBust() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestHand_IsBlackjack(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected bool
	}{
		{
			name: "Blackjack",
			cards: []Card{
				{Suit: Spades, Rank: Ace},
				{Suit: Hearts, Rank: King},
			},
			expected: true,
		},
		{
			name: "21 but not Blackjack (3 cards)",
			cards: []Card{
				{Suit: Spades, Rank: Seven},
				{Suit: Hearts, Rank: Seven},
				{Suit: Diamonds, Rank: Seven},
			},
			expected: false,
		},
		{
			name: "Not 21",
			cards: []Card{
				{Suit: Spades, Rank: Ten},
				{Suit: Hearts, Rank: King},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{Cards: tt.cards}
			if got := h.IsBlackjack(); got != tt.expected {
				t.Errorf("Hand.IsBlackjack() = %v, want %v", got, tt.expected)
			}
		})
	}
}

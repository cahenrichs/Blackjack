package models

import (
	"testing"
)

func TestSuit_String(t *testing.T) {
	tests := []struct {
		suit     Suit
		expected string
	}{
		{Spades, "Spades"},
		{Hearts, "Hearts"},
		{Diamonds, "Diamonds"},
		{Clubs, "Clubs"},
	}
	for _, tt := range tests {
		if got := tt.suit.String(); got != tt.expected {
			t.Errorf("Suit.String() = %v, want %v", got, tt.expected)
		}
	}
}

func TestRank_String(t *testing.T) {
	tests := []struct {
		rank     Rank
		expected string
	}{
		{Two, "2"},
		{Nine, "9"},
		{Ten, "10"},
		{Jack, "Jack"},
		{Queen, "Queen"},
		{King, "King"},
		{Ace, "Ace"},
	}
	for _, tt := range tests {
		if got := tt.rank.String(); got != tt.expected {
			t.Errorf("Rank.String() = %v, want %v", got, tt.expected)
		}
	}
}

func TestCard_String(t *testing.T) {
	card := Card{Suit: Spades, Rank: Ace}
	expected := "Ace of Spades"
	if got := card.String(); got != expected {
		t.Errorf("Card.String() = %v, want %v", got, expected)
	}
}

func TestCard_Value(t *testing.T) {
	tests := []struct {
		rank     Rank
		expected int
	}{
		{Two, 2},
		{Nine, 9},
		{Ten, 10},
		{Jack, 10},
		{Queen, 10},
		{King, 10},
		{Ace, 11},
	}
	for _, tt := range tests {
		card := Card{Rank: tt.rank}
		if got := card.Value(); got != tt.expected {
			t.Errorf("Card.Value() for rank %v = %v, want %v", tt.rank, got, tt.expected)
		}
	}
}

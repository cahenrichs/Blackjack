package engine

import (
	"testing"

	"github.com/cahenrichs/Blackjack/internal/models"
)

func TestGame_GetWinner(t *testing.T) {
	tests := []struct {
		name        string
		playerCards []models.Card
		dealerCards []models.Card
		expected    string
	}{
		{
			name: "Player wins (higher score)",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Nine},
			},
			dealerCards: []models.Card{
				{Suit: models.Diamonds, Rank: models.Ten},
				{Suit: models.Clubs, Rank: models.Eight},
			},
			expected: "Player wins",
		},
		{
			name: "Dealer wins (higher score)",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Eight},
			},
			dealerCards: []models.Card{
				{Suit: models.Diamonds, Rank: models.Ten},
				{Suit: models.Clubs, Rank: models.Nine},
			},
			expected: "Dealer wins",
		},
		{
			name: "Player wins (Dealer Bust)",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Eight},
			},
			dealerCards: []models.Card{
				{Suit: models.Diamonds, Rank: models.Ten},
				{Suit: models.Clubs, Rank: models.Eight},
				{Suit: models.Hearts, Rank: models.Five},
			},
			expected: "Player wins (Dealer Bust)",
		},
		{
			name: "Dealer wins (Player Bust)",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Eight},
				{Suit: models.Diamonds, Rank: models.Five},
			},
			dealerCards: []models.Card{
				{Suit: models.Clubs, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Eight},
			},
			expected: "Dealer wins (Player Bust)",
		},
		{
			name: "Push (Tie)",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Eight},
			},
			dealerCards: []models.Card{
				{Suit: models.Diamonds, Rank: models.Ten},
				{Suit: models.Clubs, Rank: models.Eight},
			},
			expected: "Push (Tie)",
		},
		{
			name: "Player Blackjack beats Dealer 21",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Ace},
				{Suit: models.Hearts, Rank: models.Ten},
			},
			dealerCards: []models.Card{
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Clubs, Rank: models.Seven},
				{Suit: models.Hearts, Rank: models.Seven},
			},
			expected: "Player wins (Blackjack)",
		},
		{
			name: "Dealer Blackjack beats Player 21",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Seven},
				{Suit: models.Hearts, Rank: models.Seven},
				{Suit: models.Diamonds, Rank: models.Seven},
			},
			dealerCards: []models.Card{
				{Suit: models.Clubs, Rank: models.Ace},
				{Suit: models.Hearts, Rank: models.Ten},
			},
			expected: "Dealer wins (Blackjack)",
		},
		{
			name: "Both Blackjack (Push)",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Ace},
				{Suit: models.Hearts, Rank: models.Ten},
			},
			dealerCards: []models.Card{
				{Suit: models.Diamonds, Rank: models.Ace},
				{Suit: models.Clubs, Rank: models.Ten},
			},
			expected: "Push (Tie)",
		},
		{
			name: "Both Bust (Dealer wins by rule)",
			playerCards: []models.Card{
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Eight},
				{Suit: models.Diamonds, Rank: models.Five},
			},
			dealerCards: []models.Card{
				{Suit: models.Clubs, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Eight},
				{Suit: models.Spades, Rank: models.Five},
			},
			expected: "Dealer wins (Player Bust)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				PlayerHand: &models.Hand{Cards: tt.playerCards},
				DealerHand: &models.Hand{Cards: tt.dealerCards},
			}
			if got := g.GetWinner(); got != tt.expected {
				t.Errorf("Game.GetWinner() = %v, want %v", got, tt.expected)
			}
		})
	}
}

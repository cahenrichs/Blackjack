package models

import "fmt"

// Suit represents one of the four suits in a deck of cards.
type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

// String returns the string representation of the suit.
func (s Suit) String() string {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	if s < 0 || int(s) >= len(suits) {
		return "Unknown"
	}
	return suits[s]
}

// Rank represents the rank of a card (Two through Ace).
type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// String returns the string representation of the rank.
func (r Rank) String() string {
	if r >= Two && r <= Ten {
		return fmt.Sprintf("%d", r)
	}
	ranks := []string{"Jack", "Queen", "King", "Ace"}
	index := int(r) - 11
	if index < 0 || index >= len(ranks) {
		return "Unknown"
	}
	return ranks[index]
}

// Card represents a single playing card with a suit and a rank.
type Card struct {
	Suit Suit
	Rank Rank
}

// String returns the string representation of the card.
func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

// Value returns the Blackjack value of the card.
// Face cards are 10, Aces are 11 (adjusted in Hand.Score).
func (c Card) Value() int {
	if c.Rank <= Ten {
		return int(c.Rank)
	}
	if c.Rank == Ace {
		return 11
	}
	return 10
}

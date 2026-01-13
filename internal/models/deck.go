package models

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// Deck represents a collection of cards.
type Deck struct {
	Cards []Card
}

// NewDeck creates a new 52-card deck.
func NewDeck() *Deck {
	deck := &Deck{
		Cards: make([]Card, 0, 52),
	}
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	ranks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	for _, suit := range suits {
		for _, rank := range ranks {
			deck.Cards = append(deck.Cards, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

// Shuffle shuffles the cards in the deck using a cryptographically secure random number generator.
func (d *Deck) Shuffle() error {
	for i := len(d.Cards) - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := int(n.Int64())
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
	return nil
}

// Draw removes and returns the top card from the deck.
// Returns an error if the deck is empty.
func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card, nil
}

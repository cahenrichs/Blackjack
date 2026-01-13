package models

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.Cards) != 52 {
		t.Errorf("NewDeck() should have 52 cards, got %d", len(deck.Cards))
	}

	// Check for duplicates
	seen := make(map[string]bool)
	for _, card := range deck.Cards {
		cardStr := card.String()
		if seen[cardStr] {
			t.Errorf("Duplicate card found: %s", cardStr)
		}
		seen[cardStr] = true
	}
}

func TestDeck_Draw(t *testing.T) {
	deck := NewDeck()
	initialCount := len(deck.Cards)

	card, err := deck.Draw()
	if err != nil {
		t.Fatalf("Draw() returned error: %v", err)
	}

	// Use the card to avoid unused variable error
	_ = card

	if len(deck.Cards) != initialCount-1 {
		t.Errorf("Deck should have %d cards after draw, got %d", initialCount-1, len(deck.Cards))
	}

	// Draw all cards
	for i := 0; i < initialCount-1; i++ {
		_, err := deck.Draw()
		if err != nil {
			t.Fatalf("Draw() returned error at index %d: %v", i, err)
		}
	}

	// Draw from empty deck
	_, err = deck.Draw()
	if err == nil {
		t.Error("Draw() from empty deck should return error")
	}
}

func TestDeck_Shuffle(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()

	// They should be identical initially
	for i := range deck1.Cards {
		if deck1.Cards[i] != deck2.Cards[i] {
			t.Fatalf("Decks should be identical before shuffle at index %d", i)
		}
	}

	err := deck1.Shuffle()
	if err != nil {
		t.Fatalf("Shuffle() returned error: %v", err)
	}

	// They should be different now (statistically very likely)
	different := false
	for i := range deck1.Cards {
		if deck1.Cards[i] != deck2.Cards[i] {
			different = true
			break
		}
	}

	if !different {
		t.Error("Deck should be different after shuffle")
	}
}

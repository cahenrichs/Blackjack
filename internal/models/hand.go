package models

// Hand represents a player's or dealer's hand of cards.
type Hand struct {
	Cards []Card
}

// AddCard adds a card to the hand.
func (h *Hand) AddCard(c Card) {
	h.Cards = append(h.Cards, c)
}

// Score calculates the best possible Blackjack score for the hand.
// It automatically adjusts Aces from 11 to 1 to avoid busting.
func (h *Hand) Score() int {
	score := 0
	aces := 0
	for _, card := range h.Cards {
		val := card.Value()
		score += val
		if card.Rank == Ace {
			aces++
		}
	}

	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}
	return score
}

// IsBust returns true if the hand's score exceeds 21.
func (h *Hand) IsBust() bool {
	return h.Score() > 21
}

// IsBlackjack returns true if the hand is a Blackjack (Ace + 10-value card).
func (h *Hand) IsBlackjack() bool {
	return len(h.Cards) == 2 && h.Score() == 21
}

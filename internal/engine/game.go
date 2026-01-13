package engine

import (
	"errors"
	"github.com/cahenrichs/Blackjack/internal/models"
)

// GameState represents the current state of the game.
type GameState int

const (
	// StateInitial is the state before cards are dealt.
	StateInitial GameState = iota
	// StateBetting is the state where the player places a bet.
	StateBetting
	// StatePlayerTurn is the state where the player can hit or stand.
	StatePlayerTurn
	// StateDealerTurn is the state where the dealer plays.
	StateDealerTurn
	// StateGameOver is the state when the game has ended.
	StateGameOver
)

// Game represents a single game of Blackjack.
type Game struct {
	Deck       *models.Deck
	PlayerHand *models.Hand
	DealerHand *models.Hand
	State      GameState
	Balance    float64
	CurrentBet float64
}

// NewGame initializes a new game with a shuffled deck and starting balance.
func NewGame() (*Game, error) {
	deck := models.NewDeck()
	err := deck.Shuffle()
	if err != nil {
		return nil, err
	}

	return &Game{
		Deck:       deck,
		PlayerHand: &models.Hand{},
		DealerHand: &models.Hand{},
		State:      StateBetting,
		Balance:    1000.0,
	}, nil
}

// PlaceBet sets the current bet and deducts it from the balance.
func (g *Game) PlaceBet(amount float64) error {
	if amount <= 0 {
		return errors.New("bet must be greater than zero")
	}
	if amount > g.Balance {
		return errors.New("insufficient balance")
	}
	g.CurrentBet = amount
	g.Balance -= amount
	return nil
}

// ResolveBet calculates the payout based on the game result and updates the balance.
func (g *Game) ResolveBet() float64 {
	pScore := g.PlayerHand.Score()
	dScore := g.DealerHand.Score()
	pBJ := g.PlayerHand.IsBlackjack()
	dBJ := g.DealerHand.IsBlackjack()

	var payout float64

	if pScore > 21 {
		// Player bust, lose bet
		payout = 0
	} else if dScore > 21 {
		// Dealer bust, player wins 1:1
		payout = g.CurrentBet * 2
	} else if pBJ && !dBJ {
		// Player Blackjack, 3:2 payout
		payout = g.CurrentBet * 2.5
	} else if dBJ && !pBJ {
		// Dealer Blackjack, lose bet
		payout = 0
	} else if pScore > dScore {
		// Player higher score, 1:1 payout
		payout = g.CurrentBet * 2
	} else if dScore > pScore {
		// Dealer higher score, lose bet
		payout = 0
	} else {
		// Push, return bet
		payout = g.CurrentBet
	}

	g.Balance += payout
	return payout
}

// ResetRound prepares the game for a new round while keeping the balance.
func (g *Game) ResetRound() error {
	if len(g.Deck.Cards) < 15 { // Reshuffle if deck is low
		g.Deck = models.NewDeck()
		err := g.Deck.Shuffle()
		if err != nil {
			return err
		}
	}
	g.PlayerHand = &models.Hand{}
	g.DealerHand = &models.Hand{}
	g.CurrentBet = 0
	g.State = StateBetting
	return nil
}

// Deal deals the initial two cards to both the player and the dealer.
func (g *Game) Deal() error {
	for i := 0; i < 2; i++ {
		card, err := g.Deck.Draw()
		if err != nil {
			return err
		}
		g.PlayerHand.AddCard(card)

		card, err = g.Deck.Draw()
		if err != nil {
			return err
		}
		g.DealerHand.AddCard(card)
	}
	g.State = StatePlayerTurn
	return nil
}

// Hit draws a card for the player. If the player busts, the game ends.
func (g *Game) Hit() error {
	card, err := g.Deck.Draw()
	if err != nil {
		return err
	}
	g.PlayerHand.AddCard(card)

	if g.PlayerHand.IsBust() {
		g.State = StateGameOver
	}
	return nil
}

// Stand ends the player's turn and starts the dealer's turn.
func (g *Game) Stand() {
	g.State = StateDealerTurn
}

// DealerPlay executes the dealer's logic: hit until score is at least 17.
func (g *Game) DealerPlay() error {
	for g.DealerHand.Score() < 17 {
		card, err := g.Deck.Draw()
		if err != nil {
			return err
		}
		g.DealerHand.AddCard(card)
	}
	g.State = StateGameOver
	return nil
}

// GetWinner determines the result of the game and returns a descriptive string.
func (g *Game) GetWinner() string {
	pScore := g.PlayerHand.Score()
	dScore := g.DealerHand.Score()
	pBJ := g.PlayerHand.IsBlackjack()
	dBJ := g.DealerHand.IsBlackjack()

	if pScore > 21 {
		return "Dealer wins (Player Bust)"
	}
	if dScore > 21 {
		return "Player wins (Dealer Bust)"
	}

	if pBJ && !dBJ {
		return "Player wins (Blackjack)"
	}
	if dBJ && !pBJ {
		return "Dealer wins (Blackjack)"
	}

	if pScore > dScore {
		return "Player wins"
	}
	if dScore > pScore {
		return "Dealer wins"
	}
	return "Push (Tie)"
}

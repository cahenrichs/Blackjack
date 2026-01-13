# Blackjack Implementation Plan

This document outlines the architecture and implementation steps for a terminal-based Blackjack game in Go.

## 1. Architecture Design

### Project Structure
```text
.
├── cmd/
│   └── blackjack/
│       └── main.go         # Entry point
├── internal/
│   ├── engine/
│   │   ├── game.go         # Core game loop and state management
│   │   └── rules.go        # Blackjack rules (scoring, bust, blackjack)
│   ├── models/
│   │   ├── card.go         # Card, Suit, Rank definitions
│   │   ├── deck.go         # Deck logic (creation, shuffling)
│   │   └── hand.go         # Hand logic (adding cards, scoring)
│   └── ui/
│       └── terminal.go     # Terminal output and input handling
├── go.mod
└── PLAN.md
```

### Data Models

#### Card
- `Suit`: (Spades, Hearts, Diamonds, Clubs)
- `Rank`: (2-10, Jack, Queen, King, Ace)
- `Value()`: Returns the numeric value (Aces handled dynamically).

#### Deck
- `Cards`: `[]Card`
- `Shuffle()`: Uses `crypto/rand` for secure shuffling.
- `Draw()`: Removes and returns the top card.

#### Hand
- `Cards`: `[]Card`
- `Score()`: Calculates the best possible score (handling Aces as 1 or 11).

#### Game
- `Deck`: The current deck.
- `PlayerHand`: The player's cards.
- `DealerHand`: The dealer's cards.
- `State`: (Initial, PlayerTurn, DealerTurn, GameOver).
- `Balance`: The player's current money.
- `CurrentBet`: The amount bet for the current round.

---

## 2. Implementation Steps

- [ ] **Phase 1: Project Setup**
    - [ ] Initialize Go module: `go mod init github.com/cahenrichs/Blackjack`.
    - [ ] Create directory structure.

- [ ] **Phase 2: Core Models (`internal/models`)**
    - [ ] Define `Suit` and `Rank` types and constants.
    - [ ] Implement `Card` struct and `String()` method.
    - [ ] Implement `Deck` struct with `NewDeck()` and `Shuffle()` using `crypto/rand`.
    - [ ] Implement `Hand` struct with `AddCard()` and `Score()` logic.

- [ ] **Phase 3: Game Engine (`internal/engine`)**
    - [ ] Define `GameState` and `Game` struct.
    - [ ] Implement `NewGame()` to initialize deck and hands.
    - [ ] Implement `Deal()` logic (2 cards each).
    - [ ] Implement `Hit()` and `Stand()` actions.
    - [ ] Implement Dealer AI (hits until 17).
    - [ ] Implement win/loss evaluation logic.

- [ ] **Phase 4: Terminal UI (`internal/ui`)**
    - [ ] Implement functions to display cards and hands.
    - [ ] Implement input handling for Hit/Stand.
    - [ ] Add color/formatting for better terminal experience (optional).

- [ ] **Phase 5: Main Loop (`cmd/blackjack`)**
    - [ ] Wire everything together in `main.go`.
    - [ ] Handle the overall game flow (Play again? Exit).

- [ ] **Phase 6: Polish & Error Handling**
    - [ ] Ensure robust error handling for deck exhaustion.
    - [ ] Validate all edge cases (e.g., Natural Blackjack on deal).
    - [ ] Add unit tests for scoring logic.

- [ ] **Phase 7: Betting System**
    - [ ] **Update `Game` struct** in `internal/engine/game.go` to include `Balance` and `CurrentBet`.
    - [ ] **Implement `PlaceBet(amount float64) error`** in `internal/engine/game.go` with validation.
    - [ ] **Implement `ResolveBet() float64`** in `internal/engine/game.go` to handle payouts (1:1, 3:2, push).
    - [ ] **Implement `ResetRound()`** in `internal/engine/game.go` to prepare for a new round.
    - [ ] **Add `GetBet(balance float64) float64`** in `internal/ui/terminal.go` to handle user bet input.
    - [ ] **Update `main.go`** to persist the `Game` state and handle the betting flow.
    - [ ] **Add game over condition** when balance reaches 0.

---

## 3. Risks & Considerations

- **Aces Logic:** Scoring Aces is the most complex part of Blackjack rules. The `Score()` function must intelligently decide whether an Ace is 1 or 11 to maximize the score without busting.
- **Randomness:** Using `crypto/rand` requires careful handling to map byte slices to integer indices for shuffling (e.g., using `binary.Read` or `rand.Int` with a crypto source).
- **Betting Logic:** Payouts for Blackjack (3:2) and regular wins (1:1) must be calculated correctly. Using `float64` for balance is acceptable for this implementation but requires care with precision if expanded.
- **Game State Management:** Transitioning between rounds while maintaining the player's balance requires careful state resets.
- **Terminal Input:** Handling user input in a loop can be messy; we should ensure clean breaks and valid input sanitization.

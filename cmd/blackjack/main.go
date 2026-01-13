package main

import (
	"fmt"
	"github.com/cahenrichs/Blackjack/internal/engine"
	"github.com/cahenrichs/Blackjack/internal/ui"
	"os"
)

func main() {
	fmt.Println("Welcome to Blackjack!")
	fmt.Println("---------------------")

	game, err := engine.NewGame()
	if err != nil {
		fmt.Printf("Error starting game: %v\n", err)
		os.Exit(1)
	}

	for {
		if game.Balance <= 0 {
			fmt.Println("You are out of money! Game Over.")
			break
		}

		if err := playRound(game); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if game.Balance <= 0 {
			fmt.Println("You are out of money! Game Over.")
			break
		}

		fmt.Printf("\nCurrent Balance: $%.2f\n", game.Balance)
		fmt.Print("Play another round? (y/n): ")
		var playAgain string
		fmt.Scanln(&playAgain)
		if playAgain != "y" {
			break
		}
		fmt.Println("\n---------------------")

		if err := game.ResetRound(); err != nil {
			fmt.Printf("Error resetting round: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("\nFinal Balance: $%.2f\n", game.Balance)
	fmt.Println("Thanks for playing!")
}

func playRound(game *engine.Game) error {
	// Betting Phase
	bet := ui.GetBet(game.Balance)
	if err := game.PlaceBet(bet); err != nil {
		return fmt.Errorf("placing bet: %w", err)
	}

	// Dealing Phase
	err := game.Deal()
	if err != nil {
		return fmt.Errorf("dealing cards: %w", err)
	}

	// Check for initial Blackjacks
	if game.PlayerHand.IsBlackjack() || game.DealerHand.IsBlackjack() {
		game.State = engine.StateGameOver
	}

	// Player Turn
	for game.State == engine.StatePlayerTurn {
		ui.DisplayHand("Dealer", game.DealerHand, true)
		ui.DisplayHand("Player", game.PlayerHand, false)

		action := ui.GetAction()
		if action == "h" {
			err := game.Hit()
			if err != nil {
				return fmt.Errorf("hitting: %w", err)
			}
			if game.PlayerHand.IsBust() {
				ui.DisplayHand("Player", game.PlayerHand, false)
				fmt.Println("Bust!")
				break
			}
		} else {
			game.Stand()
		}
	}

	// Dealer Turn
	if !game.PlayerHand.IsBust() && game.State == engine.StateDealerTurn {
		fmt.Println("\nDealer's turn...")
		err := game.DealerPlay()
		if err != nil {
			return fmt.Errorf("dealer play: %w", err)
		}
	}

	// Game Over
	fmt.Println("\nFinal Hands:")
	ui.DisplayHand("Dealer", game.DealerHand, false)
	ui.DisplayHand("Player", game.PlayerHand, false)

	fmt.Printf("Result: %s\n", game.GetWinner())

	payout := game.ResolveBet()
	if payout > 0 {
		fmt.Printf("You won $%.2f!\n", payout)
	} else {
		fmt.Println("You lost your bet.")
	}

	return nil
}

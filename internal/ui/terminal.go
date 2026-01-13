package ui

import (
	"fmt"
	"github.com/cahenrichs/Blackjack/internal/models"
)

// DisplayHand prints the cards in a hand to the terminal.
// If hideFirst is true, the first card is shown as [Hidden Card].
func DisplayHand(name string, hand *models.Hand, hideFirst bool) {
	fmt.Printf("%s's Hand:\n", name)
	for i, card := range hand.Cards {
		if i == 0 && hideFirst {
			fmt.Println("  [Hidden Card]")
		} else {
			fmt.Printf("  %s\n", card)
		}
	}
	if !hideFirst {
		fmt.Printf("Score: %d\n", hand.Score())
	}
	fmt.Println()
}

// GetAction prompts the player to either hit or stand and returns their choice.
func GetAction() string {
	var action string
	for {
		fmt.Print("Do you want to (h)it or (s)tand? ")
		fmt.Scanln(&action)
		if action == "h" || action == "s" {
			return action
		}
		fmt.Println("Invalid input. Please enter 'h' or 's'.")
	}
}

// GetBet prompts the player to enter a bet amount.
func GetBet(balance float64) float64 {
	var bet float64
	for {
		fmt.Printf("Current Balance: $%.2f\n", balance)
		fmt.Print("Enter your bet: ")
		_, err := fmt.Scanln(&bet)
		if err == nil && bet > 0 && bet <= balance {
			return bet
		}
		fmt.Printf("Invalid bet. Please enter a value between 1 and %.2f.\n", balance)
		// Clear input buffer on error
		var discard string
		fmt.Scanln(&discard)
	}
}

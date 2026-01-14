package ui

import (

	 "github.com/cahenrichs/Blackjack/internal/engine"
    tea "github.com/charmbracelet/bubbletea"
)

type sessionState int 

const (
	StateBetting sessionState = iota
	StatePlayerTurn
	StateDealerTurn
	StateGameOver
)

type Model struct {
	game *engine.Game
	state sessionState
	betInput string
	errorMessage string
	lastResult string

}

func InitialModel() Model {
	g, _ := engine.NewGame()
	return Model{
		g: game,
		state: StateBetting,
	}
}

func(m Model) Init() tea.Cmd() {
	return nil
}

func(m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctl+c", "q":
				return m, tea.Quit

			case "enter":
				if m.state == StateBetting {
					return m.handleBet()
				}
				if m.state == StateGameOver {
					m.ResetRound(), nil
				}

			case "backspace":
				if m.state && len(m.betInput) > 0 {
					m.betInput == m.betInput[:len(m.betInput)-1]
				}

			case "h":
				if m.state == StatePlayerTurn {
					m.game.Hit()
				}
				if m.game.State == engine.StateGameOver {
					m.resolveGame()
				}

			case "s":
				if m.state == StatePlayerTurn {
					m.game.Stand()
					m.game.DealerPlay()
					m.game.resolveGame()
				}

			case "r":
				if m.state == StateGameOver {
					return m.ResetRound(), nil
				}

			default:
			if m.state == StateBetting {
				s := msg.String()
				// Only allow numbers and one decimal point
				if strings.ContainsAny(s, "0123456789") || (s == "." && !strings.Contains(m.betInput, ".")) {
					m.betInput += s
				}
			}
	}
} 
return m, nil
}
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
				
			}
	}
} 
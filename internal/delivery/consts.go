package delivery

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	inputs  []textinput.Model
	focused int
	err     error
}

type (
	errMsg error
)

const (
	ticker = iota
	income
)

const (
	black     = lipgloss.Color("#000")
	blue      = lipgloss.Color("#0000FF")
	neonGreen = lipgloss.Color("#04B575")
	aqua      = lipgloss.Color("86")
	darkGray  = lipgloss.Color("#767676")
	red       = lipgloss.Color("#FF0000")
)

var (
	inputStyle = lipgloss.NewStyle().
			Foreground(aqua).
			Bold(true).
			PaddingTop(2).
			PaddingLeft(1)

	continueStyle = lipgloss.NewStyle().
			Foreground(darkGray).
			Bold(true).
			PaddingTop(2).
			PaddingLeft(1)

	errorMsg = lipgloss.NewStyle().
			Foreground(red).
			Bold(true).
			PaddingTop(2).
			PaddingLeft(1)
)

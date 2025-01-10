package delivery

import (
	"fmt"
	"log"
	"log/slog"
	"strconv"

	"github.com/andreposman/magic-number/internal/controller"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func StartTerminalLayer() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func initialModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 2)
	inputs[ticker] = textinput.New()
	inputs[ticker].Placeholder = "Example: HGLG11"
	inputs[ticker].Focus()
	inputs[ticker].CharLimit = 6
	inputs[ticker].Width = 60
	inputs[ticker].Prompt = ""
	inputs[ticker].Validate = tickerValidator

	inputs[income] = textinput.New()
	inputs[income].Placeholder = "Example: 1500"
	inputs[income].CharLimit = 10
	inputs[income].Width = 100
	inputs[income].Prompt = ""
	inputs[income].Validate = incomeValidator

	return model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Validate the current input
			if err := m.inputs[m.focused].Validate(m.inputs[m.focused].Value()); err != nil {
				m.err = err
				return m, nil // Display error and do not proceed
			}

			// If on the last input, process the inputs
			if m.focused == len(m.inputs)-1 {
				tickerInput := m.inputs[ticker].Value()
				incomeInput := m.inputs[income].Value()

				// Validate all inputs before proceeding
				if err := m.inputs[ticker].Validate(tickerInput); err != nil {
					m.err = fmt.Errorf("ticker validation failed: %v", err)
					return m, nil
				}
				if err := m.inputs[income].Validate(incomeInput); err != nil {
					m.err = fmt.Errorf("income validation failed: %v", err)
					return m, nil
				}
				slog.Info("Terminal sucessfully called, calling FII Handler now")

				fii, err := controller.FetchFIIData(tickerInput)
				if err != nil {
					slog.Error("error fetching fii data: ", slog.Any("", err))
				}

				slog.Info("", slog.Any("fii data", *fii))

				return m, tea.Quit
			}

			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	errorMessage := ""
	if m.err != nil {
		errorMessage = errorMsg.Render(fmt.Sprintf("\nError: %v", m.err))
	}

	return fmt.Sprintf(
		`
 %s
 %s

 %s 
 %s 

 %s
 %s
 
`,
		inputStyle.Width(60).Render("Type in the ticker of the FII: "),
		m.inputs[ticker].View(),
		inputStyle.Width(100).Render("Type in how much R$ do you want to receive per month (just the value): "),
		m.inputs[income].View(),
		continueStyle.Render("Continue ->"),
		errorMessage,
	) + "\n"
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func tickerValidator(s string) error {
	if len(s) != 6 {
		return fmt.Errorf("ticker must be exactly 6 characters long")
	}

	return nil
}

func incomeValidator(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return fmt.Errorf("income must be a valid number")
	}

	income, _ := strconv.ParseFloat(s, 64)
	if income <= 0 {
		return fmt.Errorf("income must be greater than zero")
	}

	return nil
}

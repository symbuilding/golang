package main

import (
	"calc/eval"
	"calc/lexer"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	expr      string
	textInput textinput.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter an expression"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		expr:      "0",
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
            val := eval.Evaluate(lexer.NewLexer(m.textInput.Value()))
            m.expr = fmt.Sprintf("%f", val)
			m.textInput.SetValue("")
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Your message: %s\n\n%s\n\n%s",
		m.expr,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}


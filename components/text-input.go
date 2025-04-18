package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInput struct {
	Placeholder string
	Label       string
	textInput   textinput.Model
	Value       *string
	Type        string
}

func NewTextInput(props TextInput) TextInput {
	ti := textinput.New()
	ti.Placeholder = props.Placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	if props.Type == "password" {
		ti.EchoMode = textinput.EchoPassword
		ti.EchoCharacter = '*'
	}
	props.textInput = ti

	return props
}

func (m TextInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.textInput.Value() != "" {
				*m.Value = m.textInput.Value()
			} else {
				*m.Value = m.Placeholder
			}
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m TextInput) View() string {
	return fmt.Sprintf("%s %s\n", m.Label, m.textInput.View())

}

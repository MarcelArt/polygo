package components

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type SingleSelect struct {
	Choices []string
	cursor  int
	Value   *string
	Label   string
}

func (m SingleSelect) Init() tea.Cmd {
	return nil
}

func (m SingleSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown:
			if m.cursor < len(m.Choices)-1 {
				m.cursor++
			}
		case tea.KeyEnter:
			*m.Value = m.Choices[m.cursor]
			log.Println(*m.Value)
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m SingleSelect) View() string {
	s := fmt.Sprintf("%s\n", m.Label)
	for i, choice := range m.Choices {
		if i == m.cursor {
			s += "> " + choice + "\n"
		} else {
			s += "  " + choice + "\n"
		}
	}
	s += "\nPress ↑/↓ to navigate, Enter to select.\n"
	return s
}

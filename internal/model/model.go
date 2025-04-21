package model

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func InitialModel(txtStyle lipgloss.Style, quitStyle lipgloss.Style) model {
	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = "John Appleseed"
	ti.Width = 30

	return model{
		ti:        ti,
		txtStyle:  txtStyle,
		quitStyle: quitStyle,
	}
}

// Just a generic tea.Model to demo terminal information of ssh.
type model struct {
	ti        textinput.Model
	txtStyle  lipgloss.Style
	quitStyle lipgloss.Style
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			os.WriteFile("output.txt", []byte(m.ti.Value()), 0644)
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.ti, cmd = m.ti.Update(msg)

	return m, cmd
}

func (m model) View() string {
	s := fmt.Sprintf("%v", m.ti.View())
	return m.txtStyle.Render(s) + "\n\n" + m.quitStyle.Render("Press 'q' to quit\n")
}

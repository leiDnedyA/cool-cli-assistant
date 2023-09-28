package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"Bookmarks", "Projects", "Misc"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string = "What would you like to do?\n\n"
	
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		} 
		
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s, err := cowsay.Say(s, cowsay.Type("default"), cowsay.BallonWidth(40));
	if err != nil{
		panic(err);
	}
	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

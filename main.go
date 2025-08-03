package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Mail struct {
	name string
	id string
}

type model struct {
	mails []Mail
	cursor int
	selected bool
	help help.Model
	table table.Model
	api_key string
	error *string

	height int
	width int

	loaded bool
}

func initialModel() model {
	columns := []table.Column{
		{ Title: "Time", Width: 20 },
		{ Title: "Description", Width: 50 },
		{ Title: "Location", Width: 25 },
		{ Title: "Source", Width: 20 },
	}

	rows := []table.Row{}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#0E74C6")).
		Bold(false)
	t.SetStyles(s)

	var model model = model{
		mails: []Mail{},
		selected: false,
		help: help.New(),
		table: t,
		loaded: false,
	}

	config, err := ReadConfig()
	if err == nil {
		model.api_key = config.api_key
	} else {
		errorMsg := err.Error()
		model.error = &errorMsg
	}

	model.help.Styles = help.Styles{
		ShortKey:       lipgloss.NewStyle(),
		ShortDesc:      lipgloss.NewStyle(),
		ShortSeparator: lipgloss.NewStyle(),
		Ellipsis:       lipgloss.NewStyle(),
		FullKey:        lipgloss.NewStyle(),
		FullDesc:       lipgloss.NewStyle(),
		FullSeparator:  lipgloss.NewStyle(),
	}

	return model
}

func (m model) Init() tea.Cmd {
	if m.api_key != "" {
		return fetchMails(m.api_key)
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Homepage keybinds handling
	if m.selected == false {
		switch msg := msg.(type) {
		    case tea.KeyMsg:
		        switch msg.String() {
			        case "up", "k":
			            if m.cursor > 0 {
			                m.cursor--
			            }
			        case "down", "j":
			            if m.cursor < len(m.mails)-1 {
			                m.cursor++
			            }
			        case "enter":
			            m.selected = true
			    }
		    }
	} else {
		// Selected keybinds handling
		switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
					case "up", "k":
						m.table.MoveUp(1)
					case "down", "j":
						m.table.MoveDown(1)
					case "backspace":
						m.selected = false
				}
		}
	}

	// Global keybind handling
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.height, m.width, m.help.Width = msg.Height, msg.Width, msg.Width
	    case tea.KeyMsg:
	        switch msg.String() {
			    case "ctrl+c", "q":
			    	return m, tea.Quit
				case "?", "h":
					m.help.ShowAll = !m.help.ShowAll
			}
		case fetchMsg:
			m.loaded = true

			if msg.error != nil {
				errorMsg := msg.error.Error()
				m.error = &errorMsg
				return m, nil
			}

			for _, mail := range msg.mails.Mails {
				Type := strings.ToUpper(mail.Type[:1]) + strings.ToLower(mail.Type[1:])
				m.mails = append(m.mails, Mail{ fmt.Sprintf("%s: %s", Type, mail.Title), mail.ID })
			}
	}

    return m, nil
}

func (m model) View() string {
	s := "\n"
	tab := "  "

	style_error := lipgloss.NewStyle().
		PaddingBottom(1).
		PaddingTop(1).
		PaddingLeft(2).
		PaddingRight(2).
		Border(lipgloss.RoundedBorder())

	if m.error != nil {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, style_error.Render(*m.error))
	}

	if !m.loaded {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, style_error.Render("Loading..."))
	}

	if m.selected {
		s += tab
		s += m.mails[m.cursor].name
		s += "\n"

		s += strings.Repeat(tab, 2)
		s += "Created at "
		s += "2025-07-28"
		s += "\n\n"

		s += strings.Repeat(tab, 2)
		s += "Status   : "
		s += "Mailed"
		s += "\n"

		s += strings.Repeat(tab, 2)
		s += "Tag      : "
		s += "summer-of-making"
		s += "\n"

		s += strings.Repeat(tab, 2)
		s += "Letter ID: "
		s += m.mails[m.cursor].id
		s += "\n\n\n"

		tableView := tab + strings.ReplaceAll(m.table.View(), "\n", "\n"+tab)
		s += tableView
	} else {
		s += tab
		s += "Your mails\n"
	    for i, mail := range m.mails {
	        cursor := " "
	        if m.cursor == i {
	            cursor = "›"
	        }

	        s += fmt.Sprintf("%s%s %s\n", tab, cursor, mail.name)
	    }
	}

    var helpView string
    if m.selected {
   		helpView = m.help.View(selectedKeys)
    } else {
   		helpView = m.help.View(homeKeys)
    }
    height := m.height - strings.Count(s, "\n") - strings.Count(helpView, "\n")

    if height < 0 {
    	s += "\n"
    } else {
    	s += strings.Repeat("\n", height - 2)
    }

    helpViewSplitted := strings.Split(helpView, "\n")
    helpView = ""
    for _, line := range helpViewSplitted {
   		helpView += tab
    	helpView += line
     	helpView += "\n"

    }
    s += helpView

    return s
}

func main() {
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

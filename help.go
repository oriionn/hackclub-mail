package main

import "github.com/charmbracelet/bubbles/key"

// Home help
type homeKeymap struct {
	Up key.Binding
	Down key.Binding
	Choose key.Binding
	Quit key.Binding
	Help key.Binding
}

var homeKeys = homeKeymap{
	Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
	Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
	Choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
	Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	Help: key.NewBinding(
			key.WithKeys("?", "h"),
			key.WithHelp("?", "toggle help"),
		),
}

func (k homeKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k homeKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Choose},
		{k.Help, k.Quit},
	}
}

// Selected help
type selectedKeymap struct {
	Up key.Binding
	Down key.Binding
	Back key.Binding
	Quit key.Binding
	Help key.Binding
}


var selectedKeys = selectedKeymap{
	Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
	Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
	Back: key.NewBinding(
			key.WithKeys("return"),
			key.WithHelp("return", "back to home"),
		),
	Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	Help: key.NewBinding(
			key.WithKeys("?", "h"),
			key.WithHelp("?", "toggle help"),
		),
}

func (k selectedKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k selectedKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Back},
		{k.Help, k.Quit},
	}
}

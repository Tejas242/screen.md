package ui

import "github.com/charmbracelet/lipgloss"

type HelpModel struct {
	keys    map[string]string
	visible bool
}

func NewHelp() HelpModel {
	return HelpModel{
		keys: map[string]string{
			"tab":    "Switch Focus",
			"ctrl+s": "Save",
			"ctrl+f": "Find",
			"ctrl+q": "Quit",
			"↑/↓":    "Scroll",
			"?":      "Toggle Help",
			"ctrl+c": "Copy",
			"ctrl+v": "Paste",
			"ctrl+x": "Cut",
		},
		visible: true,
	}
}

func (h HelpModel) View() string {
	if !h.visible {
		return ""
	}

	var helps []string
	for key, desc := range h.keys {
		helps = append(helps,
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				KeyStyle.Render(key),
				" ",
				DescStyle.Render(desc),
			),
		)
	}

	return HelpStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			helps...,
		),
	)
}

package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Color palette
	BaseColors = struct {
		Primary   lipgloss.Color
		Secondary lipgloss.Color
		Accent    lipgloss.Color
		Text      lipgloss.Color
		Subtle    lipgloss.Color
		Error     lipgloss.Color
	}{
		Primary:   lipgloss.Color("#7D56F4"),
		Secondary: lipgloss.Color("#43BF6D"),
		Accent:    lipgloss.Color("#F25D94"),
		Text:      lipgloss.Color("#FFFFFF"),
		Subtle:    lipgloss.Color("#383838"),
		Error:     lipgloss.Color("#FF5F87"),
	}

	// Base styles
	BaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(BaseColors.Subtle)

	// Editor styles
	EditorStyle = func(active bool) lipgloss.Style {
		s := BaseStyle.
			Padding(1).
			BorderStyle(lipgloss.RoundedBorder())

		if active {
			return s.BorderForeground(BaseColors.Primary)
		}
		return s.BorderForeground(BaseColors.Subtle)
	}

	// Preview styles
	PreviewStyle = func(active bool) lipgloss.Style {
		s := BaseStyle.
			Padding(1).
			BorderStyle(lipgloss.RoundedBorder())

		if active {
			return s.BorderForeground(BaseColors.Secondary)
		}
		return s.BorderForeground(BaseColors.Subtle)
	}

	// Help style
	HelpStyle = lipgloss.NewStyle().
			Foreground(BaseColors.Subtle).
			Background(lipgloss.Color("#2A2A2A")).
			Padding(0, 1).
			Width(100).
			Align(lipgloss.Center)

	// Title style
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(BaseColors.Primary).
			MarginLeft(2)

	// Status style
	StatusStyle = lipgloss.NewStyle().
			Foreground(BaseColors.Text).
			Background(BaseColors.Subtle).
			Padding(0, 1)

	KeyStyle = lipgloss.NewStyle().
			Foreground(BaseColors.Primary).
			Bold(true)

	DescStyle = lipgloss.NewStyle().
			Foreground(BaseColors.Text)

	LoadingStyle = lipgloss.NewStyle().
			Foreground(BaseColors.Primary).
			Bold(true).
			Align(lipgloss.Center)
)

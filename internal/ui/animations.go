package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SpinnerMsg struct{ Done bool }
type FadeMsg struct{ Frame int }
type TransitionMsg struct{ Progress float64 }

func NewSpinner() spinner.Model {
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(BaseColors.Primary)
    return s
}

func FadeIn() tea.Cmd {
    return tea.Tick(time.Millisecond*16, func(t time.Time) tea.Msg {
        return FadeMsg{Frame: 1}
    })
}

func Transition() tea.Cmd {
    return tea.Tick(time.Millisecond*16, func(t time.Time) tea.Msg {
        return TransitionMsg{Progress: 0.1}
    })
}

package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const initialText = `# Welcome to Markdown Editor!

## Features
- Real-time preview
- Syntax highlighting
- Beautiful UI
- Side by side view

### Try Some Code:
` + "```go" + `
func main() {
    fmt.Println("Hello, Markdown!")
}
` + "```" + `

Start typing here...`

type focusArea int

const (
	editorFocus focusArea = iota
	previewFocus
)

type model struct {
	editor    textarea.Model
	preview   viewport.Model
	renderer  *glamour.TermRenderer
	ready     bool
	width     int
	height    int
	focus     focusArea
}

var (
	// Theme colors
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	// Styles
	titleStyle = lipgloss.NewStyle().
		Foreground(highlight).
		Bold(true).
		Padding(0, 1)

	activeEditorStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(highlight).
		Padding(1).
		Background(lipgloss.Color("#1F1F1F"))

	inactiveEditorStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(subtle).
		Padding(1).
		Background(lipgloss.Color("#1F1F1F"))

	helpStyle = lipgloss.NewStyle().
		Foreground(subtle).
		Background(lipgloss.Color("#2A2A2A")).
		Padding(1).
		Width(100)
)

func initialModel() model {
	ta := textarea.New()
	ta.SetValue(initialText)
	ta.Focus()
	ta.ShowLineNumbers = true
	ta.Placeholder = "Enter markdown here..."
	ta.CharLimit = 5000
	ta.SetWidth(50)
	ta.SetHeight(20)

	// Style the textarea
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle().
		Background(lipgloss.Color("#303030"))
	ta.FocusedStyle.Base = lipgloss.NewStyle().
		Background(lipgloss.Color("#1F1F1F")).
		Foreground(lipgloss.Color("#FFFFFF"))
	ta.BlurredStyle.Base = lipgloss.NewStyle().
		Background(lipgloss.Color("#1F1F1F")).
		Foreground(lipgloss.Color("#888888"))

	vp := viewport.New(50, 20)
	vp.Style = lipgloss.NewStyle().
		Background(lipgloss.Color("#1F1F1F")).
		Foreground(lipgloss.Color("#FFFFFF"))

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
	)

	return model{
		editor:   ta,
		preview:  vp,
		renderer: renderer,
		focus:    editorFocus,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab":
			// Switch focus between editor and preview
			if m.focus == editorFocus {
				m.focus = previewFocus
				m.editor.Blur()
			} else {
				m.focus = editorFocus
				m.editor.Focus()
			}
		}

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		headerHeight := 0
		footerHeight := 3 // Space for help text
		verticalMarginHeight := 2

		bodyHeight := msg.Height - headerHeight - footerHeight - verticalMarginHeight

		horizontalMargins := 4
		halfWidth := (msg.Width - horizontalMargins) / 2

		m.editor.SetWidth(halfWidth)
		m.editor.SetHeight(bodyHeight)

		m.preview.Width = halfWidth
		m.preview.Height = bodyHeight

		if !m.ready {
			m.ready = true
		}
	}

	if m.focus == editorFocus {
		var cmd tea.Cmd
		m.editor, cmd = m.editor.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		var cmd tea.Cmd
		m.preview, cmd = m.preview.Update(msg)
		cmds = append(cmds, cmd)
	}

	rendered, _ := m.renderer.Render(m.editor.Value())
	m.preview.SetContent(rendered)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	// Apply active/inactive styles based on focus
	editorBox := inactiveEditorStyle
	previewBox := inactiveEditorStyle
	if m.focus == editorFocus {
		editorBox = activeEditorStyle
	} else {
		previewBox = activeEditorStyle
	}

	// Help text
	help := helpStyle.Render("TAB: Switch Focus • ESC/Ctrl+C: Quit • ↑/↓: Scroll")

	// Main content
	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Center,
		editorBox.Render(m.editor.View()),
		previewBox.Render(m.preview.View()),
	)

	// Combine everything
	return lipgloss.JoinVertical(
		lipgloss.Left,
		mainContent,
		help,
	)
}

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}

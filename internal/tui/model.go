package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"screen.md/internal/ui"
)

const initialText = `# Welcome to Screen.md!

This is a paragraph.

This is another paragraph.

## Features
- Real-time preview
- Split view
- Syntax highlighting

### Code Example:
` + "```go" + `
func main() {
    fmt.Println("Hello!")
}
` + "```" + `
`

type Model struct {
	editor      textarea.Model
	preview     viewport.Model
	renderer    *glamour.TermRenderer
	help        ui.HelpModel
	spinner     spinner.Model
	ready       bool
	width       int
	height      int
	focus       ui.Focus
	loading     bool
	transition  float64
	lastContent string
	clipboard   *ClipboardManager
	selection   string // Track selected/copied text
	clipboardOp string // Track last clipboard operation
}

func NewModel() Model {
	ta := textarea.New()
	ta.SetWidth(50)
	ta.SetHeight(20)
	ta.ShowLineNumbers = true
	ta.Focus()
	ta.SetValue(initialText)

	vp := viewport.New(50, 20)

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
		glamour.WithEmoji(),
		glamour.WithStylesFromJSONBytes([]byte(`{
            "document": { "margin": 1 },
            "block_quote": { "margin": 1, "indent": 2 },
            "paragraph": { "margin": 1 },
            "list": { "margin": 1, "level_indent": 2 },
            "heading": { "margin": 1 },
            "code_block": { "margin": 1 }
        }`)),
	)

	return Model{
		editor:      ta,
		preview:     vp,
		renderer:    renderer,
		help:        ui.NewHelp(),
		spinner:     ui.NewSpinner(),
		focus:       ui.EditorFocus,
		loading:     true,
		clipboard:   NewClipboardManager(),
		clipboardOp: "",
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		m.spinner.Tick,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			if m.focus == ui.EditorFocus {
				textToCopy := m.getCurrentLine()
				if err := m.clipboard.Copy(textToCopy); err == nil {
					m.selection = textToCopy
					m.clipboardOp = "Copied"
				} else {
					m.clipboardOp = "Copy failed"
				}
				return m, nil
			}
			return m, tea.Quit

		case "ctrl+v":
			if m.focus == ui.EditorFocus {
				if text, err := m.clipboard.Paste(); err == nil {
					m.editor.InsertString(text)
					m.clipboardOp = "Pasted"
				} else {
					m.clipboardOp = "Paste failed"
				}
			}

		case "ctrl+x":
			if m.focus == ui.EditorFocus {
				textToCut := m.getCurrentLine()
				if err := m.clipboard.Cut(textToCut); err == nil {
					// Delete the current line by replacing it with empty string
					currentValue := m.editor.Value()
					lines := strings.Split(currentValue, "\n")
					for i := range lines {
						if lines[i] == textToCut {
							if i == len(lines)-1 {
								// Last line
								if i > 0 {
									// Remove last line and trailing newline
									newValue := strings.Join(lines[:i], "\n")
									m.editor.SetValue(newValue)
								} else {
									// Only line
									m.editor.SetValue("")
								}
							} else {
								// Not last line
								newLines := append(lines[:i], lines[i+1:]...)
								m.editor.SetValue(strings.Join(newLines, "\n"))
							}
							break
						}
					}
					m.selection = textToCut
					m.clipboardOp = "Cut"
				} else {
					m.clipboardOp = "Cut failed"
				}
			}

		case "tab":
			if m.focus == ui.EditorFocus {
				m.focus = ui.PreviewFocus
				m.editor.Blur()
			} else {
				m.focus = ui.EditorFocus
				m.editor.Focus()
			}
		}

	case tea.WindowSizeMsg:
		m.resizeComponents(msg)
	}

	if m.focus == ui.EditorFocus {
		var editorCmd tea.Cmd
		m.editor, editorCmd = m.editor.Update(msg)
		cmds = append(cmds, editorCmd)

		currentContent := m.editor.Value()
		if currentContent != m.lastContent {
			m.lastContent = currentContent
			if rendered, err := m.renderer.Render(currentContent); err == nil {
				rendered = strings.ReplaceAll(rendered, "\n\n", "\n")
				m.preview.SetContent(rendered)
			}
		}
	} else {
		var previewCmd tea.Cmd
		m.preview, previewCmd = m.preview.Update(msg)
		cmds = append(cmds, previewCmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.ready {
		return m.loadingView()
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.getTitleBar(),
		m.getMainContent(),
		m.getStatusBar(),
		m.getHelpBar(),
	)
}

func (m Model) GetCursor() int {
	value := m.editor.Value()
	lines := strings.Split(value, "\n")
	var pos int
	for _, line := range lines {

		if pos+len(line) >= m.editor.LineInfo().ColumnOffset {
			break
		}
		pos += len(line) + 1
	}
	return pos
}

func (m Model) getCurrentLine() string {
	value := m.editor.Value()
	if len(value) == 0 {
		return ""
	}

	// Get current cursor position
	lines := strings.Split(value, "\n")
	var currentPos int
	currentLine := 0

	// Find the current line
	for i, line := range lines {
		if currentPos+len(line) >= m.editor.LineInfo().ColumnOffset {
			currentLine = i
			break
		}
		currentPos += len(line) + 1
	}

	if currentLine < len(lines) {
		return lines[currentLine]
	}

	return ""
}

func (m Model) getHelpBar() string {
	helpText := []string{
		ui.KeyStyle.Render("tab") + " Switch Focus",
		ui.KeyStyle.Render("ctrl+c") + " Copy/Quit",
		ui.KeyStyle.Render("ctrl+v") + " Paste",
		ui.KeyStyle.Render("ctrl+x") + " Cut",
		ui.KeyStyle.Render("↑/↓") + " Scroll",
	}

	return ui.HelpStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			strings.Join(helpText, " • "),
		),
	)
}

func (m *Model) resizeComponents(msg tea.WindowSizeMsg) {
	m.width = msg.Width
	m.height = msg.Height

	headerHeight := 1
	footerHeight := 3
	verticalMarginHeight := 2

	bodyHeight := msg.Height - headerHeight - footerHeight - verticalMarginHeight
	halfWidth := (msg.Width - 4) / 2

	m.editor.SetWidth(halfWidth)
	m.editor.SetHeight(bodyHeight)

	m.preview.Width = halfWidth
	m.preview.Height = bodyHeight

	if !m.ready {
		m.ready = true
	}
}

func (m Model) loadingView() string {
	return ui.LoadingStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			"Loading Screen.md...",
			m.spinner.View(),
		),
	)
}

func (m Model) getTitleBar() string {
	return ui.TitleStyle.Render("Screen.md")
}

func (m Model) getStatusBar() string {
	mode := "Editor"
	if m.focus == ui.PreviewFocus {
		mode = "Preview"
	}

	if m.clipboardOp != "" {
		mode += " | " + m.clipboardOp
	}

	if m.selection != "" {
		mode += " | Text in clipboard"
	}

	return ui.StatusStyle.Render(mode)
}

func (m Model) getMainContent() string {
	editorStyle := ui.EditorStyle(m.focus == ui.EditorFocus)
	previewStyle := ui.PreviewStyle(m.focus == ui.PreviewFocus)

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		editorStyle.Render(m.editor.View()),
		previewStyle.Render(m.preview.View()),
	)
}

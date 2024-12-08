# Screen.md 📝

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D%201.16-00ADD8.svg)
![Status](https://img.shields.io/badge/status-alpha-orange.svg)

Screen.md is a beautiful, terminal-based Markdown editor and previewer designed for distraction-free writing and blogging. Built with Go and the Charm libraries, it offers a modern, efficient interface for crafting Markdown content right in your terminal.

---

> 🚧 **Note**: This project is under active development. Features and UI may change as it evolves to better suit writing needs.

---

## 🛣️ Roadmap

- [x] Basic editor and preview layout
- [x] Real-time markdown rendering
- [x] Syntax highlighting
- [ ] Vim mode support
- [ ] File saving/loading
- [ ] Custom theme support
- [ ] Image preview support (terminal-compatible)
- [ ] Configuration file support

- [ ] LLM Integration for writing assistance
- [ ] Git integration for version control
- [ ] Custom markdown extensions
- [ ] Spell checking
- [ ] Table of contents generation

- [ ] Multiple file support with tabs
- [ ] Search and replace
- [ ] Custom snippets
- [ ] Export to various formats (PDF, HTML)

## 🎯 Quick Start

### Prerequisites
- Go 1.16 or higher
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/tejas242/screen.md.git

# Navigate to the project directory
cd screen.md

# Install dependencies
go mod tidy

# Build and run
go run main.go
```

## 🎮 Usage

### Keyboard Shortcuts
- `Tab` - Switch between editor and preview
- `Ctrl+C` / `Esc` - Exit the application
- `↑/↓` - Scroll in preview mode
- Regular text editing keys work as expected

### Writing
1. Start the application
2. Begin typing in the editor (left pane)
3. Watch your markdown render in real-time (right pane)
4. Use Tab to switch focus when you want to scroll through the preview


## 🔧 Development

MarkdownFlow is built with:
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Glamour](https://github.com/charmbracelet/glamour) - Markdown rendering
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [Chroma](https://github.com/alecthomas/chroma) - Syntax highlighting

### Project Structure (May change in future)
```
markdownflow/
├── main.go        # Main application code
├── go.mod         # Go module definition
├── go.sum         # Dependencies checksum
└── README.md      # Documentation
```

## 🤝 Contributing

While this is primarily a personal utility tool, suggestions and feedback are welcome! Feel free to:
1. Open issues for bugs or feature suggestions
2. Submit pull requests for improvements
3. Share ideas for new features

## 📝 Personal Note

I created MarkdownFlow as a personal tool for writing blog posts and documentation. It's designed to be minimal yet powerful, focusing on the writing experience while providing modern features that make markdown editing more enjoyable.

## 📸 Screenshots

[screenshots here showing:
1. Main interface with editor and preview
2. Different focus modes
3. Sample markdown rendering
4. Any special features]

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"linux-guide/data"
)

var (
	// =========================
	// CYBER MODERN PALETTE
	// =========================

	styleTitle  = "\033[1;38;5;45m"  // neon cyan
	styleHeader = "\033[1;38;5;141m" // purple soft

	styleItem   = "\033[0;38;5;250m" // light gray
	styleCursor = "\033[1;38;5;51m"  // bright cyan

	styleCmdKey = "\033[1;38;5;87m"  // electric blue
	styleDesc   = "\033[0;38;5;245m" // muted gray

	styleKeyFlag = "\033[1;38;5;220m" // amber

	styleReset = "\033[0m"
)

type model struct {
	cursor         int
	state          int
	currentSection *data.Section
	menuItems      []string

	scroll int
	lines  []string

	height int
	width  int
}

func newModel() model {
	m := model{}

	sections := data.GetAllSections()
	for i := 0; i < len(sections); i++ {
		if sec, ok := sections[i+1]; ok {
			m.menuItems = append(m.menuItems, sec.Title)
		}
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := (msg).(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			if m.state == 0 {
				if sec, ok := data.GetAllSections()[m.cursor+1]; ok {
					m.currentSection = &sec
					m.state = 1
					m.scroll = 0
					m.lines = buildLines(m)
				}
			} else {
				m.state = 0
			}

		case "esc":
			if m.state == 1 {
				m.state = 0
			} else {
				return m, tea.Quit
			}

		case "up", "k":
			if m.state == 0 {
				if m.cursor > 0 {
					m.cursor--
				}
			} else {
				if m.scroll > 0 {
					m.scroll--
				}
			}

		case "down", "j":
			if m.state == 0 {
				if m.cursor < len(m.menuItems)-1 {
					m.cursor++
				}
			} else {
				m.scroll++
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var out strings.Builder

	// =========================
	// MENU
	// =========================
	if m.state == 0 {
		out.WriteString(fmt.Sprintf("%s\n  🐧 LINUX GUIDE (CYBER)\n%s\n\n", styleTitle, styleReset))

		for i, item := range m.menuItems {
			cursor := " "
			color := styleItem

			if i == m.cursor {
				cursor = "❯"
				color = styleCursor
			}

			out.WriteString(fmt.Sprintf("  %s %s%s%s\n",
				cursor,
				color,
				item,
				styleReset,
			))
		}

		out.WriteString("\n  ↑↓ / j k • Enter • q")
		return out.String()
	}

	// =========================
	// INIT CONTENT
	// =========================
	if len(m.lines) == 0 && m.currentSection != nil {
		m.lines = buildLines(m)
	}

	// =========================
	// AUTO HEIGHT
	// =========================
	height := m.height
	if height <= 0 {
		height = 25
	}

	// =========================
	// CLAMP SCROLL
	// =========================
	maxScroll := len(m.lines) - height
	if maxScroll < 0 {
		maxScroll = 0
	}

	if m.scroll < 0 {
		m.scroll = 0
	}

	if m.scroll > maxScroll {
		m.scroll = maxScroll
	}

	end := m.scroll + height
	if end > len(m.lines) {
		end = len(m.lines)
	}

	return strings.Join(m.lines[m.scroll:end], "\n")
}

func buildLines(m model) []string {
	var out []string

	line := strings.Repeat("─", 60)

	out = append(out,
		styleTitle+line+styleReset,
		"  "+styleHeader+m.currentSection.Title+styleReset,
		styleTitle+line+styleReset,
	)

	for _, item := range m.currentSection.Items {

		switch item.Type {

		case data.TypeHeader:
			out = append(out, "", styleHeader+item.Value+styleReset)

		case data.TypeCmd:
			out = append(out,
				"  "+styleCmdKey+"$ "+item.Value+styleReset,
			)

			if item.Desc != "" {
				out = append(out,
					"     "+styleDesc+"→ "+item.Desc+styleReset,
				)
			}

		case data.TypeKey:
			out = append(out,
				"     "+styleKeyFlag+"["+item.Key+"] "+styleReset+item.Desc,
			)
		}
	}

	return out
}

func main() {
	p := tea.NewProgram(newModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"linux-guide/data"
	"linux-guide/ui"
)

const (
	stateMainMenu = iota
	stateSubMenu
	stateContent
)

type menuFrame struct {
	title  string
	items  []data.MenuItem
	cursor int
}

type model struct {
	cursor int
	state  int

	rootMenu   []data.MenuItem
	rootCursor int
	menuStack  []menuFrame

	currentSection *data.Section

	scroll int
	lines  []string

	height int
	width  int
}

func newModel() model {
	return model{
		rootMenu: data.GetMenu(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) activeItems() []data.MenuItem {
	if len(m.menuStack) == 0 {
		return m.rootMenu
	}
	return m.menuStack[len(m.menuStack)-1].items
}

func (m model) menuTitle() string {
	if len(m.menuStack) == 0 {
		return "Инструменты для DevOps"
	}
	return m.menuStack[len(m.menuStack)-1].title
}

func (m model) menuSubtitle() string {
	if len(m.menuStack) <= 1 {
		return "выберите раздел"
	}
	parts := make([]string, len(m.menuStack)-1)
	for i := 0; i < len(m.menuStack)-1; i++ {
		parts[i] = m.menuStack[i].title
	}
	return strings.Join(parts, "  ›  ")
}

func (m *model) saveMenuCursor() {
	if len(m.menuStack) == 0 {
		m.rootCursor = m.cursor
		return
	}
	m.menuStack[len(m.menuStack)-1].cursor = m.cursor
}

func (m *model) enterGroup(item data.MenuItem) {
	m.saveMenuCursor()
	m.menuStack = append(m.menuStack, menuFrame{
		title:  item.Title,
		items:  item.Children,
		cursor: 0,
	})
	m.cursor = 0
	m.state = stateSubMenu
}

func (m *model) openSection(section *data.Section) {
	m.saveMenuCursor()
	m.currentSection = section
	m.state = stateContent
	m.scroll = 0
	m.lines = ui.BuildSectionLines(section)
}

func (m *model) goBack() {
	switch m.state {
	case stateContent:
		m.currentSection = nil
		m.lines = nil
		m.scroll = 0
		if len(m.menuStack) > 0 {
			m.state = stateSubMenu
			m.cursor = m.menuStack[len(m.menuStack)-1].cursor
		} else {
			m.state = stateMainMenu
			m.cursor = m.rootCursor
		}

	case stateSubMenu:
		if len(m.menuStack) == 0 {
			m.state = stateMainMenu
			return
		}
		m.cursor = m.menuStack[len(m.menuStack)-1].cursor
		m.menuStack = m.menuStack[:len(m.menuStack)-1]
		if len(m.menuStack) == 0 {
			m.state = stateMainMenu
			m.cursor = m.rootCursor
		}
	}
}

func (m model) clampScroll() int {
	viewport := ui.ContentViewport(m.height)
	maxScroll := len(m.lines) - viewport
	if maxScroll < 0 {
		return 0
	}
	if m.scroll > maxScroll {
		return maxScroll
	}
	if m.scroll < 0 {
		return 0
	}
	return m.scroll
}

func (m model) breadcrumb() string {
	if m.currentSection == nil {
		return ""
	}
	parts := make([]string, len(m.menuStack))
	for i, f := range m.menuStack {
		parts[i] = f.title
	}
	parts = append(parts, m.currentSection.Title)
	return strings.Join(parts, "  ›  ")
}

func (m model) selectMenuItem(item data.MenuItem) model {
	if item.IsGroup() {
		m.enterGroup(item)
	} else {
		m.openSection(item.Section)
	}
	return m
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := (msg).(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q", "й":
			if m.state == stateMainMenu {
				return m, tea.Quit
			}
			m.goBack()

		case "enter":
			switch m.state {
			case stateMainMenu, stateSubMenu:
				items := m.activeItems()
				m = m.selectMenuItem(items[m.cursor])
			case stateContent:
				m.goBack()
			}

		case "esc":
			if m.state == stateMainMenu {
				return m, tea.Quit
			}
			m.goBack()

		case "up", "k":
			if m.state == stateContent {
				if m.scroll > 0 {
					m.scroll--
				}
			} else if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.state == stateContent {
				m.scroll++
				m.scroll = m.clampScroll()
			} else {
				items := m.activeItems()
				if m.cursor < len(items)-1 {
					m.cursor++
				}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {
	case stateMainMenu:
		return ui.RenderMenu(ui.MenuParams{
			Title:  "Инструменты для DevOps",
			Items:  m.rootMenu,
			Cursor: m.cursor,
			Width:  m.width,
			Height: m.height,
		})

	case stateSubMenu:
		return ui.RenderMenu(ui.MenuParams{
			Title:    m.menuTitle(),
			Subtitle: m.menuSubtitle(),
			Items:    m.activeItems(),
			Cursor:   m.cursor,
			Width:    m.width,
			Height:   m.height,
		})
	}

	if len(m.lines) == 0 && m.currentSection != nil {
		m.lines = ui.BuildSectionLines(m.currentSection)
	}

	return ui.RenderContent(ui.ContentParams{
		Title:      m.currentSection.Title,
		Breadcrumb: m.breadcrumb(),
		Lines:      m.lines,
		Scroll:     m.clampScroll(),
		Width:      m.width,
		Height:     m.height,
	})
}

func main() {
	p := tea.NewProgram(newModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

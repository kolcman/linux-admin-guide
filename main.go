package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"linux-guide/data"
)

// --------------------------------------------------
// Цветовая схема (ANSI)
// Совместима с SSH, tmux, screen и обычными терминалами
// --------------------------------------------------

var (
	// Сброс
	styleReset = "\033[0m"

	// Заголовок приложения
	styleTitle = "\033[1;94m" // Ярко-синий

	// Подзаголовки разделов
	styleHeader = "\033[1;93m" // Ярко-желтый

	// Пункты меню
	styleItem = "\033[97m" // Белый

	// Выбранный пункт
	styleCursor = "\033[30;46m" // Черный на голубом фоне

	// Команды
	styleCmdKey = "\033[1;92m" // Ярко-зеленый

	// Описание
	styleDesc = "\033[90m" // Серый

	// Ключи и параметры
	styleKeyFlag = "\033[1;96m" // Яркий голубой
)

type model struct {
	cursor         int
	state          int
	currentSection *data.Section
	menuItems      []string
}

func newModel() model {
	m := model{}

	sections := data.GetAllSections()

	for i := 1; i <= len(sections); i++ {
		if sec, ok := sections[i]; ok {
			m.menuItems = append(m.menuItems, sec.Title)
		}
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.state == 0 && m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.state == 0 && m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}

		case "enter":
			if m.state == 0 {
				if sec, ok := data.GetAllSections()[m.cursor+1]; ok {
					m.currentSection = &sec
					m.state = 1
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
		}
	}

	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	if m.state == 0 {

		s.WriteString(fmt.Sprintf(
			"%s\n  📖 LINUX ADMIN GUIDE\n%s\n\n",
			styleTitle,
			styleReset,
		))

		for i, item := range m.menuItems {

			cursor := " "

			itemStyle := styleItem

			if m.cursor == i {
				cursor = "▶"
				itemStyle = styleCursor
			}

			s.WriteString(fmt.Sprintf(
				"  %s %s%s%s\n",
				cursor,
				itemStyle,
				item,
				styleReset,
			))
		}

		s.WriteString("\n  ↑↓ / j k • Enter — открыть • q — выход")

	} else {

		line := strings.Repeat("═", 60)

		s.WriteString(fmt.Sprintf(
			"%s%s\n  %s%s\n%s\n\n",
			styleTitle,
			line,
			m.currentSection.Title,
			styleReset,
			line,
		))

		for _, item := range m.currentSection.Items {

			switch item.Type {

			case data.TypeHeader:

				s.WriteString(fmt.Sprintf(
					"\n%s  %s%s\n",
					styleHeader,
					item.Value,
					styleReset,
				))

			case data.TypeCmd:

				s.WriteString(fmt.Sprintf(
					"  %s$ %s%s\n     %s│ %s%s\n\n",
					styleCmdKey,
					item.Value,
					styleReset,
					styleDesc,
					item.Desc,
					styleReset,
				))

			case data.TypeKey:

				s.WriteString(fmt.Sprintf(
					"     %s• %s%s %s\n",
					styleKeyFlag,
					item.Key,
					styleReset,
					item.Desc,
				))

			case data.TypeVar:

				s.WriteString(fmt.Sprintf(
					"     %s• %s%s %s\n",
					styleCmdKey,
					item.Key,
					styleReset,
					item.Desc,
				))

			case data.TypeYaml:

				s.WriteString(fmt.Sprintf(
					"     %s• %s%s %s\n",
					styleKeyFlag,
					item.Key,
					styleReset,
					item.Desc,
				))
			}
		}

		s.WriteString("\n  Esc / Enter — назад")
	}

	return s.String()
}

func main() {
	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
}

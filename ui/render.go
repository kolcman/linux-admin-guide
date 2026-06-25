package ui

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
	"linux-guide/data"
)

var styles = NewStyles()

const (
	headerLines = 3
	footerLines = 2
	minWidth    = 52
)

func ContentViewport(height int) int {
	viewport := height - headerLines - footerLines
	if viewport < 8 {
		viewport = 8
	}
	return viewport
}

type MenuParams struct {
	Title    string
	Subtitle string
	Items    []data.MenuItem
	Cursor   int
	Width    int
	Height   int
}

type ContentParams struct {
	Title      string
	Breadcrumb string
	Lines      []string
	Scroll     int
	Width      int
	Height     int
}

func panelWidth(width int) int {
	if width < minWidth {
		return minWidth
	}
	return width - 4
}

// innerWidth — ширина текста внутри панели (без рамки и padding).
func innerWidth(width int) int {
	inner := panelWidth(width) - 6
	if inner < 24 {
		return 24
	}
	return inner
}

func hint(parts ...string) string {
	var b strings.Builder
	for i, part := range parts {
		if i > 0 {
			b.WriteString("  ")
		}
		b.WriteString(part)
	}
	return styles.Footer.Render(b.String())
}

func keyHint(key, label string) string {
	return styles.FooterKey.Render(key) + styles.Footer.Render(" "+label)
}

func menuFooter() string {
	return hint(
		keyHint("↑↓", "nav"),
		keyHint("enter", "open"),
		keyHint("esc", "back"),
		keyHint("q", "quit"),
	)
}

func contentFooter(scroll, total, viewport int) string {
	scrollPart := ""
	if total > viewport {
		at := scroll + 1
		if at < 1 {
			at = 1
		}
		end := scroll + viewport
		if end > total {
			end = total
		}
		scrollPart = styles.ScrollHint.Render(fmt.Sprintf("%d–%d / %d", at, end, total)) + "  "
	}

	return scrollPart + hint(
		keyHint("↑↓", "scroll"),
		keyHint("esc", "back"),
		keyHint("q", "back"),
	)
}

func renderHeader(title, subtitle string, width int) string {
	titleLine := styles.AppTitle.Render(title)
	if subtitle != "" {
		titleLine += " " + styles.AppSubtitle.Render(subtitle)
	}
	div := styles.Divider.Render(strings.Repeat("─", innerWidth(width)))
	return titleLine + "\n" + div + "\n"
}

func renderMenuItem(text string, active bool) string {
	width := runewidth.StringWidth(text)
	var b strings.Builder

	if active {
		b.WriteString(styles.MenuActive.Render(text))
	} else {
		b.WriteString(styles.MenuNormal.Render(text))
	}

	b.WriteByte('\n')

	if active {
		b.WriteString(styles.MenuActiveLine.Render(strings.Repeat("─", width)))
	} else {
		b.WriteString(strings.Repeat(" ", width))
	}

	return b.String()
}

func RenderMenu(p MenuParams) string {
	w := panelWidth(p.Width)

	var body strings.Builder
	body.WriteString(renderHeader(p.Title, p.Subtitle, p.Width))

	for i, item := range p.Items {
		body.WriteString(renderMenuItem("  "+item.Title, i == p.Cursor))
		body.WriteByte('\n')
	}

	panel := styles.Panel.Width(w).Render(body.String())
	footer := menuFooter()
	return lipglossJoinVertical(panel, footer)
}

func RenderContent(p ContentParams) string {
	w := panelWidth(p.Width)
	viewport := ContentViewport(p.Height)

	lines := p.Lines
	if len(lines) == 0 {
		lines = []string{styles.Desc.Render("  Пустой раздел")}
	}

	maxScroll := len(lines) - viewport
	if maxScroll < 0 {
		maxScroll = 0
	}

	scroll := p.Scroll
	if scroll < 0 {
		scroll = 0
	}
	if scroll > maxScroll {
		scroll = maxScroll
	}

	end := scroll + viewport
	if end > len(lines) {
		end = len(lines)
	}

	var header strings.Builder
	header.WriteString(renderHeader("Инструменты для DevOps", "", p.Width))

	if p.Breadcrumb != "" {
		header.WriteString(styles.Breadcrumb.Render("  "+p.Breadcrumb) + "\n")
	}
	header.WriteString(styles.SectionTitle.Render("  "+p.Title) + "\n")

	body := strings.Join(lines[scroll:end], "\n")
	panel := styles.Panel.Width(w).Render(header.String() + "\n" + body)
	footer := contentFooter(scroll, len(lines), viewport)

	return lipglossJoinVertical(panel, footer)
}

func lipglossJoinVertical(parts ...string) string {
	return strings.Join(parts, "\n")
}

func BuildSectionLines(section *data.Section) []string {
	if section == nil {
		return nil
	}

	var out []string

	for _, item := range section.Items {
		switch item.Type {

		case data.TypeHeader:
			out = append(out, "", styles.BlockHeader.Render(item.Value))

		case data.TypeTip:
			out = append(out, styles.Tip.Render("💡  "+item.Value))

		case data.TypeWarn:
			out = append(out, styles.Warn.Render("⚠️  "+item.Value))

		case data.TypeCmd:
			out = append(out, renderCommandLine(item.Value))
			if item.Desc != "" {
				out = append(out, styles.Desc.Render("→  "+item.Desc))
			}

		case data.TypeKey:
			out = append(out, renderKeyLine(item.Key, item.Desc))
		}
	}

	return out
}

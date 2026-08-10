package ui

import "github.com/charmbracelet/lipgloss"

// Адаптивная палитра: Light — для белого/светлого фона терминала,
// Dark — для тёмного. Иначе на дневной теме почти белый текст не читается.
var (
	colorPrimary     = lipgloss.AdaptiveColor{Light: "25", Dark: "111"}  // синий акцент
	colorSecondary   = lipgloss.AdaptiveColor{Light: "90", Dark: "177"}  // заголовки
	colorText        = lipgloss.AdaptiveColor{Light: "232", Dark: "254"} // основной текст
	colorMuted       = lipgloss.AdaptiveColor{Light: "238", Dark: "245"} // подписи / desc
	colorCmdVerb     = lipgloss.AdaptiveColor{Light: "17", Dark: "252"}  // команда
	colorCmdArg      = lipgloss.AdaptiveColor{Light: "23", Dark: "117"}  // аргументы
	colorFlag        = lipgloss.AdaptiveColor{Light: "130", Dark: "227"} // флаги -a, --name
	colorPlaceholder = lipgloss.AdaptiveColor{Light: "125", Dark: "218"} // <id>
	colorConcept     = lipgloss.AdaptiveColor{Light: "54", Dark: "141"}  // подкоманды
	colorTip         = lipgloss.AdaptiveColor{Light: "236", Dark: "250"}
	colorWarn        = lipgloss.AdaptiveColor{Light: "160", Dark: "209"}
	colorBorder      = lipgloss.AdaptiveColor{Light: "242", Dark: "238"}
	colorDim         = lipgloss.AdaptiveColor{Light: "241", Dark: "244"}
	colorPrompt      = lipgloss.AdaptiveColor{Light: "22", Dark: "86"} // $
)

type Styles struct {
	AppTitle       lipgloss.Style
	AppSubtitle    lipgloss.Style
	Panel          lipgloss.Style
	MenuNormal     lipgloss.Style
	MenuActive     lipgloss.Style
	MenuActiveLine lipgloss.Style
	Footer         lipgloss.Style
	FooterKey      lipgloss.Style
	Breadcrumb     lipgloss.Style
	SectionTitle   lipgloss.Style
	BlockHeader    lipgloss.Style
	CmdVerb        lipgloss.Style
	CmdArg         lipgloss.Style
	CmdPrompt      lipgloss.Style
	Desc           lipgloss.Style
	KeyFlag        lipgloss.Style
	KeyConcept     lipgloss.Style
	KeyLabel       lipgloss.Style
	Placeholder    lipgloss.Style
	Tip            lipgloss.Style
	Warn           lipgloss.Style
	ScrollHint     lipgloss.Style
	Divider        lipgloss.Style
}

func NewStyles() Styles {
	return Styles{
		AppTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary),

		AppSubtitle: lipgloss.NewStyle().
			Foreground(colorMuted).
			Italic(true),

		Panel: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(1, 2),

		MenuNormal: lipgloss.NewStyle().
			Foreground(colorText).
			Padding(0, 1),

		MenuActive: lipgloss.NewStyle().
			Foreground(colorPrimary).
			Bold(true).
			Padding(0, 1),

		MenuActiveLine: lipgloss.NewStyle().
			Foreground(colorPrimary).
			Padding(0, 1),

		Footer: lipgloss.NewStyle().
			Foreground(colorMuted),

		FooterKey: lipgloss.NewStyle().
			Foreground(colorPrimary).
			Bold(true),

		Breadcrumb: lipgloss.NewStyle().
			Foreground(colorMuted),

		SectionTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(colorSecondary),

		BlockHeader: lipgloss.NewStyle().
			Bold(true).
			Foreground(colorSecondary).
			MarginTop(1),

		CmdVerb: lipgloss.NewStyle().
			Bold(true).
			Foreground(colorCmdVerb),

		CmdArg: lipgloss.NewStyle().
			Foreground(colorCmdArg),

		CmdPrompt: lipgloss.NewStyle().
			Foreground(colorPrompt).
			Bold(true),

		Desc: lipgloss.NewStyle().
			Foreground(colorMuted).
			PaddingLeft(4),

		KeyFlag: lipgloss.NewStyle().
			Foreground(colorFlag).
			Bold(true),

		KeyConcept: lipgloss.NewStyle().
			Foreground(colorConcept).
			Bold(true),

		KeyLabel: lipgloss.NewStyle().
			Foreground(colorMuted),

		Placeholder: lipgloss.NewStyle().
			Foreground(colorPlaceholder).
			Italic(true),

		Tip: lipgloss.NewStyle().
			Foreground(colorTip).
			Italic(true).
			PaddingLeft(2),

		Warn: lipgloss.NewStyle().
			Foreground(colorWarn).
			Bold(true).
			PaddingLeft(2),

		ScrollHint: lipgloss.NewStyle().
			Foreground(colorDim),

		Divider: lipgloss.NewStyle().
			Foreground(colorBorder),
	}
}

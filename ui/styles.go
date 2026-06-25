package ui

import "github.com/charmbracelet/lipgloss"

// Палитра в духе продуктового CLI (256-color, SSH-safe).
const (
	colorPrimary     = "111" // sky blue — бренд, активное меню
	colorSecondary   = "177" // lilac — заголовки блоков
	colorText        = "254"
	colorMuted       = "243"
	colorCmdVerb     = "252" // команда (docker, run)
	colorCmdArg      = "117" // аргументы (nginx, пути)
	colorFlag        = "227" // жёлтые флаги -a, --name
	colorPlaceholder = "218" // <id>, <image>
	colorConcept     = "141" // подкоманды: up, prune, bridge
	colorTip         = "245"
	colorWarn        = "209"
	colorBorder      = "238"
	colorDim         = "240"
	colorPrompt      = "86" // $ prompt
)

type Styles struct {
	AppTitle     lipgloss.Style
	AppSubtitle  lipgloss.Style
	Panel        lipgloss.Style
	MenuNormal     lipgloss.Style
	MenuActive     lipgloss.Style
	MenuActiveLine lipgloss.Style
	Footer         lipgloss.Style
	FooterKey    lipgloss.Style
	Breadcrumb   lipgloss.Style
	SectionTitle lipgloss.Style
	BlockHeader  lipgloss.Style
	CmdVerb      lipgloss.Style
	CmdArg       lipgloss.Style
	CmdPrompt    lipgloss.Style
	Desc         lipgloss.Style
	KeyFlag      lipgloss.Style
	KeyConcept   lipgloss.Style
	KeyLabel     lipgloss.Style
	Placeholder  lipgloss.Style
	Tip          lipgloss.Style
	Warn         lipgloss.Style
	ScrollHint   lipgloss.Style
	Divider      lipgloss.Style
}

func NewStyles() Styles {
	return Styles{
		AppTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorPrimary)),

		AppSubtitle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			Italic(true),

		Panel: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(colorBorder)).
			Padding(1, 2),

		MenuNormal: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorText)).
			Padding(0, 1),

		MenuActive: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPrimary)).
			Bold(true).
			Padding(0, 1),

		MenuActiveLine: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPrimary)).
			Padding(0, 1),

		Footer: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)),

		FooterKey: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPrimary)).
			Bold(true),

		Breadcrumb: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)),

		SectionTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorSecondary)),

		BlockHeader: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorSecondary)).
			MarginTop(1),

		CmdVerb: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorCmdVerb)),

		CmdArg: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorCmdArg)),

		CmdPrompt: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPrompt)).
			Bold(true),

		Desc: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			PaddingLeft(4),

		KeyFlag: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorFlag)).
			Bold(true),

		KeyConcept: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorConcept)).
			Bold(true),

		KeyLabel: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)),

		Placeholder: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPlaceholder)).
			Italic(true),

		Tip: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorTip)).
			Italic(true).
			PaddingLeft(2),

		Warn: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWarn)).
			Bold(true).
			PaddingLeft(2),

		ScrollHint: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorDim)),

		Divider: lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBorder)),
	}
}

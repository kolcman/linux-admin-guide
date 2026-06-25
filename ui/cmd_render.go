package ui

import "strings"

var composeSubcommands = map[string]bool{
	"up": true, "down": true, "ps": true, "logs": true, "build": true,
	"pull": true, "restart": true, "stop": true, "start": true, "exec": true,
	"config": true, "run": true,
}

func renderCommandLine(cmd string) string {
	var b strings.Builder
	b.WriteString("  ")
	b.WriteString(styles.CmdPrompt.Render("$"))
	b.WriteByte(' ')

	parts := strings.Fields(cmd)
	for i, part := range parts {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(renderCmdToken(part, i, parts))
	}
	return b.String()
}

func renderCmdToken(part string, index int, parts []string) string {
	if strings.HasPrefix(part, "-") {
		return styles.KeyFlag.Render(part)
	}
	if strings.HasPrefix(part, "<") && strings.HasSuffix(part, ">") {
		return styles.Placeholder.Render(part)
	}
	if isComposeSubcommand(parts, index) {
		return styles.KeyConcept.Render(part)
	}
	if strings.Contains(part, "=") {
		return renderKVToken(part)
	}
	if strings.Contains(part, ":") && shouldRenderAsMount(part) {
		return renderMountToken(part)
	}
	if index == 0 {
		return styles.CmdVerb.Render(part)
	}
	return styles.CmdArg.Render(part)
}

func isComposeSubcommand(parts []string, index int) bool {
	if index < 2 || parts[0] != "docker" || parts[1] != "compose" {
		return false
	}
	return composeSubcommands[parts[index]]
}

func renderKVToken(s string) string {
	chunks := strings.Split(s, ",")
	var b strings.Builder
	for i, chunk := range chunks {
		if i > 0 {
			b.WriteString(styles.CmdArg.Render(","))
		}
		key, val, ok := strings.Cut(chunk, "=")
		if !ok {
			b.WriteString(styles.CmdArg.Render(chunk))
			continue
		}
		b.WriteString(styles.KeyFlag.Render(key))
		b.WriteString(styles.CmdArg.Render("=" + val))
	}
	return b.String()
}

func renderMountToken(s string) string {
	parts := strings.Split(s, ":")
	var b strings.Builder
	for i, p := range parts {
		if i > 0 {
			b.WriteString(styles.KeyFlag.Render(":"))
		}
		switch p {
		case "ro", "rw":
			b.WriteString(styles.KeyFlag.Render(p))
		default:
			b.WriteString(styles.CmdArg.Render(p))
		}
	}
	return b.String()
}

func shouldRenderAsMount(s string) bool {
	if strings.Contains(s, "/") {
		return true
	}
	if strings.HasSuffix(s, ":ro") || strings.HasSuffix(s, ":rw") {
		return true
	}
	parts := strings.Split(s, ":")
	if len(parts) >= 3 {
		return true
	}
	if len(parts) == 2 {
		return isPortOrPathPair(parts[0], parts[1])
	}
	return false
}

func isPortOrPathPair(a, b string) bool {
	for _, c := range a + b {
		if (c < '0' || c > '9') && c != '.' {
			return false
		}
	}
	return true
}

func renderKeyLine(key, desc string) string {
	style := styles.KeyConcept
	if isFlagKey(key) {
		style = styles.KeyFlag
	}
	return "     " + style.Render(key) + styles.KeyLabel.Render("   "+desc)
}

func isFlagKey(key string) bool {
	return strings.HasPrefix(key, "-") || strings.HasPrefix(key, ":")
}

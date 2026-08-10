package data

var PackagesSection = Section{
	Title: "📦 МЕНЕДЖЕРЫ ПАКЕТОВ",
	Items: []Item{
		{Type: TypeTip, Value: "Пакетный менеджер — как ставят и обновляют ПО на сервере.\nAPT — Ubuntu/Debian, DNF — Fedora/RHEL-клоны, Snap — универсальные пакеты."},

		{Type: TypeHeader, Value: "📦 APT (Ubuntu / Debian)"},
		{Type: TypeCmd, Value: "sudo apt update", Desc: "обновить индексы пакетов (сделай перед install)"},
		{Type: TypeCmd, Value: "sudo apt upgrade", Desc: "обновить уже установленные пакеты"},
		{Type: TypeCmd, Value: "sudo apt install <pkg>", Desc: "установить пакет"},
		{Type: TypeCmd, Value: "sudo apt remove <pkg>", Desc: "удалить пакет (конфиги часто остаются)"},
		{Type: TypeCmd, Value: "sudo apt purge <pkg>", Desc: "удалить пакет вместе с конфигами"},
		{Type: TypeCmd, Value: "apt search <name>", Desc: "поиск по имени или описанию"},
		{Type: TypeCmd, Value: "apt show <pkg>", Desc: "описание пакета, версия, зависимости"},
		{Type: TypeCmd, Value: "apt list --installed", Desc: "список установленных пакетов"},

		{Type: TypeHeader, Value: "📦 DNF (Fedora / RHEL-клон)"},
		{Type: TypeCmd, Value: "sudo dnf check-update", Desc: "проверить обновления"},
		{Type: TypeCmd, Value: "sudo dnf install <pkg>", Desc: "установить пакет"},
		{Type: TypeCmd, Value: "sudo dnf remove <pkg>", Desc: "удалить пакет"},
		{Type: TypeCmd, Value: "dnf search <name>", Desc: "поиск пакета"},
		{Type: TypeCmd, Value: "sudo dnf groupinstall \"Development Tools\"", Desc: "установить группу пакетов"},

		{Type: TypeHeader, Value: "📦 Snap"},
		{Type: TypeCmd, Value: "sudo snap install <pkg>", Desc: "установить snap-пакет"},
		{Type: TypeCmd, Value: "sudo snap refresh", Desc: "обновить все snap-пакеты"},
		{Type: TypeCmd, Value: "snap list", Desc: "список установленных snap"},
		{Type: TypeCmd, Value: "sudo snap remove <pkg>", Desc: "удалить snap-пакет"},

		{Type: TypeHeader, Value: "💡 Советы"},
		{Type: TypeTip, Value: "Всегда сначала update/check-update, потом install — иначе можешь поставить старую версию из кэша."},
		{Type: TypeWarn, Value: "purge удаляет конфиги. Если не уверен — remove, конфиги останутся в /etc."},
	},
}

package data

var PackagesSection = Section{
	Title: "📦 МЕНЕДЖЕРЫ ПАКЕТОВ",
	Items: []Item{
		{Type: TypeTip, Value: "Пакет — программа + служебные файлы (куда ставить, скрипты, версия, мейнтейнер).\nUbuntu/Debian — .deb. Fedora/RHEL — .rpm. Репозиторий — внешняя база; система ищет по скачанному индексу (кэшу)."},

		{Type: TypeHeader, Value: "🧰 Кто чем занимается (deb)"},
		{Type: TypeKey, Key: "dpkg", Desc: "локальные .deb и установленные пакеты, без репозитория"},
		{Type: TypeKey, Key: "apt-cache", Desc: "поиск и версии по локальному кэшу"},
		{Type: TypeKey, Key: "apt-get", Desc: "репозиторий: update, install, upgrade, remove"},
		{Type: TypeKey, Key: "apt", Desc: "удобная оболочка над apt-get + apt-cache"},
		{Type: TypeKey, Key: "/etc/apt/sources.list", Desc: "откуда качать пакеты (+ sources.list.d/)"},

		{Type: TypeHeader, Value: "📦 dpkg — локальный .deb"},
		{Type: TypeKey, Key: "-l", Desc: "список установленных"},
		{Type: TypeKey, Key: "-s", Desc: "инфа о пакете: мейнтейнер, версия, статус"},
		{Type: TypeKey, Key: "-L", Desc: "файлы, принадлежащие пакету"},
		{Type: TypeKey, Key: "-i", Desc: "установить уже скачанный .deb"},
		{Type: TypeKey, Key: "-r", Desc: "удалить пакет (имя пакета, не файл .deb)"},
		{Type: TypeCmd, Value: "dpkg -l", Desc: "какие пакеты установлены"},
		{Type: TypeCmd, Value: "dpkg -s <pkg>", Desc: "мейнтейнер, версия, статус"},
		{Type: TypeCmd, Value: "dpkg -L <pkg>", Desc: "файлы пакета в системе"},
		{Type: TypeCmd, Value: "sudo dpkg -i program.deb", Desc: "поставить скачанный .deb"},
		{Type: TypeCmd, Value: "sudo dpkg -r <pkg>", Desc: "удалить пакет"},
		{Type: TypeWarn, Value: "dpkg -i не тянет зависимости из репо. Для обычной установки — apt."},

		{Type: TypeHeader, Value: "📦 apt-get"},
		{Type: TypeCmd, Value: "sudo apt-get update", Desc: "скачать свежий индекс пакетов"},
		{Type: TypeCmd, Value: "sudo apt-get install <pkg>", Desc: "поставить; если стоит — обновит. Несколько имён через пробел"},
		{Type: TypeCmd, Value: "sudo apt-get install <pkg>=<ver>", Desc: "конкретная версия (смотри в policy)"},
		{Type: TypeCmd, Value: "sudo apt-get upgrade", Desc: "обновить все установленные пакеты"},
		{Type: TypeCmd, Value: "sudo apt-get remove --purge <pkg>", Desc: "удалить вместе с конфигами"},
		{Type: TypeCmd, Value: "apt-get changelog <pkg>", Desc: "лог изменений по версиям"},

		{Type: TypeHeader, Value: "🔎 apt-cache"},
		{Type: TypeCmd, Value: "apt-cache search <name>", Desc: "найти пакет по имени программы"},
		{Type: TypeCmd, Value: "apt-cache policy <pkg>", Desc: "версии в репо, что установлено, откуда ставится"},
		{Type: TypeCmd, Value: "apt-cache pkgnames", Desc: "все доступные имена (огромный список)"},

		{Type: TypeHeader, Value: "📦 APT (Ubuntu / Debian)"},
		{Type: TypeCmd, Value: "sudo apt update", Desc: "обновить индексы (сделай перед install)"},
		{Type: TypeCmd, Value: "sudo apt upgrade", Desc: "обновить установленные пакеты"},
		{Type: TypeCmd, Value: "sudo apt install <pkg>", Desc: "установить пакет"},
		{Type: TypeCmd, Value: "sudo apt remove <pkg>", Desc: "удалить пакет (конфиги часто остаются)"},
		{Type: TypeCmd, Value: "sudo apt purge <pkg>", Desc: "удалить вместе с конфигами"},
		{Type: TypeCmd, Value: "apt search <name>", Desc: "поиск по имени или описанию"},
		{Type: TypeCmd, Value: "apt show <pkg>", Desc: "описание, версия, зависимости"},
		{Type: TypeCmd, Value: "apt list --installed", Desc: "установленные пакеты"},
		{Type: TypeCmd, Value: "apt policy <pkg>", Desc: "версии в репо и что сейчас стоит"},

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
		{Type: TypeTip, Value: "Сначала update/check-update, потом install — иначе может встать старая версия из кэша."},
		{Type: TypeWarn, Value: "purge удаляет конфиги. Если не уверен — remove, конфиги останутся в /etc."},
	},
}

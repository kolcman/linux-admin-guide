package data

var PackagesSection = Section{
	Title: "📦 МЕНЕДЖЕРЫ ПАКЕТОВ",
	Items: []Item{
		{Type: TypeHeader, Value: "🔑 APT (Ubuntu):"},
		{Type: TypeKey, Key: "update", Desc: "Обновить список пакетов"},
		{Type: TypeKey, Key: "purge", Desc: "Удалить пакет с конфигами"},

		{Type: TypeHeader, Value: "🔑 SNAP:"},
		{Type: TypeKey, Key: "refresh", Desc: "Обновить все snap-пакеты"},

		{Type: TypeHeader, Value: "🔑 DNF (Fedora):"},
		{Type: TypeKey, Key: "groupinstall", Desc: "Установить группу пакетов"},
	},
}

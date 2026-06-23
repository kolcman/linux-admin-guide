package data

var NetworkSection = Section{
	Title: "🌐 СЕТЬ",
	Items: []Item{
		{Type: TypeCmd, Value: "ip -br -c a", Desc: "IP адреса интерфейсов"},
		{Type: TypeCmd, Value: "ss -tulpn", Desc: "Открытые порты и процессы"},
		{Type: TypeCmd, Value: "ping -c 4 google.com", Desc: "Проверка связи"},
	},
}

package data

var ServicesSection = Section{
	Title: "🔧 SYSTEMD",
	Items: []Item{
		{Type: TypeCmd, Value: "systemctl status service", Desc: "Статус службы"},
		{Type: TypeCmd, Value: "sudo journalctl -u service -f", Desc: "Логи в реальном времени"},
		{Type: TypeCmd, Value: "systemctl --failed", Desc: "Список упавших служб"},
	},
}

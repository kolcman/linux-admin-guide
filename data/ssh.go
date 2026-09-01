package data

var SSHSection = Section{
	Title: "🔐 SSH — ТУННЕЛИ И АДМИНИСТРИРОВАНИЕ",
	Items: []Item{
		{Type: TypeHeader, Value: "🔑 Флаги команды ssh:"},
		{Type: TypeKey, Key: "-p <port>", Desc: "Подключение по нестандартному порту"},
		{Type: TypeKey, Key: "-i <path>", Desc: "Использовать конкретный приватный ключ"},
		{Type: TypeKey, Key: "-v", Desc: "Режим отладки (показывает процесс handshake)"},
		{Type: TypeKey, Key: "-N", Desc: "Не выполнять команду (только туннель)"},
		{Type: TypeKey, Key: "-f", Desc: "Уйти в фон после подключения"},
		{Type: TypeCmd, Value: "ssh -p 2222 user@host", Desc: "порт не 22"},
		{Type: TypeCmd, Value: "ssh -i ~/.ssh/id_ed25519 user@host", Desc: "конкретный ключ"},
		{Type: TypeCmd, Value: "ssh -v user@host", Desc: "отладка рукопожатия"},

		{Type: TypeHeader, Value: "🚇 Туннелирование:"},
		{Type: TypeKey, Key: "-L <loc>:<rem>", Desc: "Локальный туннель (доступ к БД сервера)"},
		{Type: TypeKey, Key: "-R <rem>:<loc>", Desc: "Обратный туннель (доступ к своему ПК)"},
		{Type: TypeKey, Key: "-D <port>", Desc: "SOCKS прокси через сервер"},
		{Type: TypeCmd, Value: "ssh -N -L 5432:localhost:5432 user@host", Desc: "локальный туннель к БД"},
		{Type: TypeCmd, Value: "ssh -N -R 8080:localhost:80 user@host", Desc: "обратный туннель"},
		{Type: TypeCmd, Value: "ssh -N -D 1080 -f user@host", Desc: "SOCKS в фоне"},

		{Type: TypeHeader, Value: "🛠️ Управление"},
		{Type: TypeCmd, Value: "ssh-keygen -t ed25519 -C 'my_pc'", Desc: "Генерация современного ключа"},
		{Type: TypeCmd, Value: "ssh-copy-id user@host", Desc: "Быстрая отправка ключа на сервер"},
		{Type: TypeCmd, Value: "sudo systemctl restart sshd", Desc: "Перезапуск службы после правки конфига"},

		{Type: TypeHeader, Value: "📄 Конфиг клиента (~/.ssh/config)"},
		{Type: TypeTip, Value: "Host myserver\n  HostName 203.0.113.10\n  User admin\n  Port 2222\n  IdentityFile ~/.ssh/id_ed25519\n\nПотом просто: ssh myserver"},
		{Type: TypeWarn, Value: "На сервере в /etc/ssh/sshd_config не оставляй PasswordAuthentication yes и PermitRootLogin yes на продакшене без нужды."},
	},
}

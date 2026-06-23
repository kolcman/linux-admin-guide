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

		{Type: TypeHeader, Value: "🚇 Туннелирование:"},
		{Type: TypeKey, Key: "-L <loc>:<rem>", Desc: "Локальный туннель (доступ к БД сервера)"},
		{Type: TypeKey, Key: "-R <rem>:<loc>", Desc: "Обратный туннель (доступ к своему ПК)"},
		{Type: TypeKey, Key: "-D <port>", Desc: "SOCKS прокси через сервер"},

		{Type: TypeHeader, Value: "🛠️ Управление:"},
		{Type: TypeCmd, Value: "ssh-keygen -t ed25519 -C 'my_pc'", Desc: "Генерация современного ключа"},
		{Type: TypeCmd, Value: "ssh-copy-id user@host", Desc: "Быстрая отправка ключа на сервер"},
		{Type: TypeCmd, Value: "sudo systemctl restart sshd", Desc: "Перезапуск службы после правки конфига"},
	},
}

package data

var DockerComposeSection = Section{
	Title: "🐳 DOCKER COMPOSE — СТЕКИ",
	Items: []Item{
		{Type: TypeHeader, Value: "🚀 Основные команды:"},
		{Type: TypeCmd, Value: "docker compose up -d", Desc: "Запустить стек в фоне"},
		{Type: TypeCmd, Value: "docker compose down", Desc: "Остановить и удалить контейнеры/сети"},
		{Type: TypeCmd, Value: "docker compose logs -f", Desc: "Логи всех сервисов сразу"},

		{Type: TypeHeader, Value: "🔑 Ключи 'up':"},
		{Type: TypeKey, Key: "--build", Desc: "Пересобрать образы перед запуском"},
		{Type: TypeKey, Key: "-f file.yml", Desc: "Использовать другой файл конфигурации"},

		{Type: TypeHeader, Value: "🔑 Ключи 'down' (ОСТОРОЖНО):"},
		{Type: TypeKey, Key: "-v", Desc: "Удалить также тома с данными (БД!"},
		{Type: TypeKey, Key: "--rmi all", Desc: "Удалить все образы этого стека"},
	},
}

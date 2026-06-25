package docker

var CleanupSection = Section{
	Title: "🧹 ОЧИСТКА",
	Items: []Item{
		{Type: TypeTip, Value: "Docker копит остановленные контейнеры, старые образы и неиспользуемые сети — они занимают диск."},
		{Type: TypeTip, Value: "Начинай с безопасных команд prune. Полную очистку — только если понимаешь последствия."},

		{Type: TypeHeader, Value: "🧹 Безопасная очистка"},
		{Type: TypeKey, Key: "prune", Desc: "удалить только неиспользуемое, рабочее не трогает"},
		{Type: TypeKey, Key: "-f", Desc: "без интерактивного подтверждения"},
		{Type: TypeCmd, Value: "docker container prune -f", Desc: "удалить остановленные контейнеры"},
		{Type: TypeCmd, Value: "docker image prune -f", Desc: "удалить образы без тега (dangling)"},
		{Type: TypeCmd, Value: "docker volume prune -f", Desc: "тома без привязанных контейнеров"},
		{Type: TypeCmd, Value: "docker network prune -f", Desc: "сети без контейнеров"},

		{Type: TypeHeader, Value: "⚠️ Полная очистка"},
		{Type: TypeKey, Key: "-a", Desc: "all — все неиспользуемые образы, не только dangling"},
		{Type: TypeKey, Key: "--volumes", Desc: "также удалить неиспользуемые тома с данными"},
		{Type: TypeCmd, Value: "docker system prune -f", Desc: "контейнеры + сети + dangling-образы"},
		{Type: TypeCmd, Value: "docker system prune -a -f", Desc: "все неиспользуемые образы"},
		{Type: TypeWarn, Value: "prune -a удалит образы — придётся заново docker pull"},
		{Type: TypeCmd, Value: "docker system prune -a --volumes -f", Desc: "максимальная очистка диска"},
		{Type: TypeWarn, Value: "--volumes удалит данные БД и файлы в неиспользуемых томах!"},
	},
}

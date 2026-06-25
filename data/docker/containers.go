package docker

var ContainersSection = Section{
	Title: "🔧 КОНТЕЙНЕРЫ",
	Items: []Item{
		{Type: TypeTip, Value: "Контейнер — запущенное приложение. Образ (image) — шаблон, контейнер — его копия в работе."},
		{Type: TypeTip, Value: "Удалили контейнер — данные внутри пропадут (если не подключён volume)."},

		{Type: TypeHeader, Value: "📋 Список и статус"},
		{Type: TypeKey, Key: "-a", Desc: "все контейнеры, включая остановленные"},
		{Type: TypeKey, Key: "-q", Desc: "вывести только ID (удобно для скриптов)"},
		{Type: TypeCmd, Value: "docker ps", Desc: "только запущенные контейнеры"},
		{Type: TypeCmd, Value: "docker ps -a", Desc: "все контейнеры на сервере"},
		{Type: TypeCmd, Value: "docker ps -a -q", Desc: "только ID всех контейнеров"},
		{Type: TypeCmd, Value: "docker inspect <id>", Desc: "подробная информация: сеть, volumes, переменные"},

		{Type: TypeHeader, Value: "🚀 Запуск"},
		{Type: TypeKey, Key: "-d", Desc: "detached — запуск в фоне, терминал свободен"},
		{Type: TypeKey, Key: "-it", Desc: "интерактивный режим + терминал (для bash, sh)"},
		{Type: TypeKey, Key: "--name", Desc: "дать контейнеру понятное имя"},
		{Type: TypeCmd, Value: "docker run -d --name web nginx", Desc: "nginx в фоне с именем web"},
		{Type: TypeCmd, Value: "docker run -it ubuntu bash", Desc: "войти в ubuntu и открыть shell"},
		{Type: TypeCmd, Value: "docker start <id>", Desc: "запустить уже созданный контейнер"},

		{Type: TypeHeader, Value: "⏹️ Остановка и удаление"},
		{Type: TypeKey, Key: "-f", Desc: "force — принудительно (stop + rm за раз)"},
		{Type: TypeCmd, Value: "docker stop <id>", Desc: "корректно остановить контейнер"},
		{Type: TypeCmd, Value: "docker rm <id>", Desc: "удалить остановленный контейнер"},
		{Type: TypeCmd, Value: "docker rm -f <id>", Desc: "остановить и сразу удалить"},

		{Type: TypeHeader, Value: "🔍 Логи и доступ внутрь"},
		{Type: TypeKey, Key: "-f", Desc: "follow — показывать новые строки лога в реальном времени"},
		{Type: TypeKey, Key: "--tail", Desc: "показать только последние N строк"},
		{Type: TypeCmd, Value: "docker logs -f <id>", Desc: "смотреть логи как tail -f"},
		{Type: TypeCmd, Value: "docker logs --tail 100 -f <id>", Desc: "последние 100 строк + follow"},
		{Type: TypeCmd, Value: "docker exec -it <id> bash", Desc: "открыть shell внутри работающего контейнера"},
	},
}

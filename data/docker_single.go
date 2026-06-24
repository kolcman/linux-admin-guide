package data

var DockerSingleSection = Section{
	Title: "🐳 DOCKER — ОСНОВЫ",

	Items: []Item{

		// =========================
		// INTRO
		// =========================
		{Type: TypeHeader, Value: "🧠 DOCKER — ЧТО ЭТО:"},

		{Type: TypeCmd, Value: "Docker", Desc: "Система для запуска приложений в изолированных контейнерах"},
		{Type: TypeCmd, Value: "СУТЬ", Desc: "Каждое приложение работает как отдельная изолированная среда"},

		// =========================
		// CONTAINERS
		// =========================
		{Type: TypeHeader, Value: "🛠️ КОНТЕЙНЕРЫ:"},

		{Type: TypeKey, Key: "NOTE", Desc: "Контейнер = запущенная программа. Удалил контейнер → программа исчезла"},

		{Type: TypeCmd, Value: "docker ps -a", Desc: "Показать все контейнеры"},
		{Type: TypeCmd, Value: "docker run -d nginx", Desc: "Запустить контейнер в фоне"},
		{Type: TypeCmd, Value: "docker stop <id>", Desc: "Остановить контейнер"},
		{Type: TypeCmd, Value: "docker rm <id>", Desc: "Удалить контейнер"},

		// =========================
		// IMAGES
		// =========================
		{Type: TypeHeader, Value: "📦 ОБРАЗЫ:"},

		{Type: TypeKey, Key: "NOTE", Desc: "Image = шаблон. Container = запущенная копия"},

		{Type: TypeCmd, Value: "docker images", Desc: "Список образов"},
		{Type: TypeCmd, Value: "docker pull nginx", Desc: "Скачать готовый образ"},
		{Type: TypeCmd, Value: "docker build -t myapp .", Desc: "Собрать свой образ"},

		// =========================
		// PORTS
		// =========================
		{Type: TypeHeader, Value: "🌐 ПОРТЫ:"},

		{Type: TypeKey, Key: "NOTE", Desc: "Порт = способ открыть контейнер наружу"},

		{Type: TypeCmd, Value: "docker run -p 8080:80 nginx", Desc: "localhost:8080 → контейнер:80"},
		{Type: TypeCmd, Value: "docker run -p 127.0.0.1:8080:80 nginx", Desc: "Доступ только локально"},

		{Type: TypeCmd, Value: "ВАЖНО", Desc: "Без -p контейнер недоступен извне"},

		// =========================
		// VOLUMES
		// =========================
		{Type: TypeHeader, Value: "📦 VOLUMES:"},

		{Type: TypeKey, Key: "NOTE", Desc: "Volume = папка на сервере для хранения данных вне контейнера"},

		{Type: TypeCmd, Value: "docker volume ls", Desc: "Показать все хранилища данных"},
		{Type: TypeCmd, Value: "docker volume create my_vol", Desc: "Создать volume"},
		{Type: TypeCmd, Value: "docker run -v my_vol:/data ubuntu", Desc: "Подключить volume"},

		{Type: TypeCmd, Value: "ПРИМЕР", Desc: "База данных: без volume данные исчезают, с volume сохраняются"},

		// =========================
		// NETWORK
		// =========================
		{Type: TypeHeader, Value: "🌐 NETWORK:"},

		{Type: TypeKey, Key: "NOTE", Desc: "Network = внутренняя сеть Docker для общения контейнеров"},

		{Type: TypeCmd, Value: "docker network ls", Desc: "Список сетей Docker"},
		{Type: TypeCmd, Value: "docker network create my_net", Desc: "Создать сеть"},
		{Type: TypeCmd, Value: "docker run --network my_net nginx", Desc: "Запустить контейнер в сети"},

		{Type: TypeCmd, Value: "ПРИМЕР", Desc: "Backend и DB общаются внутри одной сети"},

		{Type: TypeCmd, Value: "ВАЖНО", Desc: "Без сети контейнеры изолированы"},

		// =========================
		// CLEANUP
		// =========================
		{Type: TypeHeader, Value: "🧹 ОЧИСТКА:"},

		{Type: TypeCmd, Value: "docker system prune -a", Desc: "Удалить всё неиспользуемое"},

		{Type: TypeCmd, Value: "ОПАСНО", Desc: "Полная очистка Docker (контейнеры, образы, кеш)"},
	},
}
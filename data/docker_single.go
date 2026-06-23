package data

var DockerSingleSection = Section{
	Title: "🐳 DOCKER",

	Items: []Item{

		// --------------------------------------------------
		// Контейнеры
		// --------------------------------------------------
		{Type: TypeHeader, Value: "📦 Контейнеры"},

		{Type: TypeCmd, Value: "docker ps", Desc: "Запущенные контейнеры"},
		{Type: TypeCmd, Value: "docker ps -a", Desc: "Все контейнеры"},
		{Type: TypeCmd, Value: "docker run -d --name my_app -p 8080:80 nginx", Desc: "Запустить контейнер"},
		{Type: TypeCmd, Value: "docker start <container>", Desc: "Запустить контейнер"},
		{Type: TypeCmd, Value: "docker stop <container>", Desc: "Остановить контейнер"},
		{Type: TypeCmd, Value: "docker restart <container>", Desc: "Перезапустить контейнер"},
		{Type: TypeCmd, Value: "docker rm <container>", Desc: "Удалить контейнер"},
		{Type: TypeCmd, Value: "docker logs -f <container>", Desc: "Просмотр логов"},
		{Type: TypeCmd, Value: "docker inspect <container>", Desc: "Информация о контейнере"},
		{Type: TypeCmd, Value: "docker stats", Desc: "Использование CPU и памяти"},
		{Type: TypeCmd, Value: "docker top <container>", Desc: "Процессы контейнера"},
		{Type: TypeCmd, Value: "docker diff <container>", Desc: "Изменения файловой системы"},
		{Type: TypeCmd, Value: "docker port <container>", Desc: "Проброшенные порты"},

		{Type: TypeHeader, Value: "🔑 Ключи docker ps"},
		{Type: TypeKey, Key: "-a", Desc: "Показать все контейнеры"},
		{Type: TypeKey, Key: "-q", Desc: "Только ID контейнеров"},
		{Type: TypeKey, Key: "-s", Desc: "Размер контейнеров"},
		{Type: TypeKey, Key: "--format", Desc: "Настроить формат вывода"},

		{Type: TypeHeader, Value: "🔑 Ключи docker run"},
		{Type: TypeKey, Key: "-d", Desc: "Фоновый режим"},
		{Type: TypeKey, Key: "-it", Desc: "Интерактивный терминал"},
		{Type: TypeKey, Key: "--name", Desc: "Имя контейнера"},
		{Type: TypeKey, Key: "-p", Desc: "Проброс портов"},
		{Type: TypeKey, Key: "-P", Desc: "Автопроброс EXPOSE-портов"},
		{Type: TypeKey, Key: "-v", Desc: "Подключить том"},
		{Type: TypeKey, Key: "--mount", Desc: "Новый синтаксис подключения"},
		{Type: TypeKey, Key: "-e", Desc: "Переменные окружения"},
		{Type: TypeKey, Key: "--rm", Desc: "Удалить после остановки"},
		{Type: TypeKey, Key: "--restart", Desc: "Политика перезапуска"},
		{Type: TypeKey, Key: "--network", Desc: "Подключить сеть"},
		{Type: TypeKey, Key: "--hostname", Desc: "Имя хоста"},
		{Type: TypeKey, Key: "-w", Desc: "Рабочая директория"},
		{Type: TypeKey, Key: "--cpus", Desc: "Лимит CPU"},
		{Type: TypeKey, Key: "-m", Desc: "Лимит памяти"},

		// --------------------------------------------------
		// Вход в контейнер
		// --------------------------------------------------
		{Type: TypeHeader, Value: "📦 Вход в контейнер"},

		{Type: TypeCmd, Value: "docker exec -it <container> bash", Desc: "Войти через Bash"},
		{Type: TypeCmd, Value: "docker exec -it <container> sh", Desc: "Войти через sh"},
		{Type: TypeCmd, Value: "docker exec -u root -it <container> bash", Desc: "Войти как root"},
		{Type: TypeCmd, Value: "docker cp <container>:/file .", Desc: "Копировать файл из контейнера"},
		{Type: TypeCmd, Value: "docker cp ./file <container>:/path", Desc: "Копировать файл в контейнер"},

		{Type: TypeHeader, Value: "🔑 Ключи docker exec"},
		{Type: TypeKey, Key: "-i", Desc: "Оставить STDIN открытым"},
		{Type: TypeKey, Key: "-t", Desc: "Выделить TTY"},
		{Type: TypeKey, Key: "-u", Desc: "Выполнить от имени пользователя"},
		{Type: TypeKey, Key: "-w", Desc: "Рабочая директория"},
		{Type: TypeKey, Key: "-e", Desc: "Передать переменную окружения"},

		// --------------------------------------------------
		// Образы
		// --------------------------------------------------
		{Type: TypeHeader, Value: "📦 Образы"},

		{Type: TypeCmd, Value: "docker images", Desc: "Список образов"},
		{Type: TypeCmd, Value: "docker pull ubuntu:latest", Desc: "Скачать образ"},
		{Type: TypeCmd, Value: "docker build -t my_image .", Desc: "Собрать образ"},
		{Type: TypeCmd, Value: "docker rmi <image>", Desc: "Удалить образ"},

		{Type: TypeHeader, Value: "🔑 Ключи docker pull"},
		{Type: TypeKey, Key: "-a", Desc: "Скачать все теги"},
		{Type: TypeKey, Key: "--platform", Desc: "Указать платформу"},

		{Type: TypeHeader, Value: "🔑 Ключи docker build"},
		{Type: TypeKey, Key: "-t", Desc: "Имя и тег образа"},
		{Type: TypeKey, Key: "-f", Desc: "Указать Dockerfile"},
		{Type: TypeKey, Key: "--no-cache", Desc: "Не использовать кэш"},
		{Type: TypeKey, Key: "--build-arg", Desc: "Передать ARG"},

		// --------------------------------------------------
		// Docker Volumes
		// --------------------------------------------------
		{Type: TypeHeader, Value: "📦 Docker Volumes"},

		{Type: TypeCmd, Value: "docker volume ls", Desc: "Список томов"},
		{Type: TypeCmd, Value: "docker volume create my_volume", Desc: "Создать том"},
		{Type: TypeCmd, Value: "docker volume inspect my_volume", Desc: "Информация о томе"},
		{Type: TypeCmd, Value: "docker volume rm my_volume", Desc: "Удалить том"},
		{Type: TypeCmd, Value: "docker run -v my_volume:/data ubuntu:latest", Desc: "Подключить Docker Volume"},
		{Type: TypeCmd, Value: "docker run -v $(pwd):/app ubuntu:latest", Desc: "Подключить каталог хоста"},
		{Type: TypeCmd, Value: "docker run -it --mount type=volume,source=my_volume,target=/data ubuntu:latest", Desc: "Volume через --mount"},
		{Type: TypeCmd, Value: "docker run -it --mount type=bind,source=$(pwd),target=/app ubuntu:latest", Desc: "Bind Mount через --mount"},

		{Type: TypeHeader, Value: "🔑 Ключи Docker Volume"},
		{Type: TypeKey, Key: "-v", Desc: "Короткий синтаксис"},
		{Type: TypeKey, Key: "--mount", Desc: "Полный синтаксис"},
		{Type: TypeKey, Key: "type=volume", Desc: "Использовать Docker Volume"},
		{Type: TypeKey, Key: "type=bind", Desc: "Монтировать каталог хоста"},
		{Type: TypeKey, Key: "source=<name>", Desc: "Источник (том или каталог)"},
		{Type: TypeKey, Key: "target=<path>", Desc: "Путь внутри контейнера"},
		{Type: TypeKey, Key: "readonly", Desc: "Только чтение"},

		// --------------------------------------------------
		// Очистка
		// --------------------------------------------------
		{Type: TypeHeader, Value: "🧹 Очистка"},

		{Type: TypeCmd, Value: "docker container prune", Desc: "Удалить остановленные контейнеры"},
		{Type: TypeCmd, Value: "docker image prune", Desc: "Удалить неиспользуемые образы"},
		{Type: TypeCmd, Value: "docker volume prune", Desc: "Удалить неиспользуемые тома"},
		{Type: TypeCmd, Value: "docker system prune -a", Desc: "Полная очистка Docker"},

		{Type: TypeHeader, Value: "🔑 Ключи prune"},
		{Type: TypeKey, Key: "-a", Desc: "Удалить все неиспользуемые объекты"},
		{Type: TypeKey, Key: "-f", Desc: "Без подтверждения"},
		{Type: TypeKey, Key: "--volumes", Desc: "Удалить неиспользуемые тома"},
	},
}

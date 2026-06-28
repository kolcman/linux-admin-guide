package docker

var DockerfileSection = Section{
	Title: "📄 DOCKERFILE",
	Items: []Item{
		{Type: TypeTip, Value: "Dockerfile — текстовый рецепт. Из него собирают image (шаблон), из image запускают container (инстанс)."},
		{Type: TypeTip, Value: "Файл называется Dockerfile (без расширения), лежит в папке проекта вместе с кодом."},

		{Type: TypeHeader, Value: "🧠 Главное — когда что выполняется"},
		{Type: TypeKey, Key: "RUN", Desc: "при docker build — создание image (установка пакетов, сборка)"},
		{Type: TypeKey, Key: "ENTRYPOINT", Desc: "при docker run — главная команда при старте контейнера (только одна!)"},
		{Type: TypeKey, Key: "CMD", Desc: "при docker run — аргументы к ENTRYPOINT или команда по умолчанию (только одна!)"},
		{Type: TypeTip, Value: "docker build  →  RUN, RUN, RUN…     (собираем образ)\ndocker run    →  ENTRYPOINT + CMD   (запускаем контейнер)"},
		{Type: TypeWarn, Value: "Несколько ENTRYPOINT или CMD в одном файле — не работает. Остаётся только последний."},

		{Type: TypeHeader, Value: "📄 Инструкции при сборке (docker build)"},
		{Type: TypeKey, Key: "FROM", Desc: "базовый образ — с чего начинаем (alpine, ubuntu, node:20)"},
		{Type: TypeKey, Key: "RUN", Desc: "команда при сборке — можно несколько раз"},
		{Type: TypeKey, Key: "COPY", Desc: "скопировать файлы с сервера в образ (код, конфиги)"},
		{Type: TypeKey, Key: "ADD", Desc: "как COPY + умеет распаковывать tar (чаще используй COPY)"},
		{Type: TypeKey, Key: "WORKDIR", Desc: "рабочая папка внутри образа (/app)"},
		{Type: TypeKey, Key: "ENV", Desc: "переменные окружения внутри образа"},
		{Type: TypeKey, Key: "ARG", Desc: "переменные только на этапе сборки (не в runtime)"},

		{Type: TypeHeader, Value: "🔗 Как связаны ENTRYPOINT и CMD"},
		{Type: TypeTip, Value: "Представь: ENTRYPOINT — это программа, CMD — её аргументы по умолчанию."},
		{Type: TypeTip, Value: "ENTRYPOINT [\"ping\"]  — всегда запускается ping\nCMD [\"8.8.8.8\"]       — по умолчанию пингуем Google DNS"},
		{Type: TypeTip, Value: "docker run myimage\n→ ничего не передали, берётся CMD\n→ выполнится: ping 8.8.8.8"},
		{Type: TypeTip, Value: "docker run myimage 1.1.1.1\n→ свой аргумент заменяет CMD\n→ выполнится: ping 1.1.1.1"},
		{Type: TypeTip, Value: "Аргументы docker run НЕ попадают в CMD.\nОни заменяют CMD и уходят в ENTRYPOINT."},
		{Type: TypeKey, Key: "ENTRYPOINT", Desc: "что запускать — меняется редко (--entrypoint)"},
		{Type: TypeKey, Key: "CMD", Desc: "с чем запускать по умолчанию — легко переопределить при docker run"},

		{Type: TypeHeader, Value: "⚙️ Exec form и Shell form"},
		{Type: TypeTip, Value: "Контейнер живёт, пока работает процесс с PID 1.\nОстановился PID 1 — контейнер завершился."},
		{Type: TypeKey, Key: "exec form", Desc: "CMD [\"nginx\", \"-g\", \"daemon off;\"] — программа сразу PID 1"},
		{Type: TypeKey, Key: "shell form", Desc: "CMD nginx -g 'daemon off;' — обёртка /bin/sh -c, PID 1 у sh"},
		{Type: TypeTip, Value: "exec form (рекомендуется):\nENTRYPOINT [\"nginx\"]\nCMD [\"-g\", \"daemon off;\"]"},
		{Type: TypeTip, Value: "shell form (не для production):\nENTRYPOINT nginx -g 'daemon off;'"},
		{Type: TypeWarn, Value: "В shell form PID 1 — это sh/bash, а не твоё приложение."},
		{Type: TypeTip, Value: "docker stop шлёт SIGTERM процессу PID 1.\nВ shell form сигнал получит sh — приложение может не остановиться корректно."},
		{Type: TypeTip, Value: "Рекомендация: ENTRYPOINT и CMD пиши в exec form — [\"программа\", \"аргумент\"]."},

		{Type: TypeHeader, Value: "🚀 Инструкции при запуске (docker run)"},
		{Type: TypeKey, Key: "EXPOSE", Desc: "какой порт слушает приложение (подсказка, не проброс)"},

		{Type: TypeHeader, Value: "📝 Минимальный пример"},
		{Type: TypeTip, Value: "FROM alpine:3.19\nRUN apk add --no-cache curl\nWORKDIR /app\nCOPY . .\nEXPOSE 8080\nCMD [\"./start.sh\"]"},
		{Type: TypeTip, Value: "CMD в exec form [\"./start.sh\"] — start.sh станет PID 1, не sh."},

		{Type: TypeHeader, Value: "🔨 Собрать образ"},
		{Type: TypeKey, Key: "-t", Desc: "имя и тег образа: myapp:1.0"},
		{Type: TypeKey, Key: "-f", Desc: "другой файл вместо Dockerfile"},
		{Type: TypeKey, Key: ".", Desc: "контекст — папка, откуда COPY берёт файлы"},
		{Type: TypeKey, Key: "--no-cache", Desc: "собрать заново, без кеша слоёв"},
		{Type: TypeCmd, Value: "docker build -t myapp:1.0 .", Desc: "собрать image — здесь выполняются все RUN"},
		{Type: TypeCmd, Value: "docker build -t myapp:latest -f Dockerfile.prod .", Desc: "сборка из другого файла"},
		{Type: TypeCmd, Value: "docker build --no-cache -t myapp:1.0 .", Desc: "чистая пересборка с нуля"},

		{Type: TypeHeader, Value: "▶️ Запустить контейнер из образа"},
		{Type: TypeCmd, Value: "docker run -d -p 8080:80 --name myapp myapp:1.0", Desc: "создать инстанс и запустить CMD/ENTRYPOINT"},
		{Type: TypeCmd, Value: "docker run -it myapp:1.0 sh", Desc: "sh заменит CMD — откроется shell вместо start.sh"},

		{Type: TypeHeader, Value: "🔍 Проверка и отладка"},
		{Type: TypeKey, Key: "history", Desc: "показать слои образа и команды RUN из Dockerfile"},
		{Type: TypeCmd, Value: "docker history myapp:1.0", Desc: "какие RUN/COPY создали слои"},
		{Type: TypeCmd, Value: "docker inspect myapp:1.0", Desc: "детали образа: ENV, CMD, ENTRYPOINT"},
		{Type: TypeCmd, Value: "docker run --rm -it myapp:1.0 sh", Desc: "войти в контейнер без сохранения (--rm)"},

		{Type: TypeHeader, Value: "🏗️ Multi-stage (два FROM)"},
		{Type: TypeTip, Value: "Сначала собираем в «толстом» образе, потом копируем только результат в лёгкий."},
		{Type: TypeKey, Key: "AS builder", Desc: "имя стадии: FROM node:20 AS builder"},
		{Type: TypeKey, Key: "--target", Desc: "собрать только до указанной стадии"},
		{Type: TypeCmd, Value: "docker build -t myapp:1.0 --target production .", Desc: "сборка до стадии production"},

		{Type: TypeHeader, Value: "💡 Советы новичку"},
		{Type: TypeTip, Value: "Нужно несколько действий при старте? Сделай start.sh и ENTRYPOINT [\"./start.sh\"] (exec form)"},
		{Type: TypeTip, Value: "Один RUN = один слой. Лучше: RUN apt update && apt install -y curl — одной строкой."},
		{Type: TypeTip, Value: "Добавь .dockerignore — как .gitignore, чтобы не копировать лишнее в образ."},
		{Type: TypeWarn, Value: "Не храни пароли в Dockerfile — используй ENV при запуске или secrets."},
	},
}

package data

var ProcessesSection = Section{
	Title: "⚙️ ПРОЦЕССЫ",
	Items: []Item{
		{Type: TypeTip, Value: "Запущенная программа — процесс. Один бинарник можно запустить много раз — будет много процессов с разными PID.\nДанные о процессе также лежат в /proc/PID/."},
		{Type: TypeTip, Value: "PID — ID процесса. PPID — ID родителя (кто породил).\nСостояния: готов → выполняется → ждёт I/O → завершён (zombie, пока родитель не «заберёт» статус)."},

		{Type: TypeHeader, Value: "👀 Смотреть процессы"},
		{Type: TypeCmd, Value: "ps aux", Desc: "список процессов: CPU, RAM, команда"},
		{Type: TypeCmd, Value: "ps -ef", Desc: "список с PID и PPID"},
		{Type: TypeCmd, Value: "pgrep -a nginx", Desc: "найти PID по имени (+ команда)"},
		{Type: TypeCmd, Value: "pstree -p", Desc: "дерево процессов с PID"},
		{Type: TypeCmd, Value: "top", Desc: "интерактивно: кто грузит систему (q — выход)"},
		{Type: TypeCmd, Value: "htop", Desc: "удобнее top, если установлен"},

		{Type: TypeHeader, Value: "🛑 Сигналы и kill"},
		{Type: TypeTip, Value: "Сигнал — способ ядра сообщить процессу о событии.\nСначала мягко (SIGTERM), если не помогло и точно свой процесс — SIGKILL."},
		{Type: TypeKey, Key: "SIGTERM (15)", Desc: "попросить завершиться (по умолчанию у kill)"},
		{Type: TypeKey, Key: "SIGKILL (9)", Desc: "убить принудительно — процесс не может игнорировать"},
		{Type: TypeKey, Key: "SIGHUP (1)", Desc: "часто: перечитать конфиг (nginx, sshd)"},
		{Type: TypeCmd, Value: "kill PID", Desc: "SIGTERM — попросить завершиться"},
		{Type: TypeCmd, Value: "kill -9 PID", Desc: "SIGKILL — принудительно (последний аргумент)"},
		{Type: TypeCmd, Value: "killall имя", Desc: "сигнал всем процессам с таким именем"},
		{Type: TypeCmd, Value: "pkill -f \"python app.py\"", Desc: "убить по фрагменту командной строки"},
		{Type: TypeWarn, Value: "kill -9 — только если процесс точно завис и свой. Данные могут не сохраниться."},

		{Type: TypeHeader, Value: "🎚️ Приоритет (nice)"},
		{Type: TypeTip, Value: "nice: чем больше число — тем ниже приоритет (меньше жрёт CPU относительно других).\nДиапазон обычно −20…19. Отрицательный nice — только root."},
		{Type: TypeCmd, Value: "nice -n 10 command", Desc: "запустить с пониженным приоритетом"},
		{Type: TypeCmd, Value: "renice -n 5 -p PID", Desc: "изменить nice у уже запущенного"},

		{Type: TypeHeader, Value: "💡 Советы"},
		{Type: TypeTip, Value: "Алгоритм: ps/pgrep → понять, что за процесс → kill (15) → если не помогло — kill -9."},
		{Type: TypeCmd, Value: "ls /proc/PID/", Desc: "cwd, fd, cmdline, status — всё про процесс"},
		{Type: TypeCmd, Value: "cat /proc/PID/cmdline | tr '\\0' ' '", Desc: "полная командная строка процесса"},
	},
}

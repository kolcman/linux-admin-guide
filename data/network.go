package data

var NetworkSection = Section{
	Title: "🌐 СЕТЬ",
	Items: []Item{
		{Type: TypeTip, Value: "На хосте смотришь интерфейсы (ip), кто слушает порты (ss), есть ли связь (ping).\nСокет — «разъём» для обмена данными: TCP/UDP по сети или Unix-сокет локально."},

		{Type: TypeHeader, Value: "📡 Интерфейсы и маршруты"},
		{Type: TypeCmd, Value: "ip -br -c a", Desc: "IP-адреса интерфейсов (кратко, цвет)"},
		{Type: TypeCmd, Value: "ip a", Desc: "подробно: адреса, MAC, состояние"},
		{Type: TypeCmd, Value: "ip link", Desc: "состояние линка: UP/DOWN, MTU"},
		{Type: TypeCmd, Value: "ip r", Desc: "таблица маршрутов (default via …)"},
		{Type: TypeCmd, Value: "ip neigh", Desc: "ARP-таблица (соседи в L2)"},

		{Type: TypeHeader, Value: "🔌 Порты и сокеты (ss)"},
		{Type: TypeKey, Key: "-t", Desc: "TCP"},
		{Type: TypeKey, Key: "-u", Desc: "UDP"},
		{Type: TypeKey, Key: "-l", Desc: "только слушающие"},
		{Type: TypeKey, Key: "-p", Desc: "процесс, который держит сокет"},
		{Type: TypeKey, Key: "-n", Desc: "не резолвить имена (быстрее)"},
		{Type: TypeKey, Key: "-x", Desc: "Unix-сокеты (локальные)"},
		{Type: TypeKey, Key: "-s", Desc: "краткая сводка по сокетам"},
		{Type: TypeCmd, Value: "ss -tulpn", Desc: "слушающие TCP/UDP + процессы"},
		{Type: TypeCmd, Value: "ss -x", Desc: "Unix-сокеты (локальные IPC)"},
		{Type: TypeCmd, Value: "ss -s", Desc: "сводка: сколько сокетов каких типов"},
		{Type: TypeTip, Value: "Stream (TCP) — надёжный поток. Datagram (UDP) — пакеты.\nUnix-сокет — только на этой машине (часто /run/*.sock)."},

		{Type: TypeHeader, Value: "📶 Проверка связи"},
		{Type: TypeCmd, Value: "ping -c 4 8.8.8.8", Desc: "есть ли выход в интернет (ICMP)"},
		{Type: TypeCmd, Value: "ping -c 4 google.com", Desc: "связь + резолвится ли DNS"},
		{Type: TypeCmd, Value: "traceroute google.com", Desc: "по каким хопам идёт трафик"},
		{Type: TypeCmd, Value: "curl -I https://example.com", Desc: "HTTP-заголовки ответа (жив ли сервис)"},
		{Type: TypeCmd, Value: "curl -v telnet://host:22", Desc: "проверка TCP-порта без telnet"},

		{Type: TypeHeader, Value: "🔤 DNS"},
		{Type: TypeCmd, Value: "resolvectl status", Desc: "какие DNS сейчас у systemd-resolved"},
		{Type: TypeCmd, Value: "dig example.com", Desc: "запрос A-записи (пакет dnsutils)"},
		{Type: TypeCmd, Value: "getent hosts example.com", Desc: "как система резолвит имя"},
		{Type: TypeCmd, Value: "cat /etc/resolv.conf", Desc: "указанные nameserver (может быть stub)"},

		{Type: TypeHeader, Value: "💡 Советы"},
		{Type: TypeTip, Value: "Нет сети: ip link (линк UP?) → ip a (есть IP?) → ip r (есть default?) → ping 8.8.8.8 → ping по имени (DNS)."},
		{Type: TypeTip, Value: "Порт занят: ss -tulpn | grep :80 — увидишь PID и процесс."},
	},
}

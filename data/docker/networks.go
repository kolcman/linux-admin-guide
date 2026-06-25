package docker

var NetworksSection = Section{
	Title: "🌐 СЕТИ И ПОРТЫ",
	Items: []Item{
		{Type: TypeTip, Value: "Сеть связывает контейнеры между собой. Без сети и портов контейнер недоступен снаружи."},
		{Type: TypeTip, Value: "Контейнеры в одной сети обращаются друг к другу по имени: curl http://web:80"},

		{Type: TypeHeader, Value: "🔌 Типы сетей"},
		{Type: TypeKey, Key: "bridge", Desc: "по умолчанию — изолированная сеть с NAT"},
		{Type: TypeKey, Key: "host", Desc: "контейнер использует сеть хоста напрямую"},
		{Type: TypeKey, Key: "none", Desc: "без сети — полная изоляция"},
		{Type: TypeKey, Key: "overlay", Desc: "сеть между несколькими серверами (Swarm)"},

		{Type: TypeHeader, Value: "🌐 Управление сетями"},
		{Type: TypeKey, Key: "--driver", Desc: "тип сети при создании (bridge, overlay…)"},
		{Type: TypeCmd, Value: "docker network ls", Desc: "список сетей"},
		{Type: TypeCmd, Value: "docker network create --driver bridge my_net", Desc: "создать bridge-сеть"},
		{Type: TypeCmd, Value: "docker network inspect my_net", Desc: "какие контейнеры подключены"},
		{Type: TypeCmd, Value: "docker network connect my_net <id>", Desc: "подключить контейнер к сети"},
		{Type: TypeCmd, Value: "docker network disconnect my_net <id>", Desc: "отключить от сети"},

		{Type: TypeHeader, Value: "🔌 Проброс портов"},
		{Type: TypeKey, Key: "-p", Desc: "host:container — проброс порта на сервер"},
		{Type: TypeWarn, Value: "Без -p сервис внутри контейнера недоступен снаружи"},
		{Type: TypeCmd, Value: "docker run -d -p 8080:80 --name web nginx", Desc: "nginx в фоне, порт 8080 → 80"},
		{Type: TypeCmd, Value: "docker run -p 127.0.0.1:8080:80 nginx", Desc: "только localhost, не из интернета"},

		{Type: TypeHeader, Value: "🚀 Запуск в сети"},
		{Type: TypeKey, Key: "--network", Desc: "указать сеть при запуске контейнера"},
		{Type: TypeCmd, Value: "docker run -d --network my_net --name api myapp", Desc: "запуск в пользовательской сети"},
		{Type: TypeCmd, Value: "docker run --network host nginx", Desc: "сеть хоста — без изоляции портов"},
	},
}

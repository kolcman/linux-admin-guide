package docker

var VolumesSection = Section{
	Title: "📦 VOLUMES И MOUNT",
	Items: []Item{
		{Type: TypeTip, Value: "Volume хранит данные вне контейнера. Пересоздали контейнер — данные в volume останутся."},
		{Type: TypeTip, Value: "Bind mount — привязка папки с сервера. Tmpfs — данные в RAM (пропадут при остановке)."},

		{Type: TypeHeader, Value: "📦 Docker Volume"},
		{Type: TypeKey, Key: "-v", Desc: "volume:путь — подключить именованный том"},
		{Type: TypeCmd, Value: "docker volume ls", Desc: "список томов на сервере"},
		{Type: TypeCmd, Value: "docker volume create my_data", Desc: "создать том для БД или файлов"},
		{Type: TypeCmd, Value: "docker volume inspect my_data", Desc: "где физически лежат данные"},
		{Type: TypeCmd, Value: "docker run -d -v my_data:/var/lib/mysql --name db mysql", Desc: "mysql с постоянным хранилищем"},

		{Type: TypeHeader, Value: "🔗 Bind mount — папка с сервера"},
		{Type: TypeKey, Key: ":ro", Desc: "read-only — только чтение, контейнер не пишет"},
		{Type: TypeKey, Key: ":rw", Desc: "read-write — чтение и запись (по умолчанию)"},
		{Type: TypeKey, Key: "--mount", Desc: "современный синтаксис вместо -v"},
		{Type: TypeCmd, Value: "docker run -d -v /opt/app:/app nginx", Desc: "папка сервера → /app в контейнере"},
		{Type: TypeCmd, Value: "docker run -v /opt/app:/app:ro nginx", Desc: "только чтение — для конфигов"},
		{Type: TypeCmd, Value: "docker run --mount type=bind,source=/opt/app,target=/app nginx", Desc: "явный bind mount"},

		{Type: TypeHeader, Value: "💾 Tmpfs — данные в RAM"},
		{Type: TypeKey, Key: "--tmpfs", Desc: "временное хранилище в оперативной памяти"},
		{Type: TypeCmd, Value: "docker run --tmpfs /tmp:rw,size=128m nginx", Desc: "/tmp в RAM, 128 МБ, не сохраняется"},
	},
}

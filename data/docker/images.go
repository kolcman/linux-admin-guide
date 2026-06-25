package docker

var ImagesSection = Section{
	Title: "📦 ОБРАЗЫ",
	Items: []Item{
		{Type: TypeTip, Value: "Образ (image) — шаблон приложения. Из одного образа можно запустить много контейнеров."},
		{Type: TypeTip, Value: "Образы хранятся на диске сервера. Скачиваются из Docker Hub или собираются из Dockerfile."},

		{Type: TypeHeader, Value: "📋 Список образов"},
		{Type: TypeKey, Key: "-a", Desc: "все образы, включая промежуточные слои"},
		{Type: TypeKey, Key: "--filter", Desc: "фильтр, например dangling=true"},
		{Type: TypeCmd, Value: "docker images", Desc: "образы на этом сервере"},
		{Type: TypeCmd, Value: "docker image ls -a", Desc: "полный список включая слои сборки"},
		{Type: TypeCmd, Value: "docker image ls --filter dangling=true", Desc: "только «висячие» образы без тега"},

		{Type: TypeHeader, Value: "⬇️ Скачивание"},
		{Type: TypeKey, Key: "тег", Desc: "версия образа после двоеточия, например nginx:1.25"},
		{Type: TypeCmd, Value: "docker pull nginx", Desc: "скачать последнюю версию nginx"},
		{Type: TypeCmd, Value: "docker pull nginx:1.25", Desc: "скачать конкретную версию"},

		{Type: TypeHeader, Value: "🔨 Сборка своего образа"},
		{Type: TypeKey, Key: "-t", Desc: "tag — имя и версия образа (myapp:v1)"},
		{Type: TypeKey, Key: "-f", Desc: "указать другой Dockerfile"},
		{Type: TypeKey, Key: ".", Desc: "контекст сборки — папка с Dockerfile"},
		{Type: TypeCmd, Value: "docker build -t myapp:latest .", Desc: "собрать образ из Dockerfile в текущей папке"},
		{Type: TypeCmd, Value: "docker build -t myapp:v1 -f Dockerfile.prod .", Desc: "сборка из другого файла"},

		{Type: TypeHeader, Value: "🗑️ Удаление"},
		{Type: TypeKey, Key: "-a", Desc: "удалить все неиспользуемые образы"},
		{Type: TypeKey, Key: "-f", Desc: "без подтверждения yes/no"},
		{Type: TypeCmd, Value: "docker rmi <image>", Desc: "удалить конкретный образ"},
		{Type: TypeCmd, Value: "docker image prune -f", Desc: "убрать «висячие» образы без тега"},
		{Type: TypeCmd, Value: "docker image prune -a -f", Desc: "удалить все образы без контейнеров"},
	},
}

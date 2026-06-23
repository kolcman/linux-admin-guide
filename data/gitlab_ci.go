package data

var GitLabSection = Section{
	Title: "🚀 GITLAB CI/CD — ПЛАЙПАЙНЫ",
	Items: []Item{
		{Type: TypeHeader, Value: "🏗️ Директивы .gitlab-ci.yml:"},
		{Type: TypeYaml, Key: "stages:", Desc: "Этапы (build -> test -> deploy)"},
		{Type: TypeYaml, Key: "image:", Desc: "Docker-образ для выполнения job"},
		{Type: TypeYaml, Key: "script:", Desc: "Команды shell"},
		{Type: TypeYaml, Key: "artifacts:", Desc: "Файлы для сохранения между этапами"},
		{Type: TypeYaml, Key: "cache:", Desc: "Кэш зависимостей (ускоряет сборку)"},

		{Type: TypeHeader, Value: "📋 Переменные окружения:"},
		{Type: TypeVar, Key: "CI_COMMIT_REF_NAME", Desc: "Имя ветки (main, feature/x)"},
		{Type: TypeVar, Key: "CI_PROJECT_DIR", Desc: "Абсолютный путь к проекту на раннере"},
		{Type: TypeVar, Key: "CI_PIPELINE_ID", Desc: "ID текущего пайплайна"},

		{Type: TypeHeader, Value: "🛠️ Runner:"},
		{Type: TypeCmd, Value: "sudo gitlab-runner register", Desc: "Регистрация нового исполнителя"},
		{Type: TypeCmd, Value: "sudo gitlab-runner verify", Desc: "Проверка связи с GitLab"},
	},
}

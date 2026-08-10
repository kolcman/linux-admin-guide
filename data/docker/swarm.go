package docker

var SwarmSection = Section{
	Title: "🐝 DOCKER SWARM",
	Items: []Item{
		{Type: TypeTip, Value: "Swarm — встроенная оркестрация Docker: несколько серверов, сервисы с репликами, overlay-сеть."},
		{Type: TypeTip, Value: "Для одного хоста чаще хватает Compose. Swarm — когда нужны несколько нод без Kubernetes."},

		{Type: TypeHeader, Value: "🏗️ Кластер"},
		{Type: TypeCmd, Value: "docker swarm init", Desc: "сделать текущий хост manager’ом"},
		{Type: TypeCmd, Value: "docker swarm join-token worker", Desc: "токен для worker-ноды"},
		{Type: TypeCmd, Value: "docker swarm join-token manager", Desc: "токен для дополнительного manager"},
		{Type: TypeCmd, Value: "docker swarm join --token <token> <manager-ip>:2377", Desc: "подключить ноду к кластеру"},
		{Type: TypeCmd, Value: "docker node ls", Desc: "список нод"},
		{Type: TypeCmd, Value: "docker swarm leave --force", Desc: "выйти из swarm (осторожно на manager)"},

		{Type: TypeHeader, Value: "🚀 Сервисы"},
		{Type: TypeCmd, Value: "docker service create --name web -p 80:80 --replicas 3 nginx", Desc: "сервис с 3 репликами"},
		{Type: TypeCmd, Value: "docker service ls", Desc: "список сервисов"},
		{Type: TypeCmd, Value: "docker service ps web", Desc: "задачи (tasks) сервиса по нодам"},
		{Type: TypeCmd, Value: "docker service logs -f web", Desc: "логи сервиса"},
		{Type: TypeCmd, Value: "docker service scale web=5", Desc: "изменить число реплик"},
		{Type: TypeCmd, Value: "docker service update --image nginx:1.27 web", Desc: "обновить образ (rolling update)"},
		{Type: TypeCmd, Value: "docker service rm web", Desc: "удалить сервис"},

		{Type: TypeHeader, Value: "📚 Stack"},
		{Type: TypeTip, Value: "Stack — деплой compose-файла в Swarm.\nВ файле обычно нужны deploy: (replicas, placement) — обычный Compose на одном хосте их игнорирует."},
		{Type: TypeCmd, Value: "docker stack deploy -c docker-compose.yml mystack", Desc: "задеплоить compose как stack"},
		{Type: TypeCmd, Value: "docker stack ls", Desc: "список stack’ов"},
		{Type: TypeCmd, Value: "docker stack services mystack", Desc: "сервисы внутри stack"},
		{Type: TypeCmd, Value: "docker stack rm mystack", Desc: "снять stack"},

		{Type: TypeHeader, Value: "💡 Советы"},
		{Type: TypeTip, Value: "Порт публикуется на всех нодах (routing mesh): зашёл на любой IP кластера — попал на реплику."},
		{Type: TypeWarn, Value: "swarm leave --force на последнем manager разрушит кластер. Снимай сервисы/stack заранее."},
	},
}

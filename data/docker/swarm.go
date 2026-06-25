package docker

var SwarmSection = Section{
	Title: "🐝 DOCKER SWARM",
	Items: []Item{
		{Type: TypeTip, Value: "Swarm — оркестрация контейнеров на нескольких серверах. Узлы объединяются в кластер."},
		{Type: TypeTip, Value: "Раздел в разработке. Скоро: init, join, services, stacks, nodes."},
		{Type: TypeHeader, Value: "🔜 Планируется"},
		{Type: TypeKey, Key: "swarm init", Desc: "создать кластер на текущем сервере"},
		{Type: TypeKey, Key: "swarm join", Desc: "подключить сервер к кластеру"},
		{Type: TypeKey, Key: "service", Desc: "управление сервисами в кластере"},
		{Type: TypeKey, Key: "stack deploy", Desc: "деплой compose-стека в Swarm"},
	},
}

package docker

func Submenu() []MenuItem {
	return []MenuItem{
		{Title: "🐳 Docker", Children: dockerBasicsSubmenu()},
		{Title: "🐳 Docker Compose", Section: &ComposeSection},
		{Title: "🐝 Docker Swarm", Section: &SwarmSection},
	}
}

func dockerBasicsSubmenu() []MenuItem {
	return []MenuItem{
		{Title: "🔧 Контейнеры", Section: &ContainersSection},
		{Title: "📦 Образы", Section: &ImagesSection},
		{Title: "🌐 Сети и порты", Section: &NetworksSection},
		{Title: "📦 Volumes", Section: &VolumesSection},
		{Title: "🧹 Очистка", Section: &CleanupSection},
	}
}

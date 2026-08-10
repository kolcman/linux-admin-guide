package data

func GetMenu() []MenuItem {
	return []MenuItem{
		{Title: "📁 ФАЙЛЫ И КАТАЛОГИ", Children: filesSubmenu()},
		{Title: BashSection.Title, Section: &BashSection},
		{Title: TextSection.Title, Section: &TextSection},
		{Title: RegexSection.Title, Section: &RegexSection},
		{Title: DisksSection.Title, Section: &DisksSection},
		{Title: ProcessesSection.Title, Section: &ProcessesSection},
		{Title: ServicesSection.Title, Section: &ServicesSection},
		{Title: BootSection.Title, Section: &BootSection},
		{Title: UsersSection.Title, Section: &UsersSection},
		{Title: NetworkSection.Title, Section: &NetworkSection},
		{Title: FirewallSection.Title, Section: &FirewallSection},
		{Title: SSHSection.Title, Section: &SSHSection},
		{Title: BackupSection.Title, Section: &BackupSection},
		{Title: PackagesSection.Title, Section: &PackagesSection},
		{Title: "🐳 DOCKER", Children: dockerSubmenu()},
		{Title: GitLabSection.Title, Section: &GitLabSection},
	}
}

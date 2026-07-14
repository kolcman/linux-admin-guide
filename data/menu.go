package data

func GetMenu() []MenuItem {
	return []MenuItem{
		{Title: FilesSection.Title, Section: &FilesSection},
		{Title: BashSection.Title, Section: &BashSection},
		{Title: TextSection.Title, Section: &TextSection},
		{Title: RegexSection.Title, Section: &RegexSection},
		{Title: DisksSection.Title, Section: &DisksSection},
		{Title: FirewallSection.Title, Section: &FirewallSection},
		{Title: UsersSection.Title, Section: &UsersSection},
		{Title: SSHSection.Title, Section: &SSHSection},
		{Title: NetworkSection.Title, Section: &NetworkSection},
		{Title: BackupSection.Title, Section: &BackupSection},
		{Title: ServicesSection.Title, Section: &ServicesSection},
		{Title: PackagesSection.Title, Section: &PackagesSection},
		{Title: "🐳 DOCKER", Children: dockerSubmenu()},
		{Title: GitLabSection.Title, Section: &GitLabSection},
	}
}

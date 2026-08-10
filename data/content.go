package data

// GetAllSections — плоский список (legacy). Меню: data.GetMenu().
func GetAllSections() map[int]Section {
	return map[int]Section{
		1:  DisksSection,
		2:  FirewallSection,
		3:  UsersCreateSection,
		4:  UsersGroupsSection,
		5:  UsersSudoSection,
		6:  UsersAclSection,
		7:  SSHSection,
		8:  NetworkSection,
		9:  BackupSection,
		10: ServicesSection,
		11: PackagesSection,
		12: GitLabSection,
		13: ProcessesSection,
		14: BootSection,
	}
}

package data

// GetAllSections — плоский список (legacy). Меню: data.GetMenu().
func GetAllSections() map[int]Section {
	return map[int]Section{
		1:  DisksSection,
		2:  FirewallSection,
		3:  UsersSection,
		4:  SSHSection,
		5:  NetworkSection,
		6:  BackupSection,
		7:  ServicesSection,
		8:  PackagesSection,
		9:  GitLabSection,
	}
}

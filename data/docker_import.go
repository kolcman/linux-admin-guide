package data

import "linux-guide/data/docker"

func dockerSubmenu() []MenuItem {
	return importDockerMenu(docker.Submenu())
}

func importDockerMenu(items []docker.MenuItem) []MenuItem {
	out := make([]MenuItem, len(items))
	for i, item := range items {
		mi := MenuItem{Title: item.Title}
		if item.Section != nil {
			mi.Section = importDockerSection(item.Section)
		}
		if len(item.Children) > 0 {
			mi.Children = importDockerMenu(item.Children)
		}
		out[i] = mi
	}
	return out
}

func importDockerSection(s *docker.Section) *Section {
	items := make([]Item, len(s.Items))
	for i, it := range s.Items {
		items[i] = Item{
			Type:  it.Type,
			Key:   it.Key,
			Value: it.Value,
			Desc:  it.Desc,
		}
	}
	sec := Section{Title: s.Title, Items: items}
	return &sec
}

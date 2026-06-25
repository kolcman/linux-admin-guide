package docker

const (
	TypeCmd    = "cmd"
	TypeKey    = "key"
	TypeTip    = "tip"
	TypeWarn   = "warn"
	TypeVar    = "var"
	TypeYaml   = "yaml"
	TypeHeader = "header"
)

type Item struct {
	Type  string
	Key   string
	Value string
	Desc  string
}

type Section struct {
	Title string
	Items []Item
}

type MenuItem struct {
	Title    string
	Section  *Section
	Children []MenuItem
}

func (m MenuItem) IsGroup() bool {
	return len(m.Children) > 0
}

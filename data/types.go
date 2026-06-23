package data

const (
	TypeCmd    = "cmd"
	TypeKey    = "key"
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

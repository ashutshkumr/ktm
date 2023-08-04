package api

type Topology struct {
	Name  string `yaml:"name"`
	Nodes []Node `yaml:"nodes"`
	Links []Link `yaml:"links"`
}

type Node struct {
	Name  string   `yaml:"name"`
	Image string   `yaml:"image"`
	Cmd   []string `yaml:"cmd"`
}

type Link struct {
	Name   string    `yaml:"name"`
	First  Interface `yaml:"first"`
	Second Interface `yaml:"second"`
}

type Interface struct {
	Name string `yaml:"name"`
	Node string `yaml:"node"`
}

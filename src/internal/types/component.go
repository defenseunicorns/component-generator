package types

type ComponentsConfig struct {
	Name       string      `json:"name" yaml:"name"`
	Components []Component `json:"components" yaml:"components"`
}

type Component struct {
	locals  []Local  `json:"local" yaml:"local"`
	remotes []remote `json:"remote" yaml:"remote"`
}

type Local struct {
	Name string `json:"name" yaml:"name"`
}

type remote struct {
	Git  string `json:"git" yaml:"git"`
	Path string `json:"path" yaml:"path"`
}

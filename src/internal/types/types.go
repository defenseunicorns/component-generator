package types

type JsonMap map[string]interface{}

type OscalComponentDocumentGeneric struct {
	ComponentDefinition JsonMap `json:"component-definition" yaml:"component-definition"`
}

type ComponentsConfig struct {
	Name          string    `json:"name" yaml:"name"`
	Metadata      JsonMap   `json:"metadata" yaml:"metadata"`
	Components    Component `json:"components" yaml:"components"`
	BaseDirectory string    `json:"base-directory" yaml:"base-directory"`
}

type Component struct {
	Locals  []Local  `json:"local" yaml:"local"`
	Remotes []Remote `json:"remote" yaml:"remote"`
}

type Local struct {
	Name string `json:"name" yaml:"name"`
}

type Remote struct {
	Git  string `json:"git" yaml:"git"`
	Path string `json:"path" yaml:"path"`
}

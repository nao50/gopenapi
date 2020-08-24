package gopenapi

type Header struct {
	ExtensionProps
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Deprecated  bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Required    bool                   `json:"required,omitempty" yaml:"required,omitempty"`
	Schema      *SchemaRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example     interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples    map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Content     Content                `json:"content,omitempty" yaml:"content,omitempty"`
}

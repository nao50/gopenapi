package gopenapi

type Parameters []*ParameterRef

// Parameter Object.
type Parameter struct {
	ExtensionProps
	Name            string                 `json:"name,omitempty" yaml:"name,omitempty"`
	In              string                 `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Style           string                 `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         *bool                  `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowEmptyValue bool                   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	AllowReserved   bool                   `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Deprecated      bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Required        bool                   `json:"required,omitempty" yaml:"required,omitempty"`
	Schema          *SchemaRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example         interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Content         Content                `json:"content,omitempty" yaml:"content,omitempty"`
}

package gopenapi

// MediaType Object.
type MediaType struct {
	ExtensionProps
	Schema   *SchemaRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]*Encoding   `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

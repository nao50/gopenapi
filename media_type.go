package gopenapi

// MediaType Object.
type MediaType struct {
	Schema   *Schema              `yaml:"schema"` // TODO: Schema Object | Reference Object
	Example  interface{}          `yaml:"example"`
	Examples map[string]*Example  `yaml:"examples"` // TODO: Map[ string, Example Object | Reference Object]
	Encoding map[string]*Encoding `yaml:"encoding"` // TODO:
}

package gopenapi

// Encoding Object.
type Encoding struct {
	ContentType   string             `yaml:"contentType"`
	Headers       map[string]*Header `yaml:"headers"` // TODO: Map[string, Header Object | Reference Object]
	Style         string             `yaml:"style"`
	Explode       bool               `yaml:"explode"`
	AllowReserved bool               `yaml:"allowReserved"`
}

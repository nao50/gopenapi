package gopenapi

// ServerVariable Object.
type ServerVariable struct {
	Enum        []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string        `json:"description,omitempty" yaml:"description,omitempty"`
}

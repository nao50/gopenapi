package gopenapi

// Encoding : https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#encoding-object
type Encoding struct {
	ExtensionProps
	ContentType   string                `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*HeaderRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string                `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       *bool                 `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                  `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

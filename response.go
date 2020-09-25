package gopenapi

type Responses map[string]*ResponseRef

// Response Object.
type Response struct {
	ExtensionProps
	Description *string               `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]*HeaderRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     Content               `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*LinkRef   `json:"links,omitempty" yaml:"links,omitempty"`
}

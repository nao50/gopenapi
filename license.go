package gopenapi

// License Object.
type License struct {
	ExtensionProps
	Name string `json:"name" yaml:"name"` // Required
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

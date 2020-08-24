package gopenapi

// Contact : https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#contact-object
type Contact struct {
	ExtensionProps
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

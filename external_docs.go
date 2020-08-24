package gopenapi

// ExternalDocumentation : https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#external-documentation-object
type ExternalDocs struct {
	ExtensionProps
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}

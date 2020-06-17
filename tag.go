package gopenapi

// Tag Object.
type Tag struct {
	Name         string                 `yaml:"name"`
	Description  string                 `yaml:"description"`
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs"`
}

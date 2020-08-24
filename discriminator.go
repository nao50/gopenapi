package gopenapi

// Discriminator : https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#discriminator-object
type Discriminator struct {
	ExtensionProps
	PropertyName string            `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

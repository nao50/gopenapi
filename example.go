package gopenapi

// Example : https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#example-object
type Example struct {
	ExtensionProps
	Summary       string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string      `json:"description,omitempty" yaml:"description,omitempty"`
	Value         interface{} `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue string      `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}

package gopenapi

// Example Object.
type Example struct {
	Summary       string      `yaml:"summary"`
	Description   string      `yaml:"description"`
	Value         interface{} `yaml:"value"`
	ExternalValue interface{} `yaml:"externalValue"`
}

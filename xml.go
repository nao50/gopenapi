package gopenapi

// XML Object.
type XML struct {
	Name      string `json:"name" yaml:"name"`
	Namespace string `json:"namespace" yaml:"namespace"`
	Prefix    string `json:"prefix" yaml:"prefix"`
	Attribute bool   `json:"attribute" yaml:"attribute"`
	Wrapped   bool   `json:"wrapped" yaml:"wrapped"`
}

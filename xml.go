package gopenapi

// XML Object.
type XML struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Prefix    string `yaml:"prefix"`
	Attribute bool   `yaml:"attribute"`
	Wrapped   bool   `yaml:"wrapped"`
}

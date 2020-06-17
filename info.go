package gopenapi

// Info Object.
type Info struct {
	Title          string   `yaml:"title"`
	Description    string   `yaml:"description"`
	TermsOfService string   `yaml:"termsOfService"`
	Contact        *Contact `yaml:"contact"`
	License        *License `yaml:"license"`
	Version        string   `yaml:"version"`
}

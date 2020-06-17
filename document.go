package gopenapi

// Document Object.
type Document struct {
	Version      string    `yaml:"openapi"`
	Info         *Info     `yaml:"info"`
	Servers      []*Server `yaml:"servers"`
	Paths        Paths     `yaml:"paths"`
	Components   *Components
	Security     []*SecurityRequirement
	Tags         []*Tag
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs"`
}

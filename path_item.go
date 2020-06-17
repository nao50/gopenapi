package gopenapi

// PathItem Object.
type PathItem struct {
	Ref         string       `yaml:"$ref"`
	Summary     string       `yaml:"summary"`
	Description string       `yaml:"description"`
	Get         *Operation   `yaml:"get"`
	Put         *Operation   `yaml:"put"`
	Post        *Operation   `yaml:"post"`
	Delete      *Operation   `yaml:"delete"`
	Options     *Operation   `yaml:"options"`
	Head        *Operation   `yaml:"head"`
	Patch       *Operation   `yaml:"patch"`
	Trace       *Operation   `yaml:"trace"`
	Servers     []*Server    `yaml:"servers"`
	Parameters  []*Parameter `yaml:"parameters"`
}

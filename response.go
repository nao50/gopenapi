package gopenapi

// Response Object.
type Response struct {
	Description string                `yaml:"description"`
	Headers     map[string]*Header    `yaml:"headers"` // TODO: Map[string, Header Object | Reference Object]
	Content     map[string]*MediaType `yaml:"content"`
	Links       map[string]*Link      `yaml:"links"` // TODO: Map[string, Link Object | Reference Object]
}

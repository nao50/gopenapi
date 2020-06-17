package gopenapi

// Server Object.
type Server struct {
	URL         string                     `yaml:"title"`
	Description string                     `yaml:"description"`
	Variables   map[string]*ServerVariable `yaml:"variables"`
}

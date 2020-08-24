package gopenapi

// Servers
type Servers []*Server

// Server : https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#fixed-fields-4
type Server struct {
	URL         string                     `json:"url,omitempty"`
	Description string                     `json:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty"`
}

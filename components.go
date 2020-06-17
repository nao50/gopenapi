package gopenapi

// Components Object.
type Components struct {
	Schemas         map[string]*Schema         `yaml:"schemas"`
	Responses       map[string]*Response       `yaml:"responses"`
	Parameters      map[string]*Parameter      `yaml:"parameters"`
	Examples        map[string]*Example        `yaml:"examples"`
	RequestBodies   map[string]*RequestBody    `yaml:"requestBodies"`
	Headers         map[string]*Header         `yaml:"headers"`
	SecuritySchemes map[string]*SecurityScheme `yaml:"securitySchemes"`
	Links           map[string]*Link           `yaml:"links"`
	Callbacks       map[string]*Callback       `yaml:"callbacks"`
}

package gopenapi

// Operation Object.
type Operation struct {
	Tags         []string               `yaml:"tags"`
	Summary      string                 `yaml:"summary"`
	Description  string                 `yaml:"description"`
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs"`
	OperationID  string                 `yaml:"operationId"`
	Parameters   []*Parameter           `yaml:"parameters"`  // TODO: [Parameter Object | Reference Object]
	RequestBody  *RequestBody           `yaml:"requestBody"` // TODO: Request Body Object | Reference Object
	Responses    Responses              `yaml:"responses"`
	Callbacks    map[string]*Callback   `yaml:"callbacks"`
	Deprecated   bool                   `yaml:"deprecated"`
	Security     []*SecurityRequirement `yaml:"security"`
	Servers      []*Server              `yaml:"servers"`
}

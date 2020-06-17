package gopenapi

// Link Object.
type Link struct {
	OperationRef string                 `yaml:"operationRef"`
	OperationID  string                 `yaml:"operationId"`
	Parameters   map[string]interface{} `yaml:"parameters"`  // TODO: Map[string, Any | {expression}]
	RequestBody  interface{}            `yaml:"requestBody"` // TODO: Any | {expression}
	Description  string                 `yaml:"description"`
	Server       *Server                `yaml:"server"`
}

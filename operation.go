package gopenapi

// Operation Object.
type Operation struct {
	ExtensionProps
	Tags         []string                `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                  `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID  string                  `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   Parameters              `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBodyRef         `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    Responses               `json:"responses" yaml:"responses"` // Required
	Callbacks    map[string]*CallbackRef `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                    `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     *SecurityRequirements   `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      *Servers                `json:"servers,omitempty" yaml:"servers,omitempty"`
	ExternalDocs *ExternalDocs           `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

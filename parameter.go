package gopenapi

// Parameter Object.
type Parameter struct {
	Name            string                `yaml:"name"`
	In              string                `yaml:"in"`
	Description     string                `yaml:"description"`
	Required        bool                  `yaml:"required"` // TODO: Determines whether this parameter is mandatory. If the parameter location is "path", this property is REQUIRED and its value MUST be true. Otherwise, the property MAY be included and its default value is false.
	Deprecated      string                `yaml:"deprecated"`
	AllowEmptyValue bool                  `yaml:"allowEmptyValue"`
	Style           string                `yaml:"style"` // TODO: Describes how the parameter value will be serialized depending on the type of the parameter value. Default values (based on value of in): for query - form; for path - simple; for header - simple; for cookie - form.
	Explode         bool                  `yaml:"explode"`
	AllowReserved   bool                  `yaml:"allowReserved"`
	Schema          *Schema               `yaml:"schema"` // TODO: Schema Object | Reference Object
	Example         interface{}           `yaml:"example"`
	Examples        map[string]*Example   `yaml:"examples"`
	Content         map[string]*MediaType `yaml:"content"`
	Ref             string                `yaml:"$ref"`
}

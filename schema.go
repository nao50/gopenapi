package gopenapi

// Schema Object.  定義ファイル探す
type Schema struct {
	Title                string                 `yaml:"title"`
	MultipleOf           int                    `yaml:"multipleOf"`
	Maximum              int                    `yaml:"maximum"`
	ExclusiveMaximum     bool                   `yaml:"exclusiveMaximum"`
	Minimum              int                    `yaml:"minimum"`
	ExclusiveMinimum     bool                   `yaml:"exclusiveMinimum"`
	MaxLength            int                    `yaml:"maxLength"`
	MinLength            int                    `yaml:"minLength"`
	Pattern              string                 `yaml:"pattern"`
	MaxItems             int                    `yaml:"maxItems"`
	MinItems             int                    `yaml:"minItems"`
	MaxProperties        int                    `yaml:"maxProperties"`
	MinProperties        int                    `yaml:"minProperties"`
	Required             []string               `yaml:"required"` // TODO: []string??
	Enum                 []string               `yaml:"enum"`
	Type                 string                 `yaml:"type"`
	AllOf                []*Schema              `yaml:"allOf"`
	OneOf                []*Schema              `yaml:"oneOf"`
	AnyOf                []*Schema              `yaml:"anyOf"`
	Not                  *Schema                `yaml:"not"`
	Items                *Schema                `yaml:"items"`
	Properties           map[string]*Schema     `yaml:"properties"`
	AdditionalProperties *Schema                `yaml:"additionalProperties"`
	Description          string                 `yaml:"description"`
	Format               string                 `yaml:"format"`
	Default              string                 `yaml:"default"`
	Nullable             bool                   `yaml:"nullable"`
	Discriminator        *Discriminator         `yaml:"discriminator"`
	ReadOnly             bool                   `yaml:"readOnly"`
	WriteOnly            bool                   `yaml:"writeOnly"`
	XML                  *XML                   `yaml:"xml"`
	ExternalDocs         *ExternalDocumentation `yaml:"externalDocs"`
	Example              interface{}            `yaml:"example"`
	Deprecated           bool                   `yaml:"deprecated"`
	Ref                  string                 `yaml:"$ref"`
	Extension            map[string]interface{}
}

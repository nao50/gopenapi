package gopenapi

import "regexp"

//
type Schema struct {
	ExtensionProps

	OneOf        []*SchemaRef  `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf        []*SchemaRef  `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	AllOf        []*SchemaRef  `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	Not          *SchemaRef    `json:"not,omitempty" yaml:"not,omitempty"`
	Type         string        `json:"type,omitempty" yaml:"type,omitempty"`
	Title        string        `json:"title,omitempty" yaml:"title,omitempty"`
	Format       string        `json:"format,omitempty" yaml:"format,omitempty"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	Enum         []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default      interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
	Example      interface{}   `json:"example,omitempty" yaml:"example,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// Object-related, here for struct compactness
	AdditionalPropertiesAllowed *bool `json:"-" multijson:"additionalProperties,omitempty" yaml:"-"`
	// Array-related, here for struct compactness
	UniqueItems bool `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	// Number-related, here for struct compactness
	ExclusiveMin bool `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	ExclusiveMax bool `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	// Properties
	Nullable        bool `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	ReadOnly        bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly       bool `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	XML             *XML `json:"xml,omitempty" yaml:"xml,omitempty"`
	Deprecated      bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// Number
	Min        *float64 `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Max        *float64 `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`

	// String
	MinLength       uint64  `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxLength       *uint64 `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Pattern         string  `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	compiledPattern *compiledPattern

	// Array
	MinItems uint64     `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	MaxItems *uint64    `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	Items    *SchemaRef `json:"items,omitempty" yaml:"items,omitempty"`

	// Object
	Required             []string              `json:"required,omitempty" yaml:"required,omitempty"`
	Properties           map[string]*SchemaRef `json:"properties,omitempty" yaml:"properties,omitempty"`
	MinProps             uint64                `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	MaxProps             *uint64               `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	AdditionalProperties *SchemaRef            `json:"-" multijson:"additionalProperties,omitempty" yaml:"-"`
	Discriminator        *Discriminator        `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
}

type compiledPattern struct {
	Regexp    *regexp.Regexp
	ErrReason string
}

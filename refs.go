package gopenapi

import (
	"github.com/goccy/go-yaml"
)

// CallbackRef
type CallbackRef struct {
	Ref   string
	Value Callback
}

func (value *CallbackRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *CallbackRef) UnMarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

// ExampleRef
type ExampleRef struct {
	Ref   string
	Value Example
}

func (value *ExampleRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *ExampleRef) UnMarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

// HeaderRef
type HeaderRef struct {
	Ref   string
	Value Header
}

func (value *HeaderRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *HeaderRef) UnMarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

// LinkRef
type LinkRef struct {
	Ref   string
	Value Link
}

func (value *LinkRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *LinkRef) UnMarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

// ParameterRef
type ParameterRef struct {
	Ref   string
	Value Parameter
}

func (value *ParameterRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *ParameterRef) UnMarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

// ResponseRef
type ResponseRef struct {
	Ref   string
	Value Response
}

func (value *ResponseRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *ResponseRef) UnmarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

// RequestBodyRef
type RequestBodyRef struct {
	Ref   string
	Value RequestBody
}

func (value *RequestBodyRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *RequestBodyRef) UnMarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

// SchemaRef
type SchemaRef struct {
	Ref   string
	Value Schema
}

func (value *SchemaRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *SchemaRef) UnmarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

// SecuritySchemeRef
type SecuritySchemeRef struct {
	Ref   string
	Value SecurityScheme
}

func (value *SecuritySchemeRef) MarshalYAML() ([]byte, error) {
	return MarshalRef(value.Ref, value.Value)
}

func (value *SecuritySchemeRef) UnMarshalYAML(data []byte) error {
	refProps := &refProps{}
	if err := yaml.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			value.Ref = ref
		}
	}
	return yaml.Unmarshal(data, &value.Value)
}

//////////////////////////////////////////////////////////////////////////////////////////
func MarshalRef(value string, otherwise interface{}) ([]byte, error) {
	if len(value) > 0 {
		// return json.Marshal(&refProps{
		return yaml.Marshal(&refProps{
			Ref: value,
		})
	}
	// return json.Marshal(otherwise)
	return yaml.Marshal(otherwise)
}

type refProps struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

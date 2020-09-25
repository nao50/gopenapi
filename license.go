package gopenapi

import (
	"context"
	"fmt"
)

// License information for the exposed API.
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.1.0.md#license-object
type License struct {
	// REQUIRED. The license name used for the API.
	Name string `json:"name" yaml:"name" validate:"required"`
	// An SPDX license expression for the API. The identifier field is mutually exclusive of the url field.
	Identifier string `json:"identifier" yaml:"identifier"`
	// A URL to the license used for the API. MUST be in the format of a URL. The url field is mutually exclusive of the identifier field.
	URL string `json:"url,omitempty" yaml:"url,omitempty" validate:"url"`
	// This object MAY be extended with Specification Extensions.
	ExtensionProps
}

func (value *License) validate(c context.Context) error {
	if value.Identifier != "" && value.URL != "" {
		errMsg := "Error: The identifier field is mutually exclusive of the url field."
		msg, err := yamlPathErrorMessage(value, errMsg, "identifier")
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("%s", msg)
	}
	return nil
}

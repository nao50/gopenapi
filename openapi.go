package gopenapi

import (
	"context"
	"fmt"
)

// OpenAPI is the root object of the OpenAPI document.
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.1.0.md#openapi-object
type OpenAPI struct {
	// REQUIRED. This string MUST be the version number of the OpenAPI Specification that the OpenAPI document uses. The openapi field SHOULD be used by tooling to interpret the OpenAPI document. This is not related to the API info.version string.
	OpenAPI string `json:"openapi" yaml:"openapi" validate:"required"`
	// REQUIRED. Provides metadata about the API. The metadata MAY be used by tooling as required.
	Info *Info `json:"info" yaml:"info" validate:"required"`
	// An array of Server Objects, which provide connectivity information to a target server. If the servers property is not provided, or is an empty array, the default value would be a Server Object with a url value of /.
	Servers Servers `json:"servers,omitempty" yaml:"servers,omitempty"`
	// The available paths and operations for the API.
	Paths Paths `json:"paths" yaml:"paths"`
	// webhooks
	//
	Security     SecurityRequirements `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         Tags                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	Components   Components           `json:"components,omitempty" yaml:"components,omitempty"`
	ExternalDocs *ExternalDocs        `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	ExtensionProps
}

func (doc *OpenAPI) Validate(c context.Context) error {
	// info validation
	if info := doc.Info; info != nil {
		if err := info.validate(c); err != nil {
			return fmt.Errorf("%s", err)
		}
	} else {
		errMsg := "Invalid \"OpenAPI Object\", \"info\" value expected JSON object."
		msg, err := yamlPathErrorMessage(doc, errMsg, "info")
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("%s", msg)
	}
	//

	return nil
}

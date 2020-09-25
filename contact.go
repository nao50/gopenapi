package gopenapi

import "context"

// Contact information for the exposed API.
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.1.0.md#contact-object
type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// 	The URL pointing to the contact information. MUST be in the format of a URL.
	URL string `json:"url,omitempty" yaml:"url,omitempty" validate:"url"`
	// The email address of the contact person/organization. MUST be in the format of an email address.
	Email string `json:"email,omitempty" yaml:"email,omitempty" validate:"email"`
	// This object MAY be extended with Specification Extensions.
	ExtensionProps
}

func (value *Contact) validate(c context.Context) error {
	return nil
}

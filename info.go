package gopenapi

import (
	"context"
	"fmt"
)

// Info is the object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.1.0.md#infoObject
type Info struct {
	// REQUIRED. The title of the API.
	Title string `json:"title" yaml:"title"`
	// A short summary of the API.
	Summary string `json:"summary" yaml:"summary"`
	// A short description of the API. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// A URL to the Terms of Service for the API. MUST be in the format of a URL.
	TermsOfService string `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty" validate:"url"`
	// The contact information for the exposed API.
	Contact *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	// The license information for the exposed API.
	License *License `json:"license,omitempty" yaml:"license,omitempty"`
	// REQUIRED. The version of the OpenAPI document (which is distinct from the OpenAPI Specification version or the API implementation version).
	Version string `json:"version" yaml:"version" validate:"required"`
	// This object MAY be extended with Specification Extensions.
	ExtensionProps
}

func (value *Info) validate(c context.Context) error {
	// objName := "Info Object"

	if contact := value.Contact; contact != nil {
		if err := contact.validate(c); err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	if license := value.License; license != nil {
		if err := license.validate(c); err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	// Version Required
	// if value.Version == "" {
	// 	bytes, err := yaml.Marshal(value)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	source, err := yamlSourceByPath(string(bytes), "$.version")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return fmt.Errorf("Invalid "+objName+", \"version\" value expected non-empty JSON string but actual %s:\n%s", value.Version, source)
	// }

	// Version Title
	// if value.Title == "" {
	// 	bytes, err := yaml.Marshal(value)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	source, err := yamlSourceByPath(string(bytes), "$.title")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return fmt.Errorf("Invalid "+objName+", \"title\" value expected non-empty JSON string but actual %s:\n%s", value.Version, source)
	// }

	return nil
}

// panic: [3:11] Key: 'Info.Version' Error:Field validation for 'Version' failed on the 'required' tag
//
//
//
//

// func (e *syntaxError) FormatError(p xerrors.Printer) error {
// 	var pp printer.Printer

// 	var colored, inclSource bool
// 	if mp, ok := p.(*myprinter); ok {
// 		colored = mp.colored
// 		inclSource = mp.inclSource
// 	}

// 	pos := fmt.Sprintf("[%d:%d] ", e.token.Position.Line, e.token.Position.Column)
// 	msg := pp.PrintErrorMessage(fmt.Sprintf("%s%s", pos, e.msg), colored)
// 	if inclSource {
// 		msg += "\n" + pp.PrintErrorToken(e.token, colored)
// 	}
// 	p.Print(msg)

// 	if e.verb == 'v' && e.state.Flag('+') {
// 		// %+v
// 		// print stack trace for debugging
// 		e.frame.Format(p)
// 	}
// 	return nil
// }

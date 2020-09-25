package gopenapi

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

func LoadWithValidation(filePath string) (*OpenAPI, error) {
	spec, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(spec)
	validate := validator.New()

	decoder := yaml.NewDecoder(
		buf,
		yaml.RecursiveDir(true),
		yaml.ReferenceDirs("spec"),
		yaml.Validator(validate),
	)
	var openAPI OpenAPI
	err = decoder.Decode(&openAPI)
	if err != nil {
		// fmt.Println(yaml.FormatError(err, true, true))
		return nil, fmt.Errorf("%s", yaml.FormatError(err, true, true))
	}

	// Validation
	ctx := context.Background()
	err = openAPI.Validate(ctx)
	if err != nil {
		return nil, err
	}

	return &openAPI, nil
}

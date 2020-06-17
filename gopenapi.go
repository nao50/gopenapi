package gopenapi

import (
	"bytes"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

func LoadWithValidation(filePath string) (*Document, error) {
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
	var document Document
	err = decoder.Decode(&document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

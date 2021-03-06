package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/moznion/gowrtr/generator"

	"github.com/nao50/gopenapi"
)

func genModel(doc *gopenapi.Document) {
	root := generator.NewRoot(
		generator.NewComment(" THIS CODE WAS AUTO GENERATED; DO NOT EDIT."),
		generator.NewPackage("model"),
		generator.NewNewline(),
	)

	root, _ = genStruct(doc, root)

	generated, err := root.Generate(0)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("gen/stub-server/model/model.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, generated)
}

//
func genStruct(doc *gopenapi.Document, root *generator.Root) (*generator.Root, error) {
	for schemaKey, schemaValue := range doc.Components.Schemas {
		// JSON SCHEMA RFC
		// https://json-schema.org/draft/2019-09/json-schema-core.html#rfc.section.4
		switch schemaValue.Type {
		case "null":
			return nil, fmt.Errorf("\"null\" is not supported at present")
		case "boolean":
			root = root.AddStatements(
				generator.NewRawStatement(fmt.Sprintf("type %s bool", strcase.ToCamel(schemaKey))),
			)
		case "object":
			generatedStruct := generator.NewStruct(strcase.ToCamel(schemaKey))
			for propertyKey, propertyValue := range schemaValue.Properties {
				// TODO: RFC どこにあるんだろ
				switch propertyValue.Type {
				case "string":
					generatedStruct = generatedStruct.AddField(strcase.ToCamel(propertyKey), "string", fmt.Sprintf("json:\"%s\"", propertyKey))
				case "integer":
					generatedStruct = generatedStruct.AddField(strcase.ToCamel(propertyKey), "int", fmt.Sprintf("json:\"%s\"", propertyKey))
				case "boolean":
					generatedStruct = generatedStruct.AddField(strcase.ToCamel(propertyKey), "bool", fmt.Sprintf("json:\"%s\"", propertyKey))
				}
			}
			root = root.AddStatements(generatedStruct)
		case "array":
			if schemaValue.Items.Ref != "" {
				// TODO: schemaValue.Items.Ref の存在確認と schemaKey として適当であるかのチェックをする
				slicedRef := strings.Split(schemaValue.Items.Ref, "/")
				root = root.AddStatements(
					generator.NewRawStatement(fmt.Sprintf("type %s []%s", strcase.ToCamel(schemaKey), strcase.ToCamel(slicedRef[len(slicedRef)-1]))),
				)
			} else {
				switch schemaValue.Items.Type {
				case "object":
					generatedStruct := generator.NewStruct(strcase.ToCamel(schemaKey) + "Singly")
					for itemsPropertyKey, itemsPropertyValue := range schemaValue.Items.Properties {
						switch itemsPropertyValue.Type {
						case "string":
							generatedStruct = generatedStruct.AddField(strcase.ToCamel(itemsPropertyKey), "string", fmt.Sprintf("json:\"%s\"", itemsPropertyKey))
						case "integer":
							generatedStruct = generatedStruct.AddField(strcase.ToCamel(itemsPropertyKey), "int", fmt.Sprintf("json:\"%s\"", itemsPropertyKey))
						case "boolean":
							generatedStruct = generatedStruct.AddField(strcase.ToCamel(itemsPropertyKey), "bool", fmt.Sprintf("json:\"%s\"", itemsPropertyKey))
						}
					}
					root = root.AddStatements(generatedStruct)
					root = root.AddStatements(
						generator.NewRawStatement(fmt.Sprintf("type %s []%s", strcase.ToCamel(schemaKey), strcase.ToCamel(schemaKey)+"Singly")),
					)
				case "number":
					root = root.AddStatements(
						generator.NewRawStatement(fmt.Sprintf("type %s []int", strcase.ToCamel(schemaKey))),
					)
				case "string":
					root = root.AddStatements(
						generator.NewRawStatement(fmt.Sprintf("type %s []string", strcase.ToCamel(schemaKey))),
					)
				case "boolean":
					root = root.AddStatements(
						generator.NewRawStatement(fmt.Sprintf("type %s []bool", strcase.ToCamel(schemaKey))),
					)
				}
			}
		case "number":
			root = root.AddStatements(
				generator.NewRawStatement(fmt.Sprintf("type %s int", strcase.ToCamel(schemaKey))),
			)
		case "string":
			root = root.AddStatements(
				generator.NewRawStatement(fmt.Sprintf("type %s string", strcase.ToCamel(schemaKey))),
			)
		}
	}
	return root, nil
}

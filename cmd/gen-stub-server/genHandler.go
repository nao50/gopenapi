package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/moznion/gowrtr/generator"

	"github.com/nao50/gopenapi"
)

func genHandler(doc *gopenapi.Document) {
	root := generator.NewRoot(
		generator.NewComment(" THIS CODE WAS AUTO GENERATED; DO NOT EDIT."),
		generator.NewPackage("handler"),
		generator.NewNewline(),
	)

	root, _ = genHandlerInterface(doc, root)

	generated, err := root.Generate(0)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("gen/stub-server/handler/handler.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, generated)
}

//
func genHandlerInterface(doc *gopenapi.Document, root *generator.Root) (*generator.Root, error) {
	generatedInterface := generator.NewInterface("StubHandlerInterface")
	// var generatedHandlerFuncStatements *generator.Root
	///////////////////////////////////////////////////////////////////////////
	// TODO: 下の２つは直す
	generatedStruct := generator.NewStruct("stubHandler").AddField(
		"application",
		"application.Application",
	)
	generatedConstructor := generator.NewFunc(
		nil,
		generator.NewFuncSignature("NewHandler").
			AddParameters(
				generator.NewFuncParameter("app", "application.Application"),
			).
			AddReturnTypes("StubHandlerInterface"),
	).AddStatements(
		generator.NewComment(" do something"),
		generator.NewReturnStatement("&stubHandler{application: app}"),
	)
	///////////////////////////////////////////////////////////////////////////

	for pathItemKey, pathItemValue := range doc.Paths {

		generatedInterface = generatedInterface.AddSignatures(
			generator.NewFuncSignature(strcase.ToCamel(pathItemValue.Get.OperationID)).
				AddParameters(
					generator.NewFuncParameter("w", "http.ResponseWriter"),
					generator.NewFuncParameter("r", "*http.Request"),
				),
		)

		// generatedHandlerFunc = generator.NewFunc(
		// 	generator.NewFuncReceiver("sh", "*stubHandler"),
		// 	generator.NewFuncSignature(strcase.ToCamel(pathItemValue.Get.OperationID)).
		// 		AddParameters(
		// 			generator.NewFuncParameter("w", "http.ResponseWriter"),
		// 			generator.NewFuncParameter("r", "*http.Request"),
		// 		),
		// )

		root = root.AddStatements(
			generator.NewFunc(
				generator.NewFuncReceiver("sh", "*stubHandler"),
				generator.NewFuncSignature(strcase.ToCamel(pathItemValue.Get.OperationID)).
					AddParameters(
						generator.NewFuncParameter("w", "http.ResponseWriter"),
						generator.NewFuncParameter("r", "*http.Request"),
					),
			).AddStatements(
				generator.NewRawStatement(`fmt.Println("hello, world!")`),
			),
		)

		switch {
		// case pathItemValue.Get != nil || pathItemValue.Put != nil || pathItemValue.Post != nil || pathItemValue.Delete != nil || pathItemValue.Options != nil || pathItemValue.Head != nil || pathItemValue.Patch != nil || pathItemValue.Trace != nil:
		case pathItemValue.Get != nil:
			fmt.Println("BBBBBB: ", pathItemKey)
			fmt.Printf("CCCCCC: %+v\n", pathItemValue.Get.OperationID)
		default:
		}
	}

	root = root.
		AddStatements(generatedInterface).
		AddStatements(generatedStruct).
		AddStatements(generatedConstructor).
		// AddStatements(generatedHandlerFuncStatements).
		Gofmt("-s").
		Goimports()

	return root, nil
}

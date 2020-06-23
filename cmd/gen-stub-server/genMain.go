package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moznion/gowrtr/generator"

	"github.com/nao50/gopenapi"
)

func genMain(doc *gopenapi.Document) {
	root := generator.NewRoot(
		generator.NewComment(" THIS CODE WAS AUTO GENERATED; DO NOT EDIT."),
		generator.NewPackage("main"),
		generator.NewNewline(),
	)

	root, _ = genMainFunction(doc, root)

	generated, err := root.Generate(0)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("gen/stub-server/cmd/main.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, generated)
}

//
func genMainFunction(doc *gopenapi.Document, root *generator.Root) (*generator.Root, error) {
	wd, _ := os.Getwd()
	fmt.Println("wd: ", wd)

	root = root.AddStatements(
		generator.NewImport("./application", "./handler"),
		generator.NewFunc(
			nil,
			generator.NewFuncSignature("main"),
		).AddStatements(
			generator.NewRawStatement(`app := application.NewPetApplication()`),
			generator.NewRawStatement(`petHandler := handler.NewPetHandler(app)`),

			generator.NewRawStatement(`r := chi.NewRouter()`),
			generator.NewRawStatement(`r.Use(middleware.Logger)`),
			generator.NewRawStatement(`r.Get("/pets", petHandler.ListPets)`),

			generator.NewRawStatement(`http.ListenAndServe(":5051", r)`),
		),
	).
		Gofmt("-s").
		Goimports()

	return root, nil
}

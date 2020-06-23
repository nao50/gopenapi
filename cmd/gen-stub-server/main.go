package main

import (
	"fmt"
	"os"

	"github.com/nao50/gopenapi"
)

func main() {
	doc, err := gopenapi.LoadWithValidation("./spec/petstore.yaml")
	if err != nil {
		fmt.Println(err)
	}

	//
	genDirectory()
	genModel(doc)
	genHandler(doc)

	genMain(doc)
}

//
func genDirectory() {
	if _, err := os.Stat("gen/stub-server/"); os.IsNotExist(err) {
		os.MkdirAll("gen/stub-server", 0777)
	}

	if _, err := os.Stat("gen/stub-server/application"); os.IsNotExist(err) {
		os.MkdirAll("gen/stub-server/application", 0777)
	}

	if _, err := os.Stat("gen/stub-server/cmd"); os.IsNotExist(err) {
		os.MkdirAll("gen/stub-server/cmd", 0777)
	}

	if _, err := os.Stat("gen/stub-server/config"); os.IsNotExist(err) {
		os.MkdirAll("gen/stub-server/config", 0777)
	}

	if _, err := os.Stat("gen/stub-server/handler"); os.IsNotExist(err) {
		os.MkdirAll("gen/stub-server/handler", 0777)
	}

	if _, err := os.Stat("gen/stub-server/model"); os.IsNotExist(err) {
		os.MkdirAll("gen/stub-server/model", 0777)
	}
}

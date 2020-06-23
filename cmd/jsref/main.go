package main

import (
	"fmt"

	"github.com/nao50/gopenapi"
)

func main() {
	doc, err := gopenapi.LoadWithValidation("./spec/petstore.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", doc)
}

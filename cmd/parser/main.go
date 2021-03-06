package main

import (
	"fmt"

	"github.com/nao50/gopenapi"
)

func main() {
	doc, err := gopenapi.LoadWithValidation("./spec/petstore.yaml")
	// doc, err := gopenapi.LoadWithValidation("./petstore.yaml")
	if err != nil {
		// fmt.Println(err)
		panic(err)
		// fmt.Errorf("err %s", err)
	}

	// ctx := context.Background()
	// err = doc.Validate(ctx)
	// if err != nil {
	// 	fmt.Println("Validate Error:", err)
	// }

	// fmt.Printf("%+v \n", doc)
	for _, PathsValue := range doc.Paths {
		// fmt.Printf("%+v \n", PathsValue.Get.Responses)
		for _, ResponsesValue := range PathsValue.Get.Responses {
			// fmt.Printf("%+v \n", ResponsesValue.Value.Content)
			for _, ContentValue := range ResponsesValue.Value.Content {
				fmt.Printf("%+v \n", ContentValue.Schema)
				// fmt.Printf("%+v \n", ContentValue.Schema.Ref)
			}
		}
	}
}

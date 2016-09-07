// test project main.go
package main

import (
	"fmt"

	"github.com/ahmdrz/microsoft-vision-golang"
)

func main() {
	fmt.Println("Hello World!")
	vis, err := vision.New("<KEY>")
	if err != nil {
		panic(err)
	}

	result, err := vis.OCR("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg",
		vision.OCROption{
			Language: vision.LANG_English,
		})
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

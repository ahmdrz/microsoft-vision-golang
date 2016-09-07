// test project main.go
package main

import (
	"fmt"

	"github.com/ahmdrz/microsoft-vision-golang"
)

func main() {
	fmt.Println("Hello World!")
	vision, err := vision.New("<KEY>")
	if err != nil {
		panic(err)
	}

	result, err := vision.Tag("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

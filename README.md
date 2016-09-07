# microsoft-vision-golang
A library for communication with microsoft cognitive services API written in Golang


### Dependencies 

Nothing ! I wrote this library on Golang 1.5.4 linux/amd64

So it will work on Golang 1.5 or later

### <b style="color:red;">RTFM !</b>

The full documention of microsoft cognitive services API is in [this link](https://www.microsoft.com/cognitive-services/en-us/computer-vision-api/documentation) and API reference is in [here](https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/)

### How to use ?

First of all download the library 

```bash
go get github.com/ahmdrz/microsoft-vision-golang
```

Here is a simple code

If you haven't key visit [here](https://www.microsoft.com/cognitive-services/en-US/subscriptions)

Tags :

```go
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
```

Analyze :

```go
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

	result, err := vis.Analyze("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg",
		vision.VisualFeatures{
			Adult: true,
			Tags:  true,
		})
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
```

### License 
MIT License

>Copyright (c) 2016 Ahmadreza Zibaei

>Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

>The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

>THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

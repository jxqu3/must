# must
A really, really simple (30 loc) library to make it easier to handle errors in go

instead of

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("yourfile.txt")
	if err != nil {
		panic(err)
	}
	fileLog, err := os.Open("yourfile.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fileFatal, err := os.Open("yourfile.txt")
	if err != nil {
		log.Fatal(err)

	}
	println(file.Name())
}
```

you can do:
```go
package main

import (
	"os"

	"github.com/checkm4ted/must"
)

func main() {
	file := must.Panic(os.Open("yourfile.txt")) // To log.Panic() the error
	fileLog := must.Log(os.Open("yourfile.txt")) // To log.Print() the error
	fileFatal := must.Fatal(os.Open("yourfile.txt")) // To log.Fatal() the error
	println(file.Name())
}
```
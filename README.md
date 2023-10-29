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
	file := must.Must(os.Open("yourfile.txt")) // To panic
	fileLog := must.MustLog(os.Open("yourfile.txt")) // To log.Print()
	fileFatal := must.MustFatal(os.Open("yourfile.txt")) // To log.Fatal()
	println(file.Name())
}
```
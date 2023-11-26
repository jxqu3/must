# must
A really, really simple (~30 loc) library to make it easier to handle errors in go

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
		panic("ERROR:", err)
	}
	fileLog, err := os.Open("yourfile.txt")
	if err != nil {
		log.Print("ERROR:", err)
	}
	fileFatal, err := os.Open("yourfile.txt")
	if err != nil {
		log.Fatal("ERROR:", err)
	}
  fileCustom, err := os.Open("yourfile.txt")
	if err != nil {
		// Do some custom action here
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
  
    // Sets a custom callback for the must.C function, in case you want to customise what you do with your errors.
    must.SetErrorRule(func(err error) {
        // Do some custom action here
    })
	fileCustom := must.C(os.Open("yourfile.txt")) // To use the custom callback
	fileCustom := must.C2(SomeFunctionWith2ReturnTypes("yourfile.txt"))
	fileCustom := must.C3(SomeFunctionWith3ReturnTypes("yourfile.txt"))
	must.C0(os.Chdir("dir")) // when you only want to handle the error and there are no return types (only returns an error)
	println(file.Name())
}
```
Or instead of 
```go
file, err := os.Open("anotherfile.txt")
if err != nil {
    log.Fatal(err)
}
```
```go
file, err := os.Open("anotherfile.txt")
must.H(err) // Handles the error using the SetError rule callback (log.Fatal by default)
```

## get it
Run `go get https://github.com/checkm4ted/must`
import "github.com/checkm4ted/must" in your files

package must

import (
	"os"
)

func ExampleC() {
	C0(os.Chdir("dir"))
}

func ExampleHandle() {
	Handle(os.Chdir("dir"))
	H(os.Chdir("dir"))
}

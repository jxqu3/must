package must

import (
	"fmt"
	"log"
)

// Must panics if err is not nil
func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

// MustFatal log.Fatal err if err is not nil
func MustFatal[T any](obj T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

// MustLog logs err if err is not nil
func MustLog[T any](obj T, err error) T {
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	return obj
}

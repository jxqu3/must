package must

import (
	"log"
)

// Panic panics if err is not nil
func Panic[T any](obj T, err error) T {
	if err != nil {
		log.Panic(err)
	}
	return obj
}

// Fatal log.Fatal err if err is not nil
func Fatal[T any](obj T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

// Log logs err if err is not nil
func Log[T any](obj T, err error) T {
	if err != nil {
		log.Print("ERROR: ", err)
	}
	return obj
}

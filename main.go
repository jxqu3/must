package must

import (
	"log"
)

var ErrorRule func(error)

func init() {
	SetErrorRule(func(err error) {
		log.Fatal("THIS IS THE DEFAULT ERRORRULE, USE SetErrorRule(handler)", err)
	})
}

// Sets the callback for the C (custom) function
func SetErrorRule(handler func(error)) {
	ErrorRule = handler
}

// Panic panics if err is not nil
func Panic[T any](obj T, err error) T {
	if err != nil {
		log.Panic("ERROR: ", err)
	}
	return obj
}

// Fatal log.Fatal err if err is not nil
func Fatal[T any](obj T, err error) T {
	if err != nil {
		log.Fatal("ERROR: ", err)
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

// C (Custom) action, uses the ErrorRule callback. SetErrorRule() sets that callback.
func C[T any](obj T, err error) T {
	if err != nil {
		ErrorRule(err)
	}
	return obj
}

// Package must provides functions to make error handling easier, by passing them functions that may return errors, or the errors themselves.
package must

import (
	"log"
)

var ErrorRule func(error)

func init() {
	SetErrorRule(func(err error) {
		log.Fatal("ERROR: ", err)
	})
}

// Sets the callback for the C (custom) functions, and for the H/Handle() functions.
//
//	SetErrorRule(func(err error) {
//		log.Fatal("ERROR: ", err)
//	})
func SetErrorRule(handler func(error)) {
	ErrorRule = handler
}

// Panic panics if err is not nil, ant prints it (Log.Panic())
func Panic[T any](obj T, err error) T {
	if err != nil {
		log.Panic("ERROR: ", err)
	}
	return obj
}

// Fatal() runs log.Fatal(err) if err is not nil
func Fatal[T any](obj T, err error) T {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	return obj
}

// Log() logs err if err is not nil
func Log[T any](obj T, err error) T {
	if err != nil {
		log.Print("ERROR: ", err)
	}
	return obj
}

// C0 (Custom0) action, no return types, error only. uses the ErrorRule callback. SetErrorRule() sets that callback. Defaults to
//
//	log.Fatal("Error" + err)
func C0(err error) {
	if err != nil {
		ErrorRule(err)
	}
}

// C (Custom) action, uses the ErrorRule callback. SetErrorRule() sets that callback. Defaults to
//
//	log.Fatal("Error" + err)
func C[T any](obj T, err error) T {
	if err != nil {
		ErrorRule(err)
	}
	return obj
}

// C2 (Custom2) action, with two return types. uses the ErrorRule callback. SetErrorRule() sets that callback. Defaults to
//
//	log.Fatal("Error" + err)
func C2[T, E any](obj T, obj2 E, err error) (T, E) {
	if err != nil {
		ErrorRule(err)
	}
	return obj, obj2
}

// C3 (Custom3) action, with three return types. uses the ErrorRule callback. SetErrorRule() sets that callback. Defaults to
//
//	log.Fatal("Error" + err)
func C3[T, E, D any](obj T, obj2 E, obj3 D, err error) (T, E, D) {
	if err != nil {
		ErrorRule(err)
	}
	return obj, obj2, obj3
}

// Handle() handles a given error if it's not nil. uses the ErrorRule callback. SetErrorRule() sets that callback. Defaults to:
//
//	log.Fatal("Error" + err)
//
//	err := os.Chdir("dirthatdoesnotexist")
//	Handle(err)
func Handle(err error) {
	if err != nil {
		ErrorRule(err)
	}
}

// H() is an alias to Handle(). Handles a given error if it's not nil. uses the ErrorRule callback. SetErrorRule() sets that callback. Defaults to:
//
//	log.Fatal("Error" + err)
//
//	err := os.Chdir("dirthatdoesnotexist")
//	Handle(err)
func H(err error) {
	Handle(err)
}

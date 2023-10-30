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

// C0 (Custom0) action, no return types, error only. uses the ErrorRule callback. SetErrorRule() sets that callback.
func C0(err error) {
	if err != nil {
		ErrorRule(err)
	}
}

// C (Custom) action, uses the ErrorRule callback. SetErrorRule() sets that callback.
func C[T any](obj T, err error) T {
	if err != nil {
		ErrorRule(err)
	}
	return obj
}

// C2 (Custom2) action, with two return types. uses the ErrorRule callback. SetErrorRule() sets that callback.
func C2[T, E any](obj T, obj2 E, err error) (T, E) {
	if err != nil {
		ErrorRule(err)
	}
	return obj, obj2
}

// C3 (Custom3) action, with three return types. uses the ErrorRule callback. SetErrorRule() sets that callback.
func C3[T, E, D any](obj T, obj2 E, obj3 D, err error) (T, E, D) {
	if err != nil {
		ErrorRule(err)
	}
	return obj, obj2, obj3
}

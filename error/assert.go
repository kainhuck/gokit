package error

import "fmt"

func AssertPanic(condition bool, errMsg string) {
	if !condition {
		panic(errMsg)
	}
}

func Assert(condition bool, errMsg string) {
	AssertPanic(condition, errMsg)
}

func AssertError(condition bool, errMsg string) error {
	if !condition {
		return fmt.Errorf(errMsg)
	}

	return nil
}
/* Copyright 2019 Kilobit Labs Inc. */

package assert // import "kilobit.ca/go/tested/assert"

import "testing"
import "fmt"
import "runtime"

// Simple shallow comparison of expected and actual values.
//
func Expect(tb testing.TB, expected interface{}, actual interface{}) {

	if expected != actual {
		tb.Errorf("Expected %v, Got %v", expected, actual)
	}
}

// Uses the string '%#v' representation for matching as there are
// issues matching with reflect.DeepEqual returning false negatives.
//
func ExpectDeep(tb testing.TB, expected interface{}, actual interface{}) {

	exps := fmt.Sprintf("%#v", expected)
	acts := fmt.Sprintf("%#v", actual)

	if exps != acts {
		tb.Errorf("Expected %#v, Got %#v", expected, actual)
	}
}

// Oneline check for errors.
//
// Based on a helper in benbjohnson/testing.
//
func Ok(tb testing.TB, err error) {

	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: Unexpected error: %s\033[39m\n\n", file, line, err)
		tb.FailNow()
	}
}

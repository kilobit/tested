/* Copyright 2019 Kilobit Labs Inc. */

package assert // import "kilobit.ca/go/tested/assert"

import "testing"
import "fmt"
import "runtime"
import fp "path/filepath"

// Simple shallow comparison of expected and actual values.
//
func Expect(tb testing.TB, expected interface{}, actual interface{}) {

	if expected != actual {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: Expected %v, Got %v\033[39m\n\n",
			fp.Base(file), line, expected, actual)
		tb.FailNow()
	}
}

// Uses the string '%#v' representation for matching as there are
// issues matching with reflect.DeepEqual returning false negatives.
//
func ExpectDeep(tb testing.TB, expected interface{}, actual interface{}) {

	exps := fmt.Sprintf("%#v", expected)
	acts := fmt.Sprintf("%#v", actual)

	if exps != acts {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\nExpected:\n\t%#v\nGot:\n\t%#v\033[39m\n\n",
			fp.Base(file), line, expected, actual)
		tb.FailNow()
	}
}

// Oneline check for errors.
//
// Based on a helper in benbjohnson/testing.
//
func Ok(tb testing.TB, err error) {

	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: Unexpected error\n\n%s\033[39m\n\n",
			fp.Base(file), line, err)
		tb.FailNow()
	}
}

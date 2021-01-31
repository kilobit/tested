/* Copyright 2019 Kilobit Labs Inc. */

// Standard assertions.
//
// Assertions accept additional arguments that will be stringified and
// printed in the event the assertion fails.
//
package assert // import "kilobit.ca/go/tested/assert"

import "testing"
import "fmt"
import "strings"
import "runtime"
import fp "path/filepath"

// Simple shallow comparison of expected and actual values.
//
func Expect(tb testing.TB, expected interface{}, actual interface{}, extra ...interface{}) {

	if expected != actual {
		_, file, line, _ := runtime.Caller(1)
		mstr := stringify(extra...)
		fmt.Printf("\033[31m%s:%d: Expected %v, Got %v\n\n%s\033[39m\n\n",
			fp.Base(file), line, expected, actual, mstr)
		tb.FailNow()
	}
}

// Uses the string '%#v' representation for matching as there are
// issues matching with reflect.DeepEqual returning false negatives.
//
func ExpectDeep(tb testing.TB, expected interface{}, actual interface{}, extra ...interface{}) {

	exps := fmt.Sprintf("%#v", expected)
	acts := fmt.Sprintf("%#v", actual)

	if exps != acts {
		_, file, line, _ := runtime.Caller(1)
		mstr := stringify(extra...)
		fmt.Printf("\033[31m%s:%d:\nExpected:\n\t%#v\nGot:\n\t%#v\n\n%s\033[39m\n\n",
			fp.Base(file), line, expected, actual, mstr)
		tb.FailNow()
	}
}

// Oneline check for errors.
//
// Based on a helper in benbjohnson/testing.
//
func Ok(tb testing.TB, err error, extra ...interface{}) {

	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		mstr := stringify(extra...)
		fmt.Printf("\033[31m%s:%d: Unexpected Error (%T)\n\n%s\n\n%s\033[39m\n\n",
			fp.Base(file), line, err, err, mstr)
		tb.FailNow()
	}
}

// Convert arguments of any type to a string suitable for including in
// test reports.
//
func stringify(args ...interface{}) string {

	strs := []string{}
	for _, val := range args {

		var str string
		switch v := val.(type) {
		case bool:
			str = fmt.Sprintf("%t", v)
		case string:
			str = v
		case fmt.Stringer:
			str = v.String()
		case error:
			str = fmt.Sprintf("%T: %s", v, v.Error())
		default:
			str = fmt.Sprintf("%#v", v)
		}

		strs = append(strs, str)
	}

	return strings.Join(strs, "\n")
}

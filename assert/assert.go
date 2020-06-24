/* Copyright 2019 Kilobit Labs Inc. */

package assert // import "kilobit.ca/go/tested/assert"

import "testing"
import "fmt"

// Simple shallow comparison of expected and actual values.
//
func Expect(t *testing.T, expected interface{}, actual interface{}) {

	if expected != actual {
		t.Errorf("Expected %v, Got %v", expected, actual)
	}
}

// Uses the string '%#v' representation for matching as there are
// issues matching with reflect.DeepEqual returning false negatives.
//
func ExpectDeep(t *testing.T, expected interface{}, actual interface{}) {

	exps := fmt.Sprintf("%#v", expected)
	acts := fmt.Sprintf("%#v", actual)

	if exps != acts {
		t.Errorf("Expected %#v, Got %#v", expected, actual)
	}
}


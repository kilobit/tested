/* Copyright 2019 Kilobit Labs Inc. */

package assert // import "kilobit.ca/go/tested/assert"

import "testing"
import "reflect"

func Expect(t *testing.T, expected interface{}, actual interface{}) {

	if expected != actual {
		t.Errorf("Expected %v, Got %v", expected, actual)
	}
}

func ExpectDeep(t *testing.T, expected interface{}, actual interface{}) {

	if ! reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, Got %v", expected, actual)
	}
}

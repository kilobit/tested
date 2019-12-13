/* Copyright 2019 Kilobit Labs Inc. */

package assert // import "kilobit.ca/go/tested/assert"

import "testing"

func TestAssertTest(t *testing.T) {

	Expect(t, true, true)
}

func TestExpectDeep(t *testing.T) {

	exp := map[string]int{"foo": 1, "bar": 42}
	act := map[string]int{"foo": 1, "bar": 42}

	ExpectDeep(t, exp, act)
}

func TestExpectDeepNested(t *testing.T) {

	exp := map[string]interface{}{"foo": 1, "bar": map[string]bool{"nested": true}}
	act := map[string]interface{}{"foo": 1, "bar": map[string]bool{"nested": true}}

	ExpectDeep(t, exp, act)
}

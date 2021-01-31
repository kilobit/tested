/* Copyright 2021 Kilobit Labs Inc. */

package assert_test

import "testing"
import "errors"
import . "kilobit.ca/go/tested/assert"

func TestAssertTest(t *testing.T) {

	Expect(t, true, true)
}

func TestAssertFailTest(t *testing.T) {

	t.Skip("Normally disabled as a meta-test.")

	Expect(t, true, false, "This test is normally skipped.")
}

func TestExpectDeep(t *testing.T) {

	exp := map[string]int{"foo": 1, "bar": 42}
	act := map[string]int{"foo": 1, "bar": 42}

	ExpectDeep(t, exp, act)
}

func TestExpectDeepFail(t *testing.T) {

	t.Skip("Normally disabled as a meta-test.")

	exp := map[string]int{"foo": 1, "bar": 42}
	act := map[string]int{"foo": 1, "bar": 43}

	ExpectDeep(t, exp, act, "This test is normally skipped.")
}

func TestExpectDeepNested(t *testing.T) {

	exp := map[string]interface{}{"foo": 1, "bar": map[string]bool{"nested": true}}
	act := map[string]interface{}{"foo": 1, "bar": map[string]bool{"nested": true}}

	ExpectDeep(t, exp, act)
}

func TestOk(t *testing.T) {

	var err error = nil
	Ok(t, err, "Won't be printed.")
}

func TestNotOk(t *testing.T) {

	t.Skip("Normally disabled as a meta-test.")

	var err error = errors.New("Test Error.")
	Ok(t, err, "This test is normally skipped.")
}

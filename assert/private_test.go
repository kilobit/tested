/* Copyright 2021 Kilobit Labs Inc. */

package assert

import "strings"
import "unicode"

import "testing"

func TestPrivateTest(t *testing.T) {
	Expect(t, true, true)
}

func isNotPrint(r rune) bool {
	return !unicode.IsPrint(r)
}

func TestStringify(t *testing.T) {

	tests := []struct {
		args []interface{}
		exp  string
	}{
		{[]interface{}{}, ""},

		{
			[]interface{}{
				"foo",
				1,
				true,
				[]byte{'t'},
			},
			"foo\n1\ntrue\n[]byte{0x74}\n",
		},
	}

	for _, test := range tests {

		act := stringify(test.args...)
		Expect(t,
			strings.TrimFunc(test.exp, isNotPrint),
			strings.TrimFunc(act, isNotPrint), test.args,
		)
	}
}

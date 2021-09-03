package str

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{s: "", exp: ""},
		{s: "x", exp: "x"},
		{s: "ab", exp: "ba"},
		{s: "ћирилица", exp: "ацилирић"},
	}

	for _, test := range tests {
		res := Reverse(test.s)
		if res != test.exp {
			t.Errorf("test %q failed, expected %q, but got %q", test.s, test.exp, res)
		}
	}
}

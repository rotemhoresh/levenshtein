package levenshtein

import "testing"

var tests = []struct {
	A string
	B string
	D int
}{
	{
		A: "",
		B: "",
		D: 0,
	},
	{
		A: "flaw",
		B: "lawn",
		D: 2,
	},
}

func TestRecursiveDistance(t *testing.T) {
	for _, test := range tests {
		if d := RecursiveDistance(test.A, test.B); d != test.D {
			t.Errorf("RecursiveDistance(%s, %s) = %d, want %d", test.A, test.B, d, test.D)
		}
	}
}

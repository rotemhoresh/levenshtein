package levenshtein

import "testing"

func TestRecursiveDistance(t *testing.T) {
	for _, test := range tests {
		if d := RecursiveDistance(test.A, test.B); d != test.D {
			t.Errorf("RecursiveDistance(%s, %s) = %d, want %d", test.A, test.B, d, test.D)
		}
	}
}

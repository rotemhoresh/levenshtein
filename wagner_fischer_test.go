package levenshtein

import "testing"

func TestFullMatrixDistance(t *testing.T) {
	for _, test := range tests {
		if d := FullMatrixDistance(test.A, test.B); d != test.D {
			t.Errorf("FullMatrixDistance(%s, %s) = %d, want %d", test.A, test.B, d, test.D)
		}
	}
}

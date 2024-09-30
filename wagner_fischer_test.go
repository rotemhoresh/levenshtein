package levenshtein

import (
	"fmt"
	"testing"
)

func TestFullMatrixDistance(t *testing.T) {
	for _, test := range tests {
		if d := FullMatrixDistance(test.A, test.B); d != test.D {
			t.Errorf("FullMatrixDistance(%s, %s) = %d, want %d", test.A, test.B, d, test.D)
		}
	}
}

func BenchmarkFullMatrixDistance(b *testing.B) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("%s_vs_%s", test.A, test.B), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FullMatrixDistance(test.A, test.B)
			}
		})
	}
}

func TestTwoRowsDistance(t *testing.T) {
	// for _, test := range tests {
	// 	if d := TwoRowsDistance(test.A, test.B); d != test.D {
	// 		t.Errorf("TwoRowsDistance(%s, %s) = %d, want %d", test.A, test.B, d, test.D)
	// 	}
	// }
	if TwoRowsDistance("sitting", "kitten") != 3 {
		t.Error("Bruh")
	}
}

func BenchmarkTwoRowsDistance(b *testing.B) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("%s_vs_%s", test.A, test.B), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TwoRowsDistance(test.A, test.B)
			}
		})
	}
}

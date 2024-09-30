package levenshtein

import (
	"fmt"
	"testing"
)

func TestRecursiveDistance(t *testing.T) {
	for _, test := range tests {
		if d := RecursiveDistance(test.A, test.B); d != test.D {
			t.Errorf("RecursiveDistance(%s, %s) = %d, want %d", test.A, test.B, d, test.D)
		}
	}
}

func BenchmarkRecursiveDistance(b *testing.B) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("%s_vs_%s", test.A, test.B), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RecursiveDistance(test.A, test.B)
			}
		})
	}
}

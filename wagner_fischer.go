package levenshtein

// FullMatrixDistance calculate the distance between two strings using the Levenshtein
// Distance algorithm, with the [Wagner Fischer] algorithm.
//
// The Wagner-Fischer algorithm can be optimized by utilizing only two matrix rows,
// which is implemented by [TwoRowsDistance].
//
// [Wagner Fischer]: https://en.wikipedia.org/wiki/Levenshtein_distance#Iterative_with_full_matrix
func FullMatrixDistance(a, b string) int {
	la, lb := len(a), len(b)
	d := make([][]int, la+1)
	for i := 0; i <= la; i++ {
		d[i] = make([]int, lb+1)
	}

	for i := 0; i <= la; i++ {
		d[i][0] = i
	}
	for j := 1; j <= lb; j++ {
		d[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			substitutionCost := 1
			if a[i-1] == b[j-1] {
				substitutionCost = 0
			}
			d[i][j] = min(
				d[i-1][j]+1,                  // deletion
				d[i][j-1]+1,                  // insertion
				d[i-1][j-1]+substitutionCost, // substitution
			)
		}
	}

	return d[la][lb]
}

// TwoRowsDistance calculate the distance between two strings using the Levenshtein
// Distance algorithm, with the Wagner–Fischer algorithm, with the [two matrix rows]
// optimization.
//
// This is an optimization of the [FullMatrixDistance] function.
//
// [two matrix rows]: https://en.wikipedia.org/wiki/Levenshtein_distance#Iterative_with_two_matrix_rows
func TwoRowsDistance(a, b string) int {
	la, lb := len(a), len(b)
	prev, curr := make([]int, lb+1), make([]int, lb+1)

	for j := 0; j <= lb; j++ {
		prev[j] = j
	}

	for i := 0; i < la; i++ {
		curr[0] = i + 1

		for j := 0; j < lb; j++ {
			substitutionCost := 1
			if a[i] == b[j] {
				substitutionCost = 0
			}
			curr[j+1] = min(
				prev[j+1]+1,              // deletion
				curr[j]+1,                // insertion
				prev[j]+substitutionCost, // substitution
			)
		}

		prev, curr = curr, prev
	}

	return prev[lb]
}

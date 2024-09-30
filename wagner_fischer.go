package levenshtein

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

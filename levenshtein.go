package levenshtein

// RecursiveDistance calculates the edit distance between two strings using the
// Levenshtein Distance algorithm.
//
// It implements the original, recursive algorithm, which is not efficient, thus
// should not be used. It was created for learning purposes.
func RecursiveDistance(a, b string) int {
	la, lb := len(a), len(b)
	if lb == 0 {
		return la
	}
	if la == 0 {
		return lb
	}
	if head(a) == head(b) {
		return RecursiveDistance(tail(a), tail(b))
	}
	return 1 + min(
		RecursiveDistance(tail(a), b),
		RecursiveDistance(a, tail(b)),
		RecursiveDistance(tail(a), tail(b)),
	)
}

func head(s string) string {
	if len(s) == 0 {
		return s
	}
	return s[:1]
}

func tail(s string) string {
	if len(s) <= 1 {
		return ""
	}
	return s[1:]
}

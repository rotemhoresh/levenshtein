// package levenshtein contains multiple implementations for the [Levenshtein Distance]
// algorithm.
//
// [Levenshtein Distance]: https://en.wikipedia.org/wiki/Levenshtein_distance
package levenshtein

// Distance is an alias for [TwoRowsDistance] as it is the most performant implementation
// of the Levenshtein Distance algorithm in this package.
var Distance = TwoRowsDistance

# levenshtein

[![Go Reference](https://pkg.go.dev/badge/github.com/rotemhoresh/levenshtein.svg)](https://pkg.go.dev/github.com/rotemhoresh/levenshtein)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/rotemhoresh/levenshtein)](https://goreportcard.com/report/github.com/rotemhoresh/levenshtein)


Go implementations for the [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance).

There are existing packages that implement Levenshtein Distance, I have created this package for learning.

## Overview

Implementations:
- Full Matrix - the [Wagner-Fischer algorithm](https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm) for computing the Levenshtein Distance, which utilizes dynamic programing.
- Two Rows - an optimization for the Wagner-Fischer algorithm, which takes only two matrix rows regardless of the input. this is the most performant implementation in this package.
- Recursive - the original implementaion. an inefficient implementation, created for learning purposes.

When using this package, you should use the `Distance` - alias for the Two Rows implementaion.

### Example

Using this package:

```bash
go get github.com/rotemhoresh/levenshtein
```

```go
d := levenshtein.Distance("kitten", "sitting")
fmt.Printf("edit distance between 'kitten' and 'sitting' is %d", d)

// Output:
// edit distance between 'kitten' and 'sitting' is 3
```

## Testing and Benchmarking

Test and Benchmark fucntions are available for all implementations.

```bash
# test
go test

# benchmark
go test -bench=. -benchmem
```

## License

MIT

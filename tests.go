package levenshtein

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
	{
		A: "sitting",
		B: "kitten",
		D: 3,
	},
	{
		A: "Sunday",
		B: "Saturday",
		D: 3,
	},
	{
		A: "abc",
		B: "abc",
		D: 0,
	},
	{
		A: "a very long string to test the performance",
		B: "a very long string to test the benchmark performance",
		D: 10,
	},
}

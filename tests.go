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
}

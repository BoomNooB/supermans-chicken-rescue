package constant

const (
	FirstLine  = "firstLine"
	SecondLine = "secondLine"
)

var (
	// Assume the first line and second line as slice
	TC1 = map[string][]int{
		FirstLine:  {5, 5},             // [n,k]
		SecondLine: {2, 5, 10, 12, 15}, // [position 1D-axis]
	}

	TC2 = map[string][]int{
		FirstLine:  {6, 10},
		SecondLine: {1, 11, 30, 34, 35, 37},
	}

	TC3 = map[string][]int{
		FirstLine:  {10, 100},
		SecondLine: {1, 2, 3, 4, 5, 6, 7, 8, 9, 120},
	}
)

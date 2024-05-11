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

	TC4 = map[string][]int{
		FirstLine:  {10, 1200},
		SecondLine: {1, 2, 3, 4, 5, 6, 7, 8, 35, 120},
	}

	TC5 = map[string][]int{
		FirstLine:  {10, 0},
		SecondLine: {1, 2, 3, 4, 5, 6, 7, 8, 35, 120},
	}

	TC6 = map[string][]int{
		FirstLine:  {11, 1000000000000000000},
		SecondLine: {1, 19, 22, 23, 24, 25, 26, 255, 3367, 3452134, 43563452345},
	}
)

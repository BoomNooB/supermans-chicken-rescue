package main

import (
	"fmt"
	"p2/constant"
	"slices"
)

func main() {
	testCases := []map[string][]int{
		constant.TC1,
		constant.TC2,
		constant.TC3,
		constant.TC4,
		constant.TC5,
		constant.TC6,
	}

	for _, testCase := range testCases {
		fmt.Println(howManyChickenAreRescue(testCase))
	}

}

func howManyChickenAreRescue(tc map[string][]int) int {
	// if the roof are zero of less then no chicken will be survived
	if tc[constant.FirstLine][1] <= 0 {
		return 0
	}

	// clone original slice to new one for easier to manipulate slices
	// without messing with old one
	positionAxis := slices.Clone(tc[constant.SecondLine])

	// init result of max chicken that can be saved
	mostChickenSaveNumber := 0

	// define loop name for break case
outerLoop:
	for modIndex, modValue := range positionAxis {
		// itself position as starting point of roof and minus one for counting itself as a start point
		modValue = modValue + tc[constant.FirstLine][1] - 1
		temp := 0

		// inner loop for iterate all value for compare
		for _, orgValue := range tc[constant.SecondLine][modIndex+1:] {

			// if this true mean roof can cover that chicken
			if modValue >= orgValue {
				temp++
			}

			// if this occur -> we can cover all chicken, no need to be cal for the rest
			if temp == tc[constant.FirstLine][0]-1 {
				mostChickenSaveNumber = temp + 1
				break outerLoop
			}
		}
		// if the temp is more than previous starting point then save the new max into result
		if temp > mostChickenSaveNumber {
			mostChickenSaveNumber = temp + 1
		}

		// if the left chicken is less than maximum number, that's mean that the maximum that we can save
		// no need to cal the rest
		if mostChickenSaveNumber >= (len(tc[constant.SecondLine]) - modIndex) {
			break outerLoop
		}
	}
	return mostChickenSaveNumber
}

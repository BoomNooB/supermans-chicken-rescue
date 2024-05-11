package main

import (
	"fmt"
	"p2/constant"
	"slices"
)

func main() {
	testCases := []map[string][]int{
		// constant.TC1,
		// constant.TC2,
		// constant.TC3,
		// constant.TC4,
		// constant.TC5,
		constant.TC6,
	}

	for _, testCase := range testCases {
		fmt.Println(howManyChickenAreRescue(testCase))
	}

}

func howManyChickenAreRescue(tc map[string][]int) int {
	// clone axis slice
	if tc[constant.FirstLine][1] <= 0 {
		// if the roof are zero of less then no chicken will be survived
		return 0
	}
	// clone original slice to new one for easier to manipulate slices
	// without messing with old one
	positionAxis := slices.Clone(tc[constant.SecondLine])

	// init result of max chicken that can be saved
	mostChickenSaveNumber := 0

outerLoop:
	for modIndex, modValue := range positionAxis {
		modValue = modValue + tc[constant.FirstLine][1] - 1
		temp := 0
		for _, orgValue := range tc[constant.SecondLine][modIndex+1:] {
			if modValue >= orgValue {
				temp++
			}
			if temp == tc[constant.FirstLine][0]-1 {
				mostChickenSaveNumber = temp + 1
				break outerLoop
			}
		}
		if temp > mostChickenSaveNumber {
			mostChickenSaveNumber = temp + 1
		}
		if mostChickenSaveNumber >= (len(tc[constant.SecondLine]) - modIndex) {
			break outerLoop
		}
	}
	return mostChickenSaveNumber
}

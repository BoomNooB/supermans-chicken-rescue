package main

import (
	"log"
	"p2/constant"
	"slices"
)

func main() {
	// clone axis slice
	positionAxis := slices.Clone(constant.TC1[constant.SecondLine])
	mostChickenSaveNumber := 0

outerLoop:
	for modIndex, modValue := range positionAxis {
		modValue = modValue + constant.TC1[constant.FirstLine][1] - 1
		temp := 0
		for _, orgValue := range constant.TC1[constant.SecondLine][modIndex+1:] {
			if modValue >= orgValue {
				temp++
			}
			if temp == constant.TC1[constant.FirstLine][0]-1 {
				log.Println("All chicken are saved")
				mostChickenSaveNumber = temp + 1
				break outerLoop
			}
		}
		if temp > mostChickenSaveNumber {
			mostChickenSaveNumber = temp + 1
		}
	}

	log.Println(mostChickenSaveNumber)
}

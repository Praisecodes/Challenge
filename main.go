package main

import (
	"fmt"
)

func Copy(firstArr [15]int, secondArr [15]int) {
	firstArrLength := len(firstArr)
	AddedCount := 0

	for i := 0; i < firstArrLength; i++ {
		secondArr[i] = firstArr[i]
		AddedCount++

		percent := (100 * AddedCount) / firstArrLength
		fmt.Println(percent)
	}
}

func main() {
	var firstArray = [15]int{34, 89, 39, 58, 40, 58, 93, 34, 58, 40, 58, 39, 58, 35, 58}
	var secondArray [15]int

	Copy(firstArray, secondArray)
}

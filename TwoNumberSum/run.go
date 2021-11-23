package main

import "fmt"

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	fmt.Printf("%v", TwoNumberSum(array, targetSum))
}

func TwoNumberSum(array []int, target int) [2]int {
	var result [2]int
	for index1 := range array {
		for index2 := index1 + 1; index2 < len(array); index2++ {
			if (array[index2] + array[index1]) == target {
				result[0] = array[index1]
				result[1] = array[index2]
				return result
			}
		}
	}
	return result
}

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}

func main() {
	var result = twoSum([]int{2, 7, 11, 15}, 9)

	fmt.Println(result)
}

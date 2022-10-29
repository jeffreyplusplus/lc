package main

import "fmt"

func maxSubArray(nums []int) int {
	curSum, maxSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		curSum = max(curSum+nums[i], nums[i])
		maxSum = max(curSum, maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var result = maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

	fmt.Println(result)
}

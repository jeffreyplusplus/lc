package main

import "fmt"

func containsDuplicate(nums []int) bool {

	m := make(map[int]bool)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		} else {
			m[n] = true
		}
	}

	return false
}

func main() {
	var result = containsDuplicate([]int{1, 2, 3, 1})

	fmt.Println(result)
}

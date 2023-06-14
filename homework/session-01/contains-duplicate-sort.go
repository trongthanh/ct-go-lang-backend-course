package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	for i := 0; i < len(sorted) - 1; i++ {
		if sorted[i] == sorted[i+1] {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(nums, containsDuplicate(nums))
	nums = nums[:3]
	fmt.Println(nums, containsDuplicate(nums))
	nums = append(nums, 2)
	fmt.Println(nums, containsDuplicate(nums))
}



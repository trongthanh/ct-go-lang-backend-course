package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			return true
		}
		numMap[nums[i]] = true
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



package array

import (
	"sort"
)

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
func smallerNumbersThanCurrent(nums []int) []int {

	count := make(map[int]int)
	copy := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		copy[i] = nums[i]
	}

	sort.Slice(copy, func(i, j int) bool {
		return copy[i] < copy[j]
	})

	for i := 0; i < len(nums); i++ {
		_, ok := count[copy[i]]
		if !ok {
			count[copy[i]] = i
		}
	}

	for i := 0; i < len(nums); i++ {
		val, ok := count[nums[i]]
		if ok {
			copy[i] = val
		}
	}

	return copy
}

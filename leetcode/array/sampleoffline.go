package array

import (
	"math/rand"
	"time"
)

func sampleOfflineData(nums []int, k int) []int {

	if k >= len(nums)-1 {
		return nums
	}

	for i := 0; i < k; i++ {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(nums) - i)
		nums[len(nums)-1-i], nums[index] = nums[index], nums[len(nums)-1-i]
	}

	return nums[len(nums)-k:]
}

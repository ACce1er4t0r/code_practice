package threeSumClosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	diff, result := math.MaxInt32, 0
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp < target {
				left++
				if diff > target-tmp {
					diff = target - tmp
					result = tmp
				}
			} else if tmp > target {
				right--
				if diff > tmp-target {
					diff = tmp - target
					result = tmp
				}
			} else {
				return tmp
			}
		}
	}
	return result
}

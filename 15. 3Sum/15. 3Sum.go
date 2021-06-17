package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	//if don't have three elements, means that we can't find three nums to append into result
	if len(nums) < 3 {
		return result
	}
	//Sorted from largest to smallest
	sort.Ints(nums)
	//If the first position of the sorted array is greater than 0, it is not possible to add up to 0
	if nums[0] > 0 {
		return result
	}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for nums[left] == nums[left+1] && left+1 < right {
					left++
				}
				for nums[right] == nums[right-1] && left < right {
					right--
				}
				left++
				right--
			} else if tmp > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

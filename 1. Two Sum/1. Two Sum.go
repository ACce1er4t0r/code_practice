package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		diff := target - nums[i]
		if j, ok := m[diff]; ok {
			return []int{j, i}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

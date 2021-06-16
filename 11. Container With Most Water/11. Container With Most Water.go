package maxArea

func maxArea(height []int) int {
	left, right, area := 0, len(height)-1, 0
	for left < right {
		width := right - left
		high := 0
		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}
		tmp := width * high
		if tmp > area {
			area = tmp
		}
	}
	return area
}

package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	loc := [256]int{}
	for i := range loc {
		loc[i] = -1
	}
	length, left := 0, 0
	for i := 0; i < len(s); i++ {
		if loc[s[i]] >= left {
			left = loc[s[i]] + 1
		} else if i+1-left > length {
			length = i + 1 - left
		}
		loc[s[i]] = i
	}
	return length
}

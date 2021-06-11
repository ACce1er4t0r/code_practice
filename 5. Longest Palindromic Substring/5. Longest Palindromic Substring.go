package lengthOfLongestSubstring

// first
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}
func maxPalindrome(s string, i int, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(sub) > len(res) {
		return sub
	}
	return res
}

// second
func longestPalindrome2(s string) string {
	if len(s) <= 0 {
		return ""
	}
	left, right, l, r := 0, -1, 0, 0
	for left < len(s) {
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > r-l {
			l, r = left, right
		}
		left = (left+right)/2 + 1
		right = left
	}
	return s[l : r+1]
}

// third
//TODO manachar... ummm, next time, next time...
func longestPalindrome3(s string) string {

}

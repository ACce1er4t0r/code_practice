package isPalindrome

import (
	"math"
	"strconv"
)

// int convert to string
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

// without convert to string
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x%10 == 0 {
		return false
	} else {
		bits := 1
		tmp := x
		for tmp >= 10 {
			tmp /= 10
			bits++
		}
		i, j := 1, bits
		for i < j {
			left := (x / int(math.Pow10(i))) % 10
			right := (x / int(math.Pow10(j))) % 10
			if left != right {
				return false
			}
			i++
			j--
		}
	}
	return true
}

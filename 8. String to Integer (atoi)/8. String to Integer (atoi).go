package myAtoi

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	sign, tmp, res := 1, "", 0
	s = strings.TrimSpace(s)
	if s == "" {
		tmp = ""
	} else {
		switch s[0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sign, tmp = 1, s
		case '+':
			sign, tmp = 1, s[1:]
		case '-':
			sign, tmp = -1, s[1:]
		default:
			tmp = ""
		}
		for i, j := range tmp {
			if j < '0' || j > '9' {
				tmp = tmp[:i]
				break
			}
		}
	}
	for _, n := range tmp {
		res = res*10 + int(n-'0')
		if sign == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && sign*res < math.MinInt32 {
			return math.MinInt32
		}
	}
	return res * sign
}

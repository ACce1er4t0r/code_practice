package intToRoman

//Normal solution
func intToRoman(num int) string {
	result, i := "", 0
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for num > 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		result += symbals[i]
	}
	return result
}

//another solution
func intToRoman2(num int) string {
	// cuz num < 3999, so just set to MMMCMXCIX
	romanMap := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
	digits, tmp := 0, num
	result := ""
	for tmp > 0 {
		remain := tmp % 10
		result = romanMap[digits][remain] + result
		tmp = tmp / 10
		digits++
	}
	return result
}

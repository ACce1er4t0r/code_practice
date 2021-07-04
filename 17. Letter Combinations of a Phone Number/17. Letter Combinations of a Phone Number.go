package letterCombinations

var letterMap = []string{
	" ",    //0
	"",     //1
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	result := []string{""}

	for i := 0; i < len(digits); i++ {
		tmp := []string{}
		for j := 0; j < len(result); j++ {
			for k := 0; k < len(letterMap[digits[i]-'0']); k++ {
				tmp = append(tmp, result[j]+string(letterMap[digits[i]-'0'][k]))
			}
		}
		result = tmp
	}
	return result
}

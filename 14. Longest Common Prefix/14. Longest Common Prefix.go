package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	tmp := strs[0]
	for _, s := range strs {
		if len(tmp) > len(s) {
			tmp = s
		}
	}
	for k, v := range tmp {
		for j := 0; j < len(strs); j++ {
			if strs[j][k] != byte(v) {
				return strs[j][:k]
			}
		}
	}
	return tmp
}

package isMatch

//Using recursive matching
func isMatch(s string, p string) bool {
	i, j := len(s), len(p)
	if i == 0 && j == 0 {
		return true
	} else if i != 0 && j == 0 {
		return false
	} else if i == 0 && j != 0 {
		if p[j-1] == '*' {
			return isMatch(s, p[:j-2])
		}
		return false
	} else if s[i-1] == p[j-1] || p[j-1] == '.' {
		return isMatch(s[:i-1], p[:j-1])
	} else if p[j-1] == '*' {
		if isMatch(s, p[:j-2]) {
			return true
		}
		if s[i-1] == p[j-2] || p[j-2] == '.' {
			return isMatch(s[:i-1], p)
		}
		return false
	}
	return false
}

//TODO I don't quite understand the dynamic programming approach yet, but I'm trying to understand...
func isMatch2(s, p string) bool {
	sSize := len(s)
	pSize := len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	/* dp[i][j] represents whether s[:i] can be matched with p[:j] */

	dp[0][0] = true
	/**
	 * According to the setting of the title, "" can be matched with "a*b*c*"
	 * So, the corresponding dp needs to be set to true
	 */
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == '.' || p[j] == s[i] {
				/* p[j] and s[i] can be matched, so as long as the previous matches, it will match here */
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				/* At this point, the match of p[j] is related to the content of p[j-1] */
				if p[j-1] != s[i] && p[j-1] != '.' {
					/**
					 * p[j] cannot be matched with s[i]
					 * p[j-1:j+1] can only be taken as ""
					 */
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					/**
					 * p[j] is matched with s[i]
					 * p[j-1;j+1] as "x*", which can be interpreted in three ways
					 */
					dp[i+1][j+1] = dp[i+1][j-1] || /* "x*" is interpreted as "" */
						dp[i+1][j] || /* "x*" is interpreted as "x" */
						dp[i][j+1] /* "x*" is interpreted as "xx..." */
				}
			}
		}
	}

	return dp[sSize][pSize]
}

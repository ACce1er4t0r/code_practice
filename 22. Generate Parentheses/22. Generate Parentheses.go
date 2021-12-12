package generateParenthesis

func generateParenthesis1(n int) []string {
    dp := make([][]string ,n + 1)
	dp[0] = make([]string, 0)
	if n == 0{
		return dp[0]
	}
	dp[0] = append(dp[0], "")
	
	for i := 1; i <= n; i++{
		dp[i] = make([]string, 0)
		for j := 0; j < i; j++{
			for _, s1 := range dp[j]{
				for _, s2 := range dp[i-j-1]{
					s := "(" + s1 + ")" + s2
					dp[i] = append(dp[i], s)
				}

			}
		}
	}
    return dp[n]
}

func generateParenthesis2(n int) []string {
    ret := make([]string, 0)
    if n == 0 {
        return ret
    }
    ret = generateParenthesisDfs("", n, n, ret)
    
    return ret
}

func generateParenthesisDfs(s string, left int, right int, ret []string) []string {
    if left == 0 && right == 0{
		ret = append(ret, s)
		return ret
	}
	if left > 0{
		ret = generateParenthesisDfs(s + "(", left - 1, right, ret)
	}

	if right > 0 && left < right{
		ret = generateParenthesisDfs(s + ")", left, right - 1, ret)
	}
	return ret
}

func generateParenthesis3(n int) []string {
    m := make([]string, 0)
    ret := make([]string, 0)
    if n == 0 {
        return m
    }
    if n == 1 {
        return append(ret, "()")
    }
    
    tmp1 := generateParenthesis(n-1)
    for _,s := range(tmp1) {
        for i,_ := range(s) {
            tmp2 := s[:i] + "()" + s[i:]
            m = append(m, tmp2)
        }
    }
    sort.Strings(m)
    l := len(m)
    for i:=0; i < l; i++{
        if (i > 0 && m[i-1] == m[i]) || len(m[i])==0 {
            continue;
        }
        ret = append(ret, m[i])
    }
    return ret
}
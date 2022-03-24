package problems

func longestCommonPrefix(strs []string) string {
	lcp := ""
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if i >= len(strs[j+1]) || i >= len(strs[j]) || strs[j][i] != strs[j+1][i] {
				return lcp
			}
		}
		lcp = strs[0][:i+1]
	}
	return lcp
}

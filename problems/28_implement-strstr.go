package problems

func strStr(haystack string, needle string) int {

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}

	}
	return -1
}

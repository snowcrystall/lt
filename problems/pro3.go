package problems

func LengthOfLongestSubstring(s string) int {
	var i, j, t, sublen, maxlen int = 0, 1, 0, 0, 0
	len := len(s)
	if len <= 1 {
		return len
	}
	for j < len {
		t = i
		sublen = 1
		for t < j {
			if s[t] == s[j] {
				i = t + 1
				break
			} else {
				sublen++
				t++
			}
		}
		j++
		if sublen > maxlen {
			maxlen = sublen
		}
	}
	return maxlen
}

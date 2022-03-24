package problems

func romanToInt(s string) int {
	res := 0
	numMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s); i++ {

		if int(i+1) < len(s) && numMap[s[i+1]] > numMap[s[i]] {

			res -= numMap[s[i]]
		} else {
			res += numMap[s[i]]
		}

	}
	return res
}

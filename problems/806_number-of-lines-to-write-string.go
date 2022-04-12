package problems

func numberOfLines(widths []int, s string) []int {
	res := 0
	line := 0

	for _, v := range s {
		if res+widths[v-97] > 100 {
			res = 0
			line++
		}
		res += widths[v-97]
	}
	if res != 0 {
		line++
	}
	return []int{line, res}
}

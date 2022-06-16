package medium

func findPairs(nums []int, k int) int {

	res := map[int]byte{}
	for _, num := range nums {
		if _, ok := res[num-k]; ok {
			res[num-k] = 1
		}
		if _, ok := res[num+k]; ok {
			res[num] = 1
		} else {
			res[num] = 0
		}
	}
	sum := 0
	for _, v := range res {
		sum += int(v)
	}
	//fmt.Println(res)
	return sum

}

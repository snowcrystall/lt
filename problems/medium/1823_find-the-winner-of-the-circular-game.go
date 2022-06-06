package medium

func findTheWinner(n int, k int) int {
	nlist := []int{}
	for i := 1; i <= n; i++ {
		nlist = append(nlist, i)

	}
	//fmt.Println(nlist)
	pos := 0
	for len(nlist) > 1 {
		pos = (k+pos)%(len(nlist)) - 1
		if pos < 0 {
			pos = pos + len(nlist)
		}
		nlist = append(nlist[0:pos], nlist[pos+1:]...)
		//fmt.Println(pos,nlist)
	}
	return nlist[0]
}

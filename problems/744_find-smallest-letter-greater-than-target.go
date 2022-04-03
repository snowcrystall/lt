package problems

func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)-1
	mid := 0
	for lo <= hi {
		mid = (lo + hi) / 2
		if letters[mid] > target {
			hi = mid
			if lo == hi {
				break
			}
		} else {
			lo = mid + 1
		}
	}
	return letters[lo%len(letters)]
}

package problems

func removeDuplicates(nums []int) int {
	len := len(nums)
	if len == 0 {
		return len
	}
	i, j := 0, 1
	for j < len {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		} else {
			j++
		}
	}

	return i + 1
}

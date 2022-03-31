package problems

func removeElement(nums []int, val int) int {
	len := len(nums)
	if len == 0 {
		return len
	}
	slow := 0
	for fast := 0; fast < len; fast++ {
		if nums[slow] == val && nums[fast] != val {
			nums[slow] = nums[fast]
			nums[fast] = val
			slow++
		} else if nums[slow] == val {

		} else {
			slow++
		}
	}
	return slow
}

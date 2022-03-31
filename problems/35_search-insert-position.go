package problems

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	mid := 0
	for lo < hi {
		mid = (lo + hi) / 2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid
		} else {
			return mid
		}
	}
	return hi
}

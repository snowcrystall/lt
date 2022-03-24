package problems

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	var tmp, i, num, j int
	var ok bool
	for i, num = range nums {
		tmp = target - num
		if j, ok = numsMap[tmp]; ok && i != j {
			return []int{i, j}
		}
		numsMap[num] = i
	}
	return nil
}

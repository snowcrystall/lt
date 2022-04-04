/*给你一个数组 nums ，请你完成两类查询。

其中一类查询要求 更新 数组 nums 下标对应的值
另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
实现 NumArray 类：

NumArray(int[] nums) 用整数数组 nums 初始化对象
void update(int index, int val) 将 nums[index] 的值 更新 为 val
int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）
*/
/*
 *TODO SumRange不是O(1) 空间占用也很大，优化思路： LRU？
 */
package problems

import (
	"log"
)

type NumArray struct {
	arr    []int
	sumMap map[int]int // sum for arr[index] ~ arr[index+chunk-1] of arr
	chunk  int
}

func ConstructorNumArray(nums []int) NumArray {
	numarr := NumArray{[]int{}, map[int]int{}, 1000}
	for i, v := range nums {
		numarr.arr = append(numarr.arr, v)

		numarr.sumMap[i/numarr.chunk] = numarr.sumMap[i/numarr.chunk] + v

	}

	return numarr
}

func (this *NumArray) Update(index int, val int) {
	change := val - this.arr[index]
	this.arr[index] = val
	this.sumMap[index/this.chunk] = this.sumMap[index/this.chunk] + change

}

func (this *NumArray) SumRange(left int, right int) int {
	log.Println(this.arr, left, right, this.sumMap)
	sum := 0
	if left == right {
		return this.arr[left]
	}
	if left > right {
		left, right = right, left
	}
	if left/this.chunk == right/this.chunk {

		for i := left; i < right+1; i++ {
			sum += this.arr[i]
		}
		return sum
	}
	leftpart := this.sumMap[left/this.chunk]
	mod := left % this.chunk
	for mod > 0 {
		leftpart -= this.arr[left-mod]
		mod--
	}
	rightpart := this.sumMap[right/this.chunk]
	mod = this.chunk - right%this.chunk - 1
	for i := 1; mod > 0 && right+i < len(this.arr); i++ {
		rightpart -= this.arr[right+i]
		mod--

	}
	midpart := 0
	for i := left/this.chunk + 1; i < right/this.chunk; i++ {
		midpart += this.sumMap[i]
	}

	return leftpart + rightpart + midpart
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

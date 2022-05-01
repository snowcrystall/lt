package problems

/* https://leetcode-cn.com/problems/number-of-1-bits/
 * 191. 位1的个数
 * 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
 */
func hammingWeight(num uint32) int {
	one := 0
	for ; num > 0; num &= num - 1 {
		one++
	}
	return one
}

/*
 * 868. Binary Gap
 * Given a positive integer n, find and return the longest distance between any two adjacent 1's in the binary representation of n. If there are no two adjacent 1's, return 0.

Two 1's are adjacent if there are only 0's separating them (possibly no 0's). The distance between two 1's is the absolute difference between their bit positions. For example, the two 1's in "1001" have a distance of 3
 https://leetcode-cn.com/problems/binary-gap/
*/

func binaryGap(n int) int {
	max := 0
	dis := 0
	for n > 0 {
		if n&1 == 1 {

			if max < dis {
				max = dis
			}

			dis = 1

		} else if dis != 0 {
			dis++
		}
		n = n >> 1
	}
	return max
}

//118. Pascal's Triangle
func generate(numRows int) [][]int {
	res := [][]int{}

	for i := 0; i < numRows; i++ {
		tmp := []int{}
		if i == 0 {
			tmp = append(tmp, 1)
		} else {
			tmp = append(tmp, 1)
			for j := 1; j < i; j++ {
				tmp = append(tmp, res[i-1][j-1]+res[i-1][j])
			}
			tmp = append(tmp, 1)
		}
		res = append(res, tmp)
	}
	return res
}

//119. https://leetcode-cn.com/problems/pascals-triangle-ii/
func getRow(rowIndex int) []int {
	tmp1 := []int{}
	tmp2 := []int{}
	for i := 0; i < rowIndex+1; i++ {
		if i == 0 {
			tmp1 = append(tmp1, 1)
		} else {
			tmp2 = append(tmp2, 1)
			for j := 1; j < i; j++ {
				tmp2 = append(tmp2, tmp1[j-1]+tmp1[j])
			}
			tmp2 = append(tmp2, 1)

			tmp1, tmp2 = tmp2, tmp1

			tmp2 = tmp2[:0]
		}
	}
	return tmp1
}

//121. 买卖股票的最佳时机 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	maxsub := 0
	min := prices[0]
	for _, v := range prices {
		if v < min {
			min = v
		} else {
			if v-min > maxsub {
				maxsub = v - min
			}
		}
	}
	return maxsub
}

//136. 只出现一次的数字
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

//169. 多数元素
func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}

	}

	return candidate

}

//168. https://leetcode-cn.com/problems/excel-sheet-column-title/
func convertToTitle(columnNumber int) string {

	ans := []byte{}
	for columnNumber > 0 {
		a0 := (columnNumber-1)%26 + 1
		ans = append(ans, 'A'+byte(a0-1))
		columnNumber = (columnNumber - a0) / 26
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)

}

//171. https://leetcode-cn.com/problems/excel-sheet-column-number/
func titleToNumber(columnTitle string) int {
	sum := 0
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		sum += int(k) * multiple
		multiple *= 26
	}
	return sum
}

//190. https://leetcode-cn.com/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; i < 32 && num > 0; i++ {
		rev |= num & 1 << (31 - i)
		num >>= 1
	}
	return rev
}

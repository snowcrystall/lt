package jianzhiofferii

import "strconv"

// 此方法超长字符串溢出
func addBinary(a string, b string) string {

	var x, y, res, carry uint64
	x, _ = strconv.ParseUint(a, 2, 64)
	y, _ = strconv.ParseUint(b, 2, 64)

	for y > 0 {
		res = (x ^ y)
		carry = (x & y) << 1
		x, y = res, carry
	}
	return strconv.FormatUint(x, 2)

}

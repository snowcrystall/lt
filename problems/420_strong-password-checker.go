package problems

import (
	"fmt"
	"unicode"
)

func StrongPasswordChecker(password string) int {
	res := 0
	if len(password) < 6 {
		addNum := 6 - len(password)
		meanNum := meanN(password)
		res = max(addNum, meanNum)

	} else if len(password) > 20 {
		smap := sameNumMap(password)

		delNum := len(password) - 20
		delNum2 := delNum
		chanNum := smap.changeN()
		meanNum := meanN(password)
		fmt.Println(delNum, chanNum, meanNum)
		for delNum2 > 0 && chanNum > 0 {
		first:
			for i, _ := range smap {
				if delNum2 > 0 && chanNum > 0 && smap[i]%3 == 0 {
					delNum2--
					chanNum--
					smap[i]--
				}
				if smap[i] < 3 {
					delete(smap, i)
				}
			}

			for i, _ := range smap {
				if delNum2 > 0 && smap[i]%3 == 1 {
					delNum2--
					smap[i]--
					goto first
				}
			}
			for i, _ := range smap {
				if delNum2 > 0 {
					delNum2--
					smap[i]--
					goto first
				}
			}

			fmt.Println(delNum2, chanNum, smap)
		}

		fmt.Println(delNum, chanNum, meanNum, password, len(password))
		res = delNum + max(chanNum, meanNum)
	} else {
		smap := sameNumMap(password)
		chanNum := smap.changeN()

		meanNum := meanN(password)
		res = max(chanNum, meanNum)
	}

	return res

}

type sameMap map[int]int

func sameNumMap(password string) sameMap {
	sameNumMap := map[int]int{}

	j := 1
	for i := 1; i < len(password); i = j {
		for ; j < len(password); j++ {
			if password[j-1] != password[j] {
				break
			}
		}

		sameNumMap[i] = (j - i + 1)
		j++
	}
	return sameNumMap
}

func (smap sameMap) changeN() int {

	changeNum := 0
	for _, v := range smap {
		changeNum += v / 3
	}
	return changeNum
}

func meanN(password string) int {
	hasLower, hasUpper, hasDigit, meanNum := 0, 0, 0, 0
	for _, v := range password {
		if unicode.IsLower(v) {
			hasLower = 1
		} else if unicode.IsLetter(v) {
			hasUpper = 1
		} else if unicode.IsDigit(v) {
			hasDigit = 1
		}
	}
	meanNum = hasDigit + hasLower + hasUpper
	return 3 - meanNum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

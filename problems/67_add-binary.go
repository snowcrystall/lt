package problems

func addBinary(a string, b string) string {
	sh, lo := []byte{}, []byte{}
	res := []byte{}
	if len(a) > len(b) {
		lo = []byte(a)
		sh = []byte(b)
		res = []byte(a)
	} else {
		lo = []byte(b)
		sh = []byte(a)
		res = []byte(b)
	}

	var carry byte = 0
	i := len(lo)
	j := len(sh)
	for j > 0 {
		i--
		j--
		res[i] = lo[i] + sh[j] + carry - 48
		if res[i] > 49 {
			res[i] -= 2
			carry = 1
		} else {
			carry = 0
		}

	}
	for i > 0 {
		i--
		res[i] = lo[i] + carry
		if res[i] > 49 {
			res[i] -= 2
			carry = 1
		} else {
			carry = 0
		}
	}

	if carry == 1 {
		res = append([]byte{'1'}, res...)
	}
	return string(res)
}

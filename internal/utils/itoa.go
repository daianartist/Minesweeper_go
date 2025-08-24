package utils

func Itoa(n int) string {
	res := []byte{}
	slice := [20]byte{}
	i := 0
	if n == 0 {
		return "0"
	}

	if n < 0 {
		res = append(res, '-')
		n = -n
	}

	for n > 0 {
		digit := n % 10
		slice[i] = byte('0' + digit)
		n /= 10
		i++
	}

	for j := i - 1; j >= 0; j-- {
		res = append(res, slice[j])
	}

	return string(res)
}

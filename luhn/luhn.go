package luhn

import "strings"

func toInt(c rune) int {
	i := int(c - '0')
	if i < 0 || i > 9 {
		return -1
	}
	return i
}

func singleDegitDouble(n int) int {
	n = n * 2
	if n < 9 {
		return n
	}
	return n - 9
}

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	l := len(id)
	if l <= 1 {
		return false
	}
	sum := 0
	for i, c := range id {
		n := toInt(c)
		if n == -1 {
			return false
		}
		if (l-i)%2 == 1 {
			sum += n
		} else {
			sum += singleDegitDouble(n)
		}
	}
	return (sum%10 == 0)
}

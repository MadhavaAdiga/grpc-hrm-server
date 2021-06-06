package utils

func NumberLen(n int) int {
	var count int = 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

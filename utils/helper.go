package utils

import "net/mail"

func NumberLen(n int) int {
	var count int = 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func ValidateMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

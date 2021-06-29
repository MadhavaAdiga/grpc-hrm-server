package utils

import (
	"net/mail"

	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
)

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

// check for a valid permission
// return a bool
func CheckPermission(permissionType hrm.Permission, permissionList []int32) bool {
	for _, v := range permissionList {
		if v == int32(permissionType.Number()) {
			return true
		}
	}
	return false
}

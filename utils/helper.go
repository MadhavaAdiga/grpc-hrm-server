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
func CheckPermission(permissionType hrm.Permission, permissionList []int32) bool {
	for _, v := range permissionList {
		if v == int32(permissionType.Number()) {
			return true
		}
	}
	return false
}

// check if a valid permission exists in a list of given permission
func CheckPermissions(validPermissions []hrm.Permission, permissionList []int32) bool {
	var count int16 = 0

	m := make(map[int32]string)
	// create a map for validating
	for _, v := range validPermissions {
		m[int32(v.Number())] = v.String()
	}

	for _, v := range permissionList {
		if _, ok := m[v]; ok {
			count++
		}
	}

	return count > 0
}

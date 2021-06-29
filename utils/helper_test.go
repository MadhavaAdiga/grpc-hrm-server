package utils_test

import (
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/stretchr/testify/require"
)

func TestNumberLen(t *testing.T) {
	t.Parallel()

	count := utils.NumberLen(1234567890)
	require.Equal(t, 10, count)
	count = utils.NumberLen(12347890)
	require.Equal(t, 8, count)
	count = utils.NumberLen(1234567891234)
	require.Equal(t, 13, count)
}

func TestValidateMail(t *testing.T) {
	s := "asdas"
	isemail := utils.ValidateMail(s)
	require.False(t, isemail)

	valid := "a@example.com"
	isemail = utils.ValidateMail(valid)
	require.True(t, isemail)
}

func TestCheckPermission(t *testing.T) {
	permissions1 := []int32{0, 3, 2, 6}
	permissionType := hrm.Permission_ADMIN

	v := utils.CheckPermission(permissionType, permissions1)
	require.True(t, v)

	permissions2 := []int32{0, 3, 2, 6}
	permissionType2 := hrm.Permission_CAN_VIEW_SALARIES

	v = utils.CheckPermission(permissionType2, permissions2)
	require.False(t, v)
}

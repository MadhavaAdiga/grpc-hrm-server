package utils_test

import (
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/stretchr/testify/require"
)

func TestValidatePassword(t *testing.T) {
	t.Parallel()

	password := utils.RandomString(8)

	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = utils.ValidatePassword(hashedPassword, password)
	require.NoError(t, err)
}

package auth_test

import (
	"testing"
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/internal/auth"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetomanager(t *testing.T) {
	t.Parallel()

	manager, err := auth.NewPasetoManager()
	require.NoError(t, err)
	require.NotNil(t, manager)

	userName := utils.RandomName()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := manager.CreateToken(userName, duration)
	require.NoError(t, err)
	require.NotNil(t, token)

	payload, err := manager.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, payload.UserName, userName)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpireAt, time.Second)
}

func TestJwtManager(t *testing.T) {
	t.Parallel()

	manager, err := auth.NewPasetoManager()
	require.NoError(t, err)
	require.NotNil(t, manager)

	_, err = manager.CreateToken("asd", time.Duration(1))
	require.NoError(t, err)
}

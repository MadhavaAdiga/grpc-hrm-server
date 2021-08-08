package auth_test

import (
	"testing"
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/internal/auth"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPasetomanager(t *testing.T) {
	t.Parallel()

	manager, err := auth.NewPasetoManager()
	require.NoError(t, err)
	require.NotNil(t, manager)

	shortUID := uuid.New()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	permissions := []hrm.Permission{hrm.Permission_ADMIN, hrm.Permission_CAN_ADD_EMPLOYEE}

	token, err := manager.CreateToken(shortUID, duration, permissions)
	require.NoError(t, err)
	require.NotNil(t, token)

	payload, err := manager.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, payload.ShortUid, shortUID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpireAt, time.Second)
}

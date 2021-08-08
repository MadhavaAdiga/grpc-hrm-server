package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/auth"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSetPrinciple(t *testing.T) {
	t.Parallel()

	manager := auth.NewRefreshTokenManager("secret", time.Minute)
	require.NotNil(t, manager)
	id := uuid.New()

	refreshToken, err := manager.Generate()
	require.NoError(t, err)
	require.NotNil(t, refreshToken)

	principle := db.Priniple{
		UserName:     utils.RandomName(),
		RefreshToken: refreshToken,
	}

	err1 := testCacheStore.SetPrinciple(context.Background(), id.String(), principle)
	require.NoError(t, err1)
}

func TestGetPrinciple(t *testing.T) {
	t.Parallel()

	manager := auth.NewRefreshTokenManager("secret", time.Minute)
	require.NotNil(t, manager)
	id := uuid.New()

	refreshToken, err := manager.Generate()
	require.NoError(t, err)
	require.NotNil(t, refreshToken)

	principle := db.Priniple{
		UserName:     utils.RandomName(),
		RefreshToken: refreshToken,
	}

	err1 := testCacheStore.SetPrinciple(context.Background(), id.String(), principle)
	require.NoError(t, err1)

	principle1, err1 := testCacheStore.GetPrinciple(context.Background(), id.String())
	require.NoError(t, err1)
	require.NotNil(t, principle1)

	require.Equal(t, principle.RefreshToken, principle1.RefreshToken)
	require.Equal(t, principle.UserName, principle1.UserName)
}

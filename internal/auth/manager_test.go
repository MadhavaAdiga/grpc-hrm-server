package auth_test

import (
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/internal/auth"
	"github.com/stretchr/testify/require"
)

func TestPasetomanager(t *testing.T) {
	t.Parallel()

	manager, err := auth.NewPasetoManager()
	require.NoError(t, err)
	require.NotNil(t, manager)
}

package utils_test

import (
	"testing"

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

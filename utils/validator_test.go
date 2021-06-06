package utils_test

import (
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/stretchr/testify/require"
)

func TestNilValidator(t *testing.T) {
	t.Parallel()

	s := struct {
		Number      int            `validate:"NotNil"`
		Float       float32        `validate:"NotNil"`
		Str         string         `validate:"NotNil"`
		Boolean     bool           `validate:"NotNil"`
		IntSlice    []int          `validate:"NotNil"`
		StringSlice []string       `validate:"NotNil"`
		FloatSlice  []float32      `validate:"NotNil"`
		BoolSlice   []bool         `validate:"NotNil"`
		Pointer     *int           `validate:"NotNil"`
		MapVar      map[int]string `validate:"NotNil"`
	}{}

	errs := utils.ValidateStruct(s)
	for _, err := range errs {
		require.Error(t, err)
	}

}

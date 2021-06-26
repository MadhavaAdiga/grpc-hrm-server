package db_test

import (
	"context"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreatePayroll(t *testing.T) {
	t.Parallel()

	createPayroll(t)
}

func createPayroll(t *testing.T) db.Payroll {
	emp := createEmployee(t)

	arg := db.CreatePayrollParam{
		Employee:  emp.ID,
		Ctc:       utils.RandomInt(0, 6),
		Allowance: utils.RandomInt(0, 3),
		CreatedBy: uuid.New(),
	}

	payroll, err := testSQLStore.CreatePayroll(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, payroll)
	require.NotEqual(t, payroll.ID, uuid.Nil)

	require.Equal(t, payroll.Ctc, arg.Ctc)
	require.Equal(t, payroll.Allowance, arg.Allowance)
	require.Equal(t, payroll.CreateBy, arg.CreatedBy)

	require.NotZero(t, emp.CreatedAt)

	return payroll
}

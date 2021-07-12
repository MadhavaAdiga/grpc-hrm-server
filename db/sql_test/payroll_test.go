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

func TestFindPayRollByID(t *testing.T) {
	t.Parallel()

	payroll := createPayroll(t)

	payroll1, err := testSQLStore.FindPayrollByEmpID(context.Background(), payroll.Employee.ID)
	require.NoError(t, err)

	require.Equal(t, payroll.ID, payroll1.ID)
	require.Equal(t, payroll.Ctc, payroll1.Ctc)
	require.Equal(t, payroll.Employee.ID, payroll1.Employee.ID)
	require.Equal(t, payroll.Allowance, payroll1.Allowance)
	require.Equal(t, payroll.CreateBy, payroll1.CreateBy)
	require.Equal(t, payroll.CreatedAt, payroll1.CreatedAt)

	require.Equal(t, payroll.Employee.ID, payroll1.Employee.ID)
}

func TestFindPayRollByName(t *testing.T) {
	t.Parallel()

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

	user, err := testSQLStore.FindUserById(context.Background(), emp.User.ID)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.NotEqual(t, user.ID, uuid.Nil)

	payroll1, err := testSQLStore.FindPayrollByEmpName(context.Background(), user.UserName)
	require.NoError(t, err)

	require.Equal(t, payroll.ID, payroll1.ID)
	require.Equal(t, payroll.Ctc, payroll1.Ctc)
	require.Equal(t, payroll.Employee.ID, payroll1.Employee.ID)
	require.Equal(t, payroll.Allowance, payroll1.Allowance)
	require.Equal(t, payroll.CreateBy, payroll1.CreateBy)
	require.Equal(t, payroll.CreatedAt, payroll1.CreatedAt)

	require.Equal(t, payroll.Employee.ID, payroll1.Employee.ID)
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
	require.Equal(t, payroll.Employee.ID, emp.ID)
	require.Equal(t, payroll.Allowance, arg.Allowance)
	require.Equal(t, payroll.CreateBy, arg.CreatedBy)

	require.NotZero(t, emp.CreatedAt)

	return payroll
}

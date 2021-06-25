package db_test

import (
	"context"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateEmployee(t *testing.T) {
	t.Parallel()
}

func createEmployee(t *testing.T, arg db.CreateEmployeeParam) db.Employee {
	if (db.CreateEmployeeParam{}) == arg {
		org := createOrganization(t)
		usr := createUser(t, db.CreateUserParam{})
		role := createRole(t)

		arg = db.CreateEmployeeParam{
			User_id:         usr,
			Organization_id: org.ID,
			Role_id:         role.ID,
			Status:          1,
			CreatedBy:       uuid.New(),
		}
	}

	emp, err := testSQLStore.CreateEmployee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, emp.ID, uuid.Nil)

	require.Equal(t, emp.Organization.ID, arg.Organization_id)
	require.Equal(t, emp.User.ID, arg.User_id)
	require.Equal(t, emp.Role, arg.Role_id)
	require.Equal(t, emp.CreateBy, arg.CreatedBy)

	require.NotZero(t, emp.CreatedAt)

	return emp

}

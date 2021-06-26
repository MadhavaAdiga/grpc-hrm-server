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

	createEmployee(t)
}

func createEmployee(t *testing.T) db.Employee {
	org := createOrganization(t)
	usr := createUser(t)
	role := createRole(t)

	arg := db.CreateEmployeeParam{
		UserId:         usr.ID,
		OrganizationId: org.ID,
		RoleId:         role.ID,
		Status:         1,
		CreatedBy:      uuid.New(),
	}

	emp, err := testSQLStore.CreateEmployee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, emp.ID, uuid.Nil)

	require.Equal(t, emp.Organization.ID, arg.OrganizationId)
	require.Equal(t, emp.User.ID, arg.UserId)
	require.Equal(t, emp.Role.ID, arg.RoleId)
	require.Equal(t, emp.CreateBy, arg.CreatedBy)

	require.NotZero(t, emp.CreatedAt)

	return emp

}

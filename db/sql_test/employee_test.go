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

func TestFindEmpbyUNameAndOrg(t *testing.T) {
	t.Parallel()

	emp := createEmployee(t)

	usr, err := testSQLStore.FindUserById(context.Background(), emp.User.ID)
	require.NoError(t, err)

	org, err := testSQLStore.FindOrganizationByID(context.Background(), emp.Organization.ID)
	require.NoError(t, err)

	findArg := db.FindEmployeeUnameAndOrgParam{
		OrganizationName: org.Name,
		Username:         usr.UserName,
	}

	emp1, err := testSQLStore.FindEmployeeByUnameAndOrg(context.Background(), findArg)
	require.NoError(t, err)
	require.NotNil(t, emp1)

	require.Equal(t, emp.ID, emp1.ID)
	require.Equal(t, emp.User.ID, emp1.User.ID)
	require.Equal(t, usr.UserName, emp1.User.UserName)
	require.Equal(t, emp.Organization.ID, emp1.Organization.ID)
	require.Equal(t, emp.Role.ID, emp1.Role.ID)
	require.Equal(t, emp.Status, emp1.Status)
	require.Equal(t, emp.CreateBy, emp1.CreateBy)
	require.Equal(t, emp.CreatedAt, emp1.CreatedAt)
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

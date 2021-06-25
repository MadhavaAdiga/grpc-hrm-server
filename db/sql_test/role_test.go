package db_test

import (
	"context"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateRole(t *testing.T) {
	t.Parallel()

	createRole(t)
}

func TestFindRoleByNameAndOrg(t *testing.T) {
	t.Parallel()

	role := createRole(t)

	findArg := db.FindRoleByOrgIDParam{
		Name:         role.Name,
		Organization: role.Organization.ID,
	}

	role1, err := testSQLStore.FindRoleByOrganizationID(context.Background(), findArg)
	require.NoError(t, err)
	require.NotEmpty(t, role1)
	require.NotEqual(t, role1.ID, uuid.Nil)

	require.Equal(t, role.ID, role1.ID)
	require.Equal(t, role.Organization.ID, role1.Organization.ID)
	for i := range role.Permissions {
		require.Equal(t, role.Permissions[i], role1.Permissions[i])
	}
	require.Equal(t, role.CreatedBy, role1.CreatedBy)

	require.NotZero(t, role.CreatedAt)
}

func TestFindRoleByNameAndOrgName(t *testing.T) {
	t.Parallel()

	role := createRole(t)

	org, err := testSQLStore.FindOrganizationByID(context.Background(), role.Organization.ID)
	require.NoError(t, err)

	findArg := db.FindRoleByOrgNameParam{
		Name:             role.Name,
		OrganizationName: org.Name,
	}

	role1, err := testSQLStore.FindRoleByOrganizationName(context.Background(), findArg)
	require.NoError(t, err)
	require.NotEmpty(t, role1)
	require.NotEqual(t, role1.ID, uuid.Nil)

	require.Equal(t, role.ID, role1.ID)
	require.Equal(t, role.Organization.ID, role1.Organization.ID)
	for i := range role.Permissions {
		require.Equal(t, role.Permissions[i], role1.Permissions[i])
	}
	require.Equal(t, role.CreatedBy, role1.CreatedBy)

	require.NotZero(t, role.CreatedAt)
}

func createRole(t *testing.T) db.Role {
	org := createOrganization(t)

	arg := db.CreateRoleParam{
		Name:         utils.RandomName(),
		Active:       true,
		Organization: org.ID,
		Permissions: []int32{
			int32(hrm.Permission_ADMIN),
		},
		CreatedBy: uuid.New(),
	}

	role, err := testSQLStore.CreateRole(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotEqual(t, role.ID, uuid.Nil)

	require.Equal(t, role.Organization.ID, org.ID)
	for i := range role.Permissions {
		require.Equal(t, role.Permissions[i], arg.Permissions[i])
	}
	require.Equal(t, role.CreatedBy, arg.CreatedBy)

	require.NotZero(t, role.CreatedAt)

	return role
}

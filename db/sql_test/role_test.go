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

	createRole(t, arg)
}

func TestFindRoleByNameAndOrg(t *testing.T) {
	t.Parallel()
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

	id := createRole(t, arg)

	findArg := db.FindRoleByOrgIDParam{
		Name:         arg.Name,
		Organization: org.ID,
	}

	role, err := testSQLStore.FindRoleByOrganizationID(context.Background(), findArg)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotEqual(t, role.ID, uuid.Nil)

	require.Equal(t, role.ID, id)
	require.Equal(t, role.Organization.ID, org.ID)
	for i := range role.Permissions {
		require.Equal(t, role.Permissions[i], arg.Permissions[i])
	}
	require.Equal(t, role.CreatedBy, arg.CreatedBy)

	require.NotZero(t, role.CreatedAt)
}

func TestFindRoleByNameAndOrgName(t *testing.T) {
	t.Parallel()
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

	id := createRole(t, arg)

	findArg := db.FindRoleByOrgNameParam{
		Name:             arg.Name,
		OrganizationName: org.Name,
	}

	role, err := testSQLStore.FindRoleByOrganizationName(context.Background(), findArg)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotEqual(t, role.ID, uuid.Nil)

	require.Equal(t, role.ID, id)
	require.Equal(t, role.Organization.ID, org.ID)
	for i := range role.Permissions {
		require.Equal(t, role.Permissions[i], arg.Permissions[i])
	}
	require.Equal(t, role.CreatedBy, arg.CreatedBy)

	require.NotZero(t, role.CreatedAt)
}

func createRole(t *testing.T, arg db.CreateRoleParam) uuid.UUID {
	id, err := testSQLStore.CreateRole(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotEqual(t, id, uuid.Nil)

	return id
}

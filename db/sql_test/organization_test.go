package db_test

import (
	"context"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateOrganization(t *testing.T) {
	t.Parallel()
	createOrganization(t)
}

func TestFindOrganization(t *testing.T) {
	t.Parallel()
	organization1 := createOrganization(t)

	organization2, err := testSQLStore.FindOrganizationByName(context.Background(), organization1.Name)
	require.NoError(t, err)
	require.NotNil(t, organization2)

	require.Equal(t, organization1.ID, organization2.ID)

	require.Equal(t, organization1.Name, organization2.Name)
	require.Equal(t, organization1.Status, organization2.Status)
	// require.Equal(t, organization1.CreatedBy, organization2.CreatedBy)
	require.Equal(t, organization1.CreatorID, organization2.CreatorID)

	require.Equal(t, organization1.CreatedAt, organization2.CreatedAt)
	require.Equal(t, organization1.UpdatedAt, organization2.UpdatedAt)
}

func TestFindOrganizationById(t *testing.T) {
	t.Parallel()
	organization1 := createOrganization(t)

	organization2, err := testSQLStore.FindOrganizationByID(context.Background(), organization1.ID)
	require.NoError(t, err)
	require.NotNil(t, organization2)

	require.Equal(t, organization1.ID, organization2.ID)

	require.Equal(t, organization1.Name, organization2.Name)
	require.Equal(t, organization1.Status, organization2.Status)
	// require.Equal(t, organization1.CreatedBy, organization2.CreatedBy)
	require.Equal(t, organization1.CreatorID, organization2.CreatorID)

	require.Equal(t, organization1.CreatedAt, organization2.CreatedAt)
	require.Equal(t, organization1.UpdatedAt, organization2.UpdatedAt)
}

func createOrganization(t *testing.T) db.Organization {
	arg := db.CreateOrganizationParam{
		Name: utils.RandomName(),
		// CreatedBy: utils.RandomName(),
		Status:    0,
		CreatorID: uuid.New(),
	}

	organization, err := testSQLStore.CreateOrganization(context.Background(), arg)
	// organization.ID = uuid.Nil
	require.NoError(t, err)
	require.NotNil(t, organization)

	require.NotEqual(t, organization.ID, uuid.Nil)

	require.Equal(t, organization.Name, arg.Name)
	require.Equal(t, organization.Status, arg.Status)
	// require.Equal(t, organization.CreatedBy, arg.CreatedBy)
	require.Equal(t, organization.CreatorID, arg.CreatorID)

	require.NotZero(t, organization.CreatedAt)

	return organization
}

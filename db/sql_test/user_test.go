package db_test

import (
	"context"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	arg := db.CreateUserParam{
		FirstName:      utils.RandomName(),
		LastName:       utils.RandomName(),
		HashedPassword: "secret",
		UserName:       utils.RandomName(),
		Address:        utils.RandomString(15),
		Email:          utils.RandomString(4),
		ContactNumber:  uint32(utils.RandomContactNum()),
	}
	createUser(t, arg)
}

func TestFindUserByName(t *testing.T) {
	t.Parallel()
	arg := db.CreateUserParam{
		FirstName:      utils.RandomName(),
		LastName:       utils.RandomName(),
		HashedPassword: "secret",
		UserName:       utils.RandomName(),
		Address:        utils.RandomString(15),
		Email:          utils.RandomString(4),
		ContactNumber:  uint32(utils.RandomContactNum()),
	}
	createUser(t, arg)

	user, err := testSQLStore.FindUserByName(context.Background(), arg.UserName)
	require.NoError(t, err)
	require.NotNil(t, user)

	require.NotEqual(t, user.ID, uuid.Nil)
	require.Equal(t, user.FirstName, arg.FirstName)
	require.Equal(t, user.LastName, arg.LastName)
	require.Equal(t, user.UserName, arg.UserName)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
	require.Equal(t, user.Address, arg.Address)
	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.ContactNumber, arg.ContactNumber)

	require.NotZero(t, user.CreatedAt)
}

func createUser(t *testing.T, arg db.CreateUserParam) uuid.UUID {
	if arg.UserName == "" {
		arg = db.CreateUserParam{
			FirstName:      utils.RandomName(),
			LastName:       utils.RandomName(),
			HashedPassword: "secret",
			UserName:       utils.RandomName(),
			Address:        utils.RandomString(15),
			Email:          utils.RandomString(4),
			ContactNumber:  uint32(utils.RandomContactNum()),
		}
	}
	id, err := testSQLStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotEqual(t, id, uuid.Nil)

	return id
}

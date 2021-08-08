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

	createUser(t)
}

func TestFindUserByName(t *testing.T) {
	t.Parallel()

	user := createUser(t)

	user1, err := testSQLStore.FindUserByName(context.Background(), user.UserName)
	require.NoError(t, err)
	require.NotNil(t, user)

	require.NotEqual(t, user.ID, uuid.Nil)

	require.Equal(t, user.FirstName, user1.FirstName)
	require.Equal(t, user.LastName, user1.LastName)
	require.Equal(t, user.UserName, user1.UserName)
	require.Equal(t, user.HashedPassword, user1.HashedPassword)
	require.Equal(t, user.Address, user1.Address)
	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.ContactNumber, user1.ContactNumber)
	require.Equal(t, user.CreatedAt, user1.CreatedAt)
}

func TestFindUserById(t *testing.T) {
	t.Parallel()

	user := createUser(t)

	user1, err := testSQLStore.FindUserById(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotNil(t, user)

	require.NotEqual(t, user.ID, uuid.Nil)

	require.Equal(t, user.FirstName, user1.FirstName)
	require.Equal(t, user.LastName, user1.LastName)
	require.Equal(t, user.UserName, user1.UserName)
	require.Equal(t, user.HashedPassword, user1.HashedPassword)
	require.Equal(t, user.Address, user1.Address)
	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.ContactNumber, user1.ContactNumber)
	require.Equal(t, user.CreatedAt, user1.CreatedAt)
}

func createUser(t *testing.T) db.User {

	password, err := utils.HashPassword("secret")
	require.NoError(t, err)

	arg := db.CreateUserParam{
		FirstName:      utils.RandomName(),
		LastName:       utils.RandomName(),
		HashedPassword: password,
		UserName:       utils.RandomName(),
		Address:        utils.RandomString(15),
		Email:          utils.RandomString(4),
		ContactNumber:  uint32(utils.RandomContactNum()),
	}

	user, err := testSQLStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEqual(t, user.ID, uuid.Nil)

	require.Equal(t, user.FirstName, arg.FirstName)
	require.Equal(t, user.LastName, arg.LastName)
	require.Equal(t, user.UserName, arg.UserName)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
	require.Equal(t, user.Address, arg.Address)
	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.ContactNumber, arg.ContactNumber)

	require.NotZero(t, user.CreatedAt)
	return user
}

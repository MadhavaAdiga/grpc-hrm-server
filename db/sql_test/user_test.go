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
	createUser(t)
}

func createUser(t *testing.T) {
	arg := db.CreateUserParam{
		FirstName:      utils.RandomName(),
		LastName:       utils.RandomName(),
		HashedPassword: "secret",
		UserName:       utils.RandomName(),
		Address:        utils.RandomString(15),
		Email:          utils.RandomString(4),
		ContactNumber:  uint32(utils.RandomContactNum()),
	}

	id, err := testSQLStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotEqual(t, id, uuid.Nil)
}

func TestFindUser(t *testing.T) {

}

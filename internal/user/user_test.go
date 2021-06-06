package user_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	mockdb "github.com/MadhavaAdiga/grpc-hrm-server/db/mock"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/user"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	usrId := uuid.New()

	req := &hrm.CreateUserRequest{
		FirstName:     utils.RandomName(),
		LastName:      utils.RandomName(),
		UserName:      utils.RandomName(),
		Password:      "secret",
		Address:       utils.RandomString(8),
		EmailId:       "a@example.com",
		ContactNumber: 1234567890,
	}

	testcase := []struct {
		name          string
		req           *hrm.CreateUserRequest
		buildStub     func(store *mockdb.MockStore)
		checkresponse func(t *testing.T, res *hrm.CreateUserResponse)
	}{
		{
			name: "Best case",
			req:  req,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(usrId, nil)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateUserResponse) {
				id, err := uuid.Parse(res.GetId())
				require.NoError(t, err)
				require.NotEqual(t, uuid.Nil, id)
			},
		},
	}

	for i := range testcase {
		test := testcase[i]

		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)

			test.buildStub(store)

			serverAddr := startTestServer(t, store)
			client := createTestClient(t, serverAddr)

			// create new user
			res, err := client.CreateUser(context.Background(), test.req)
			require.NoError(t, err)
			require.NotNil(t, res)

			test.checkresponse(t, res)
		})
	}

}

func TestFindUserByName(t *testing.T) {
	t.Parallel()

	userName := utils.RandomName()

	user := db.User{
		ID:             uuid.New(),
		FirstName:      utils.RandomName(),
		LastName:       utils.RandomName(),
		UserName:       userName,
		HashedPassword: "secret",
		Address:        utils.RandomString(8),
		Email:          "a@example.com",
		ContactNumber:  uint32(utils.RandomContactNum()),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(1).Return(user, nil)

	serverAddr := startTestServer(t, store)
	client := createTestClient(t, serverAddr)

	arg := &hrm.FindUserRequest{
		UserName: userName,
	}
	res, err := client.FindUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, res.User.Id, uuid.Nil.String())
	require.NotZero(t, res.User.Createdat)

}

func startTestServer(t *testing.T, store db.Store) string {

	server := user.NewUserServer(store, hclog.Default())

	grpcServer := grpc.NewServer()
	hrm.RegisterUserServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

func createTestClient(t *testing.T, serverAddress string) hrm.UserServiceClient {
	// connect to grpc server, insecure is used for testing
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return hrm.NewUserServiceClient(conn)
}

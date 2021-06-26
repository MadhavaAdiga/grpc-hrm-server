package user_test

import (
	"context"
	"database/sql"
	"errors"
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
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

/*
 Table driven test for CreateUser RPC
*/
func TestCreateUser(t *testing.T) {
	t.Parallel()

	user := db.User{
		ID:             uuid.New(),
		FirstName:      utils.RandomName(),
		LastName:       utils.RandomName(),
		HashedPassword: "secret",
		UserName:       utils.RandomName(),
		Address:        utils.RandomString(15),
		Email:          "abc@email.com",
		ContactNumber:  1234567890,
	}

	reqStub := &hrm.CreateUserRequest{
		FirstName:     user.UserName,
		LastName:      user.LastName,
		UserName:      user.UserName,
		Password:      user.HashedPassword,
		Address:       user.Address,
		EmailId:       user.Email,
		ContactNumber: user.ContactNumber,
	}

	testcase := []struct {
		name          string
		buildReq      func(t *testing.T, req *hrm.CreateUserRequest) *hrm.CreateUserRequest
		buildStub     func(store *mockdb.MockStore)
		checkresponse func(t *testing.T, res *hrm.CreateUserResponse, err error)
	}{
		{
			name: "Best case",
			buildReq: func(t *testing.T, req *hrm.CreateUserRequest) *hrm.CreateUserRequest {
				return req
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(user, nil)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateUserResponse, err error) {
				require.NoError(t, err)
				id, err := uuid.Parse(res.GetId())
				require.NoError(t, err)
				require.NotEqual(t, uuid.Nil, id)
			},
		}, {
			name: "Repeated Username",
			buildReq: func(t *testing.T, req *hrm.CreateUserRequest) *hrm.CreateUserRequest {
				return req
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(db.User{}, sql.ErrNoRows)

			},
			checkresponse: func(t *testing.T, res *hrm.CreateUserResponse, err error) {
				require.Error(t, err, "unique_violation")
				id, err := uuid.Parse(res.GetId())
				require.Error(t, err, "invalid UUID length")
				require.Equal(t, uuid.Nil, id)
			},
		}, {
			name: "invalid email id",
			buildReq: func(t *testing.T, req *hrm.CreateUserRequest) *hrm.CreateUserRequest {
				other := &hrm.CreateUserRequest{}
				err := copier.Copy(other, req)
				require.NoError(t, err)

				other.EmailId = "asdsa"
				return other
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkresponse: func(t *testing.T, res *hrm.CreateUserResponse, err error) {
				require.Error(t, err)
			},
		}, {
			name: "Worst case",
			buildReq: func(t *testing.T, req *hrm.CreateUserRequest) *hrm.CreateUserRequest {
				other := &hrm.CreateUserRequest{}
				err := copier.Copy(other, req)
				require.NoError(t, err)

				other.FirstName = ""
				other.LastName = ""
				other.UserName = ""
				other.Password = ""
				other.ContactNumber = 12
				return other
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.All()).
					Times(0).Return(db.User{}, errors.New("missing argument"))
			},
			checkresponse: func(t *testing.T, res *hrm.CreateUserResponse, err error) {
				require.Error(t, err)
				id, err := uuid.Parse(res.GetId())
				require.Error(t, err)
				require.Equal(t, uuid.Nil, id)
			},
		},
	}

	for i := range testcase {
		test := testcase[i]

		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)

			// building stub for mock db
			test.buildStub(store)

			// create server and client for test
			serverAddr := startTestServer(t, store)
			client := createTestClient(t, serverAddr)

			// get test request
			req := test.buildReq(t, reqStub)

			// create new user
			res, err := client.CreateUser(context.Background(), req)
			// checking for valid response by test
			test.checkresponse(t, res, err)
		})
	}

}

/*
	Table drive test for FindUser RPC
*/
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

	testcase := []struct {
		name          string
		username      string
		buildStub     func(store *mockdb.MockStore)
		checkresponse func(t *testing.T, res *hrm.FindUserResponse, err error)
	}{
		{
			name:     "Best Case",
			username: user.UserName,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(1).Return(user, nil)
			},
			checkresponse: func(t *testing.T, res *hrm.FindUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEqual(t, res.User.Id, uuid.Nil.String())
				require.NotZero(t, res.User.Createdat)
			},
		}, {
			name:     "Username not found",
			username: "does not exisits",
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(1).Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, res *hrm.FindUserResponse, err error) {
				require.Error(t, err, sql.ErrNoRows)
				require.Nil(t, res)
			},
		}, {

			name:     "invalid username",
			username: "",
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(0)
			},
			checkresponse: func(t *testing.T, res *hrm.FindUserResponse, err error) {
				require.Error(t, err)
			},
		},
	}

	for i := range testcase {
		test := testcase[i]

		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)

			// create server and client for test
			serverAddr := startTestServer(t, store)
			client := createTestClient(t, serverAddr)

			// building stub for mock db
			test.buildStub(store)

			arg := &hrm.FindUserRequest{
				UserName: test.username,
			}
			res, err := client.FindUser(context.Background(), arg)
			// checking for valid response by test
			test.checkresponse(t, res, err)

		})

	}

}

// Helper function start a test server
func startTestServer(t *testing.T, store db.Store) string {

	server := user.NewUserServer(store, hclog.Default())

	grpcServer := grpc.NewServer()
	hrm.RegisterUserServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

// Helper function create a client
func createTestClient(t *testing.T, serverAddress string) hrm.UserServiceClient {
	// connect to grpc server, insecure is used for testing
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return hrm.NewUserServiceClient(conn)
}

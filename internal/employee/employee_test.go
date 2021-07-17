package employee_test

import (
	"context"
	"database/sql"
	"net"
	"testing"
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	mockdb "github.com/MadhavaAdiga/grpc-hrm-server/db/mock"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/employee"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestCreateEmployee(t *testing.T) {
	t.Parallel()

	orgName := utils.RandomName()
	empId := uuid.New()
	userName := utils.RandomName()
	roleName := utils.RandomName()

	employee := db.Employee{
		ID: uuid.New(),
	}

	creator := db.Employee{
		ID:           empId,
		Organization: db.Organization{ID: uuid.New(), Name: orgName},
		Role:         db.Role{ID: uuid.New(), Permissions: []int32{1, 6}},
	}

	user := db.User{
		ID:       uuid.New(),
		UserName: userName,
	}

	role := db.Role{
		ID:   uuid.New(),
		Name: roleName,
	}

	reqStub := &hrm.CreateEmployeeRequest{
		UserName:         userName,
		OrganizationName: orgName,
		RoleName:         roleName,
		CreatorId:        empId.String(),
	}

	testcase := []struct {
		name          string
		buildReq      func(t *testing.T, req *hrm.CreateEmployeeRequest) *hrm.CreateEmployeeRequest
		buildStubs    func(store *mockdb.MockStore)
		checkresponse func(t *testing.T, res *hrm.CreateEmployeeResponse, err error)
	}{
		{
			name: "Best case",
			buildReq: func(t *testing.T, req *hrm.CreateEmployeeRequest) *hrm.CreateEmployeeRequest {
				return req
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(creator, nil)
				store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(1).Return(user, nil)
				store.EXPECT().FindRoleByOrganizationName(gomock.Any(), gomock.All()).Times(1).Return(role, nil)
				store.EXPECT().CreateEmployee(gomock.Any(), gomock.All()).Times(1).Return(employee, nil)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateEmployeeResponse, err error) {
				require.NoError(t, err)
				require.NotEqual(t, uuid.Nil, res)
			},
		},
		{
			name: "Invalid creatorId",
			buildReq: func(t *testing.T, req *hrm.CreateEmployeeRequest) *hrm.CreateEmployeeRequest {
				return req
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(db.Employee{}, sql.ErrNoRows)
				store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(0)
				store.EXPECT().FindRoleByOrganizationName(gomock.Any(), gomock.All()).Times(0)
				store.EXPECT().CreateEmployee(gomock.Any(), gomock.All()).Times(0)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateEmployeeResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
			},
		}, {
			name: "Invalid username",
			buildReq: func(t *testing.T, req *hrm.CreateEmployeeRequest) *hrm.CreateEmployeeRequest {
				other := &hrm.CreateEmployeeRequest{}
				err := copier.Copy(other, req)
				require.NoError(t, err)

				other.UserName = utils.RandomName()

				return other

			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(creator, nil)
				store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(1).Return(user, nil)
				store.EXPECT().FindRoleByOrganizationName(gomock.Any(), gomock.All()).Times(1).Return(db.Role{}, sql.ErrNoRows)
				store.EXPECT().CreateEmployee(gomock.Any(), gomock.All()).Times(0)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateEmployeeResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
			},
		}, {
			name: "Invalid roleName",
			buildReq: func(t *testing.T, req *hrm.CreateEmployeeRequest) *hrm.CreateEmployeeRequest {
				other := &hrm.CreateEmployeeRequest{}
				err := copier.Copy(other, req)
				require.NoError(t, err)

				other.RoleName = utils.RandomName()

				return other

			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(creator, nil)
				store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(1).Return(db.User{}, sql.ErrNoRows)
				store.EXPECT().FindRoleByOrganizationName(gomock.Any(), gomock.All()).Times(0)
				store.EXPECT().CreateEmployee(gomock.Any(), gomock.All()).Times(0)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateEmployeeResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
			},
		},
	}

	for _, test := range testcase {

		t.Run(test.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			// building stub for mock db
			test.buildStubs(store)

			// create server and client for test
			serverAddr := startTestServer(t, store)
			client := createTestClient(t, serverAddr)

			// get test request
			req := test.buildReq(t, reqStub)

			// create new user
			res, err := client.CreateEmployee(context.Background(), req)
			// checking for valid response by test
			test.checkresponse(t, res, err)
		})
	}
}

func TestFindEmployee(t *testing.T) {
	t.Parallel()

	orgName := utils.RandomName()
	userName := utils.RandomName()

	employee := db.Employee{
		ID:           uuid.New(),
		User:         db.User{UserName: userName},
		Organization: db.Organization{Name: orgName},
		Role:         db.Role{},
		Status:       0,
		CreateBy:     [16]byte{},
		UpdatedBy:    [16]byte{},
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	reqStub := &hrm.FindEmployeeRequest{
		Filter: &hrm.EmployeeFilter{
			OrganizationName: orgName,
			UserName:         userName,
		},
	}

	testcase := []struct {
		name          string
		buildReq      func(t *testing.T, req *hrm.FindEmployeeRequest) *hrm.FindEmployeeRequest
		buildStubs    func(store *mockdb.MockStore)
		checkresponse func(t *testing.T, res *hrm.FindEmployeeResponse, err error)
	}{
		{
			name: "Best case",
			buildReq: func(t *testing.T, req *hrm.FindEmployeeRequest) *hrm.FindEmployeeRequest {
				return req
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindEmployeeByUnameAndOrg(gomock.Any(), gomock.All()).Times(1).Return(employee, nil)
			},
			checkresponse: func(t *testing.T, res *hrm.FindEmployeeResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
			},
		}, {
			name: "Not found",
			buildReq: func(t *testing.T, req *hrm.FindEmployeeRequest) *hrm.FindEmployeeRequest {
				return req
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindEmployeeByUnameAndOrg(gomock.Any(), gomock.All()).Times(1).Return(db.Employee{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, res *hrm.FindEmployeeResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
			},
		},
	}

	for _, test := range testcase {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			// building stub for mock db
			test.buildStubs(store)

			// create server and client for test
			serverAddr := startTestServer(t, store)
			client := createTestClient(t, serverAddr)

			// get test request
			req := test.buildReq(t, reqStub)

			// create new user
			res, err := client.FindEmployee(context.Background(), req)
			// checking for valid response by test
			test.checkresponse(t, res, err)
		})
	}

}

// Helper function start a test server
func startTestServer(t *testing.T, store db.Store) string {

	server := employee.NewEmployeeServer(store, hclog.Default())

	grpcServer := grpc.NewServer()
	hrm.RegisterEmployeeServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

// Helper function create a client
func createTestClient(t *testing.T, serverAddress string) hrm.EmployeeServiceClient {
	// connect to grpc server, insecure is used for testing
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return hrm.NewEmployeeServiceClient(conn)
}

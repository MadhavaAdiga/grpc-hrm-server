package role_test

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	mockdb "github.com/MadhavaAdiga/grpc-hrm-server/db/mock"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/role"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestCreateRole(t *testing.T) {
	t.Parallel()

	empId := uuid.New()
	orgName := utils.RandomName()
	// stubs
	role := db.Role{
		ID: uuid.New(),
	}
	emp := db.Employee{
		ID:           empId,
		Organization: db.Organization{Name: orgName},
		Role:         db.Role{ID: uuid.New(), Permissions: []int32{6}},
	}
	org := db.Organization{
		ID:        uuid.New(),
		Name:      orgName,
		CreatorID: empId,
	}
	reqStub := &hrm.CreateRoleRequest{
		Name:             utils.RandomName(),
		OrganizationName: orgName,
		Permissions:      []hrm.Permission{hrm.Permission_ADMIN, hrm.Permission_CAN_ADD_EMPLOYEE},
		CreatorId:        uuid.New().String(),
	}

	testcase := []struct {
		name          string
		buildReq      func(t *testing.T, req *hrm.CreateRoleRequest) *hrm.CreateRoleRequest
		buildStubs    func(store *mockdb.MockStore)
		checkresponse func(t *testing.T, res *hrm.CreateRoleResponse, err error)
	}{
		{
			name: "Best case",
			buildReq: func(t *testing.T, req *hrm.CreateRoleRequest) *hrm.CreateRoleRequest {
				return reqStub
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(emp, nil)
				store.EXPECT().FindOrganizationByName(gomock.Any(), gomock.Eq(orgName)).Times(1).Return(org, nil)
				store.EXPECT().CreateRole(gomock.Any(), gomock.Any()).Times(1).Return(role, nil)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateRoleResponse, err error) {
				require.NoError(t, err)
				require.NotEqual(t, uuid.Nil, res)
			},
		},
		{
			name: "invalid creatorId",
			buildReq: func(t *testing.T, req *hrm.CreateRoleRequest) *hrm.CreateRoleRequest {
				return req
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(db.Employee{}, errors.New("not found"))
				store.EXPECT().FindOrganizationByName(gomock.Any(), gomock.Eq(orgName)).Times(0)
				store.EXPECT().CreateRole(gomock.Any(), gomock.Any()).Times(0)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateRoleResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
			},
		},
		{
			name: "Invalid Permission set",
			buildReq: func(t *testing.T, req *hrm.CreateRoleRequest) *hrm.CreateRoleRequest {
				other := &hrm.CreateRoleRequest{}
				err := copier.Copy(other, req)
				require.NoError(t, err)

				other.Permissions = []hrm.Permission{13}

				return other
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(0)
				store.EXPECT().FindOrganizationByName(gomock.Any(), gomock.Any()).Times(0)
				store.EXPECT().CreateRole(gomock.Any(), gomock.Any()).Times(0)
			},
			checkresponse: func(t *testing.T, res *hrm.CreateRoleResponse, err error) {
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
			res, err := client.CreateRole(context.Background(), req)
			// checking for valid response by test
			test.checkresponse(t, res, err)

		})

	}

}

// Helper function start a test server
func startTestServer(t *testing.T, store db.Store) string {

	server := role.NewRoleServer(store, hclog.Default())

	grpcServer := grpc.NewServer()
	hrm.RegisterRoleServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

// Helper function create a client
func createTestClient(t *testing.T, serverAddress string) hrm.RoleServiceClient {
	// connect to grpc server, insecure is used for testing
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return hrm.NewRoleServiceClient(conn)
}

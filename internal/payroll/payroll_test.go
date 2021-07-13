package payroll_test

import (
	"context"
	"net"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	mockdb "github.com/MadhavaAdiga/grpc-hrm-server/db/mock"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/payroll"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestCreatePayroll(t *testing.T) {
	t.Parallel()

	creatorId := uuid.New()
	orgName := utils.RandomName()
	empName := utils.RandomName()

	creator := db.Employee{
		ID:           creatorId,
		Organization: db.Organization{Name: orgName},
		Role:         db.Role{ID: uuid.New(), Permissions: []int32{6, 4}},
	}
	emp := db.Employee{
		ID: uuid.New(),
	}
	reqStub := &hrm.AddPayrollRequest{
		Username: empName,
		Ctc: &hrm.AddPayrollRequest_Yearly{
			Yearly: utils.RandomInt(0, 10000),
		},
		Allowance:        utils.RandomInt(0, 3000),
		OrganizationName: orgName,
		CreatorId:        creatorId.String(),
	}

	testcase := []struct {
		name          string
		buildReq      func(t *testing.T, req *hrm.AddPayrollRequest) *hrm.AddPayrollRequest
		buildStubs    func(store *mockdb.MockStore)
		checkresponse func(t *testing.T, res *hrm.PayrollResponse, err error)
	}{
		{
			name: "Best case",
			buildReq: func(t *testing.T, req *hrm.AddPayrollRequest) *hrm.AddPayrollRequest {
				return req
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(creator, nil)
				store.EXPECT().FindEmployeeByUnameAndOrg(gomock.Any(), gomock.All()).Times(1).Return(emp, nil)
				store.EXPECT().CreatePayroll(gomock.Any(), gomock.All()).Times(1).Return(db.Payroll{}, nil)
			},
			checkresponse: func(t *testing.T, res *hrm.PayrollResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
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
			res, err := client.AddPayroll(context.Background(), req)
			// checking for valid response by test
			test.checkresponse(t, res, err)
		})
	}

}

// Helper function start a test server
func startTestServer(t *testing.T, store db.Store) string {

	server := payroll.NewPayrollServe(store, hclog.Default())

	grpcServer := grpc.NewServer()
	hrm.RegisterPayrollServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

// Helper function create a client
func createTestClient(t *testing.T, serverAddress string) hrm.PayrollServiceClient {
	// connect to grpc server, insecure is used for testing
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return hrm.NewPayrollServiceClient(conn)
}

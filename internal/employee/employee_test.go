package employee_test

import (
	"context"
	"net"
	"testing"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	mockdb "github.com/MadhavaAdiga/grpc-hrm-server/db/mock"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/employee"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestCreateEMployee(t *testing.T) {
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

	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(creator, nil)
	store.EXPECT().FindUserByName(gomock.Any(), gomock.All()).Times(1).Return(user, nil)
	store.EXPECT().FindRoleByOrganizationName(gomock.Any(), gomock.All()).Times(1).Return(role, nil)
	store.EXPECT().CreateEmployee(gomock.Any(), gomock.All()).Times(1).Return(employee, nil)

	serverAddr := startTestServer(t, store)
	client := createTestClient(t, serverAddr)

	res, err := client.CreateEmployee(context.Background(), reqStub)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, res)
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

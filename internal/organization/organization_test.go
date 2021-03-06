package organization_test

import (
	"context"
	"net"
	"testing"

	mockdb "github.com/MadhavaAdiga/grpc-hrm-server/db/mock"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/internal/organization"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestCreateOrganization(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	empId := uuid.New()

	emp := db.Employee{
		ID:           empId,
		Organization: db.Organization{Name: "HRM_GRPC"},
		User:         db.User{},
		Role:         db.Role{ID: uuid.New(), Permissions: []int32{6}},
		Status:       1,
		CreateBy:     uuid.New(),
	}

	store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).Times(1).Return(emp, nil)
	store.EXPECT().CreateOrganization(gomock.Any(), gomock.Any()).Times(1)

	// store.EXPECT().FindAdminEmployee(gomock.Any(), gomock.All()).AnyTimes().Return(db.Employee{}, sql.ErrNoRows)
	// store.EXPECT().CreateOrganization(gomock.Any(), gomock.Any()).Times(0)

	serverAddr := startTestServer(t, store)
	client := createTestClient(t, serverAddr)

	req := &hrm.CreateOrganizationRequest{
		Name:      utils.RandomName(),
		CreatorId: uuid.New().String(),
	}

	res, err := client.CreateOrganization(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, res.Id, uuid.Nil)

}

// TODO : do a table test
func TestFindOrganization(t *testing.T) {
	t.Parallel()

	org := db.Organization{
		ID:   uuid.New(),
		Name: utils.RandomName(),
		// CreatedBy: utils.RandomName(),
		Status:    0,
		CreatorID: uuid.New(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	serverAdr := startTestServer(t, store)
	client := createTestClient(t, serverAdr)

	req := &hrm.CreateOrganizationRequest{
		Name:      org.Name,
		CreatorId: org.CreatorID.String(),
	}

	store.EXPECT().FindOrganizationByName(gomock.Any(), gomock.Eq(req.Name)).Times(1).Return(org, nil)

	res2, err := client.FindOrganization(context.Background(), &hrm.FindOrganizationRequest{
		Name: req.Name,
	})
	require.NoError(t, err)
	require.NotNil(t, res2)

	require.Equal(t, res2.GetOrganization().Name, req.Name)
}

// Helper function start a test server
func startTestServer(t *testing.T, store db.Store) string {

	server := organization.NewOrganizationServer(store, hclog.Default())

	grpcServer := grpc.NewServer()
	hrm.RegisterOrganizationServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

// Helper function create a client
func createTestClient(t *testing.T, serverAddress string) hrm.OrganizationServiceClient {
	// connect to grpc server, insecure is used for testing
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return hrm.NewOrganizationServiceClient(conn)
}

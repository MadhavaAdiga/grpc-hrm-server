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

	store.EXPECT().CreateOrganization(gomock.Any(), gomock.Any()).Times(1)

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

// func TestFindOrganization(t *testing.T) {
// 	t.Parallel()

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	store := mockdb.NewMockStore(ctrl)

// 	store.EXPECT().CreateOrganization(gomock.Any(), gomock.Any()).Times(1)

// 	serverAdr := startTestServer(t, store)
// 	client := createTestClient(t, serverAdr)

// 	req := &hrm.CreateOrganizationRequest{
// 		Name:      utils.RandomName(),
// 		CreatorId: uuid.New().String(),
// 	}
// 	// create a new organiztion
// 	res1, err := client.CreateOrganization(context.Background(), req)
// 	require.NoError(t, err)
// 	require.NotNil(t, res1)

// 	var organization db.Organization
// 	store.EXPECT().FindOrganizationByName(gomock.Any(), gomock.Any()).Times(1).Return(organization, nil)

// 	res2, err := client.FindOrganization(context.Background(), &hrm.FindOrganizationRequest{
// 		Name: req.Name,
// 	})
// 	require.NoError(t, err)
// 	require.NotNil(t, res2)

// 	require.Equal(t, res1.Id, res2.GetOrganization().Id)
// 	// require.Equal(t, res2.GetOrganization().Name, req.Name)
// }

func startTestServer(t *testing.T, store db.Store) string {

	server := organization.NewOrganizationServer(store, hclog.Default())

	grpcServer := grpc.NewServer()
	hrm.RegisterOrganizatoinServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

func createTestClient(t *testing.T, serverAddress string) hrm.OrganizatoinServiceClient {
	// connect to grpc server, insecure is used for testing
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return hrm.NewOrganizatoinServiceClient(conn)
}

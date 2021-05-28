package organization

import (
	"context"

	db "github.com/MadhavaAdiga/grpc-hrm-server/db/sql"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
)

type OrganizationServer struct {
	store db.Store
	hrm.UnimplementedOrganizatoinServiceServer
}

func NewOrganizationServer(store db.Store) *hrm.OrganizatoinServiceServer {
	return &OrganizationServer{
		store: store,
	}
}

// rpc to create a new organization
func (service *OrganizationServer) CreateOrganization(ctx context.Context, req *hrm.CreateOrganizationRequest) (*hrm.CreateOrganizationResponse, error) {

}

// rpc to search organization
func (service *OrganizationServer) SearchOrganization(ctx context.Context, req *hrm.SearchOrganizationRequest) (*hrm.SearchOrganizationResponse, error) {

}

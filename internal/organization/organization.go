package organization

import (
	"context"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrganizationServer struct {
	store db.Store
	hrm.UnimplementedOrganizatoinServiceServer
}

func NewOrganizationServer(store db.Store) *OrganizationServer {
	return &OrganizationServer{
		store: store,
	}
}

// rpc to create a new organization
func (server *OrganizationServer) CreateOrganization(ctx context.Context, req *hrm.CreateOrganizationRequest) (*hrm.CreateOrganizationResponse, error) {

	title := req.GetName()
	creatorID := req.GetCreatorId()

	// check for valid title
	if len(title) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "title is required")
	}

	// check for a valid uuid
	_, err := uuid.Parse(creatorID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "creator id is not a valid uuid: %w", err)
	}

	// todo check if its a valid cretor id by checking in db

	arg := db.CreateOrganizationParam{
		Name:      title,
		CreatedBy: utils.RandomName(),
		CreatorID: creatorID,
		Status:    uint16(hrm.Organization_ACTIVE),
	}

	// save to store
	organization, err := server.store.CreateOrganization(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create new organization: %w", err)
	}

	res := &hrm.CreateOrganizationResponse{
		Id: organization.ID.String(),
	}

	return res, nil
}

// rpc to search organization
func (service *OrganizationServer) FindOrganization(ctx context.Context, req *hrm.FindOrganizationRequest) (*hrm.FindOrganizationResponse, error) {
	return nil, nil
}

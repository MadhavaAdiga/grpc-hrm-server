package organization

import (
	"context"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrganizationServer struct {
	store db.Store
	log   hclog.Logger
	hrm.UnimplementedOrganizationServiceServer
}

func NewOrganizationServer(store db.Store, l hclog.Logger) hrm.OrganizationServiceServer {
	return &OrganizationServer{
		store: store,
		log:   l.Named("organization_server"),
	}
}

// RPC implementation to create a new organization
func (server *OrganizationServer) CreateOrganization(ctx context.Context, req *hrm.CreateOrganizationRequest) (*hrm.CreateOrganizationResponse, error) {
	// name of organization
	title := req.GetName()
	creatorID := req.GetCreatorId()

	// check for valid title
	if len(title) < 0 {
		server.log.Info("invalid title, empty")
		return nil, status.Errorf(codes.InvalidArgument, "title is required")
	}

	// check for a valid uuid
	_, err := uuid.Parse(creatorID)
	if err != nil {
		server.log.Info("invalid uuid", "error", err)
		return nil, status.Errorf(codes.InvalidArgument, "creator id is not a valid uuid: %v", err)
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
		pqError, ok := err.(*pq.Error)
		if ok {
			switch pqError.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "organization with name:%s alreaddy exists", title)
			}
		}
		server.log.Info("failed to create organization", "error", err)
		return nil, status.Errorf(codes.Internal, "unable to create new organization: %v", err)
	}

	res := &hrm.CreateOrganizationResponse{
		Id: organization.ID.String(),
	}

	return res, nil
}

// RPC implemetation to search organization
func (server *OrganizationServer) FindOrganization(ctx context.Context, req *hrm.FindOrganizationRequest) (*hrm.FindOrganizationResponse, error) {
	// name of organization
	title := req.GetName()

	// find record from db
	organization, err := server.store.FindOrganizationByName(ctx, title)
	if err != nil {
		server.log.Info("request organization not found", "error", err)
		return nil, status.Errorf(codes.NotFound, "organization is not found: %v", err)
	}

	// map db.Organization to protobuf organization message
	o := &hrm.Organization{
		Id:        organization.ID.String(),
		CreatedBy: organization.CreatedBy,
		CreatorId: organization.CreatorID.String(),
		Name:      organization.Name,
		Status:    hrm.Organization_Status(organization.Status),
		UpdatedBy: organization.UpdatedBy,
		UpdaterId: organization.UpdaterID.String(),
		CreatedAt: timestamppb.New(organization.CreatedAt),
		UpdatedAt: timestamppb.New(organization.UpdatedAt),
	}

	res := &hrm.FindOrganizationResponse{
		Organization: o,
	}

	return res, nil
}

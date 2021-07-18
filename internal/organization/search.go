package organization

import (
	"context"
	"database/sql"

	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// RPC implemetation to search organization
func (server *OrganizationServer) FindOrganization(ctx context.Context, req *hrm.FindOrganizationRequest) (*hrm.FindOrganizationResponse, error) {
	// name of organization
	title := req.GetName()

	// handele context error
	if err := utils.ContextError(ctx); err != nil {
		return nil, err
	}

	// find record from db
	organization, err := server.store.FindOrganizationByName(ctx, title)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("not found", "error", err)
			return nil, status.Errorf(codes.NotFound, "organization is not found: %v", err)
		}
		server.log.Info("request organization not found", "error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// map db.Organization to protobuf organization message
	o := &hrm.Organization{
		Id:        organization.ID.String(),
		Name:      organization.Name,
		CreatedBy: "",
		CreatorId: organization.CreatorID.String(),
		Status:    hrm.Organization_Status(organization.Status),
		UpdatedBy: "",
		UpdaterId: organization.UpdaterID.String(),
		CreatedAt: timestamppb.New(organization.CreatedAt),
		UpdatedAt: timestamppb.New(organization.UpdatedAt),
	}

	res := &hrm.FindOrganizationResponse{
		Organization: o,
	}

	return res, nil
}

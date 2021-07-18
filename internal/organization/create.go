package organization

import (
	"context"
	"database/sql"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RPC implementation to create a new organization
func (server *OrganizationServer) CreateOrganization(ctx context.Context, req *hrm.CreateOrganizationRequest) (*hrm.CreateOrganizationResponse, error) {
	// name of organization
	title := req.GetName()

	// check for valid title
	if len(title) < 0 {
		server.log.Info("invalid title, empty")
		return nil, status.Errorf(codes.InvalidArgument, "title is required")
	}

	// handele context error
	if err := utils.ContextError(ctx); err != nil {
		return nil, err
	}

	// check for a valid uuid
	creatorID, err := uuid.Parse(req.GetCreatorId())
	if err != nil {
		server.log.Info("invalid uuid", "error", err)
		return nil, status.Errorf(codes.InvalidArgument, "creator id is not a valid uuid: %v", err)
	}

	// param for find operation
	// employee of HRM_GRPC has previlage to create a company
	creatorArg := db.FindAdminEmployeeParam{
		OrganizationName: "HRM_GRPC",
		EmployeeId:       creatorID,
	}
	// check if creator exists and belongs to HRM_GRPC
	creator, err := server.store.FindAdminEmployee(ctx, creatorArg)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("invalid creator", "error", err)
			return nil, status.Errorf(codes.NotFound, "creator does not belong to organization: %v", err)
		}
		server.log.Info("failed to find creator", "error", err)
		return nil, status.Errorf(codes.Internal, "unable to create new organization: %v", err)
	}
	// check for admin previlage
	if !utils.CheckPermission(hrm.Permission_ADMIN, creator.Role.Permissions) {
		server.log.Info("invalid creator")
		return nil, status.Errorf(codes.PermissionDenied, "creator does not have Admin previlage")
	}

	arg := db.CreateOrganizationParam{
		Name: title,
		// CreatedBy: utils.RandomName(),
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

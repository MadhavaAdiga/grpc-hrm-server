package role

import (
	"context"
	"database/sql"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// create a role for the orgaization
func (server *RoleServer) CreateRole(ctx context.Context, req *hrm.CreateRoleRequest) (*hrm.CreateRoleResponse, error) {

	if len(req.GetName()) == 0 {
		server.log.Info("invalid argument", "role name is required")
		return nil, status.Errorf(codes.InvalidArgument, "role name is required")
	}

	// check for required permission
	if req.GetPermissions() == nil {
		server.log.Info("invalid argument", "permissions are not specified for the role")
		return nil, status.Errorf(codes.InvalidArgument, "permissions are not specified for the role")
	}
	// check for valid set of permissions
	for _, v := range req.GetPermissions() {
		if _, ok := hrm.Permission_value[v.String()]; !ok {
			server.log.Info("invalid argument", "contains invalid permission set")
			return nil, status.Errorf(codes.InvalidArgument, "contains invalid permission set")
		}
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
		OrganizationName: req.GetOrganizationName(),
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
		server.log.Info("invalid creator", "error", err)
		return nil, status.Errorf(codes.PermissionDenied, "creator does not have Admin previlage: %v", err)
	}

	// check if organization exists and use the uuid to create role
	o, err := server.store.FindOrganizationByName(ctx, req.GetOrganizationName())
	if err != nil {
		server.log.Info("internal", "error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// convert hrm.Permission to int32
	var permissions []int32 = make([]int32, len(req.GetPermissions()))
	for i, v := range req.GetPermissions() {
		permissions[i] = int32(v.Number())
	}

	// db argument
	arg := db.CreateRoleParam{
		Name:         req.GetName(),
		Active:       true,
		Organization: o.ID,
		Permissions:  permissions,
		CreatedBy:    creatorID,
	}
	// store to db
	role, err := server.store.CreateRole(ctx, arg)
	if err != nil {
		server.log.Info("failed to create role", "error", err)
		return nil, status.Errorf(codes.Internal, "unable to create new role: %v", err)
	}

	res := &hrm.CreateRoleResponse{
		Id: role.ID.String(),
	}

	return res, nil
}

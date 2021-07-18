package employee

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

// create a user employed by the organization
func (server *EmployeeServer) CreateEmployee(ctx context.Context, req *hrm.CreateEmployeeRequest) (*hrm.CreateEmployeeResponse, error) {

	// validate
	if len(req.GetUserName()) == 0 {
		server.log.Error("Username is empty")
		return nil, status.Errorf(codes.InvalidArgument, "Username is a required field")
	}
	if len(req.OrganizationName) == 0 {
		server.log.Error("Organization is empty")
		return nil, status.Errorf(codes.InvalidArgument, "Organization is a required field")
	}
	if len(req.GetRoleName()) == 0 {
		server.log.Error("RoleName is empty")
		return nil, status.Errorf(codes.InvalidArgument, "RoleName is a required field")
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

	// check for previlages to add employee
	reqiredPermissions := []hrm.Permission{hrm.Permission_CAN_ADD_EMPLOYEE, hrm.Permission_ADMIN}
	if !utils.CheckPermissions(reqiredPermissions, creator.Role.Permissions) {
		server.log.Info("invalid creator")
		return nil, status.Errorf(codes.PermissionDenied, "creator does not have reqired previlage")
	}

	// check if user exists
	user, err := server.store.FindUserByName(ctx, req.GetUserName())
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("not found", "error", err)
			return nil, status.Errorf(codes.NotFound, "user %s is not found: %v", req.GetUserName(), err)
		}
		server.log.Error("internal server error", "error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// check if role exists
	rArg := db.FindRoleByOrgNameParam{
		Name:             req.GetRoleName(),
		OrganizationName: req.GetOrganizationName(),
	}
	role, err := server.store.FindRoleByOrganizationName(ctx, rArg)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("not found", "error", err)
			return nil, status.Errorf(codes.NotFound, "role %s is not found: %v", req.GetRoleName(), err)
		}
		server.log.Error("internal server error", "error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	empArg := db.CreateEmployeeParam{
		UserId:         user.ID,
		OrganizationId: creator.Organization.ID,
		RoleId:         role.ID,
		Status:         int16(hrm.Employee_EMPLOYEED),
		CreatedBy:      creatorID,
	}
	employee, err := server.store.CreateEmployee(ctx, empArg)
	if err != nil {
		server.log.Info("failed to create employee", "error", err)
		return nil, status.Errorf(codes.Internal, "unable to create new employee: %v", err)
	}

	res := &hrm.CreateEmployeeResponse{
		Id: employee.ID.String(),
	}

	return res, nil
}

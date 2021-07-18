package payroll

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

// add payroll to and existing employee
func (server *PayrollServer) AddPayroll(ctx context.Context, req *hrm.AddPayrollRequest) (*hrm.PayrollResponse, error) {

	if len(req.GetUsername()) == 0 {
		server.log.Info("invalid username, empty")
		return nil, status.Errorf(codes.InvalidArgument, "username is required")
	}
	if len(req.OrganizationName) == 0 {
		server.log.Error("Organization name is empty")
		return nil, status.Errorf(codes.InvalidArgument, "Organization is a required field")
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
	// check if creator exists
	creator, err := server.store.FindAdminEmployee(ctx, creatorArg)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("invalid creator", "error", err)
			return nil, status.Errorf(codes.NotFound, "creator does not belong to organization: %v", err)
		}
		server.log.Info("failed to find creator", "error", err)
		return nil, status.Errorf(codes.Internal, "unable to create new organization: %v", err)
	}
	// check for previlages
	premissions := []hrm.Permission{hrm.Permission_ADMIN, hrm.Permission_CAN_ADD_PAY}
	if !utils.CheckPermissions(premissions, creator.Role.Permissions) {
		server.log.Info("invalid creator")
		return nil, status.Errorf(codes.PermissionDenied, "creator does not have required previlage")
	}

	// find employee
	empArg := db.FindEmployeeUnameAndOrgParam{
		OrganizationName: req.GetOrganizationName(),
		Username:         req.GetUsername(),
	}
	employee, err := server.store.FindEmployeeByUnameAndOrg(ctx, empArg)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("not found", "error", err)
			return nil, status.Errorf(codes.NotFound, "employee is not found: %v", err)
		}
		server.log.Info("request organization not found", "error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// store payroll of employee
	var ctc int32 = 0
	if req.GetYearly() != 0 {
		ctc = req.GetYearly()
	} else {
		ctc = req.GetMonthly() * 12
	}

	arg := db.CreatePayrollParam{
		Employee:  employee.ID,
		Ctc:       ctc,
		Allowance: req.GetAllowance(),
		CreatedBy: creatorID,
	}
	_, err = server.store.CreatePayroll(ctx, arg)
	if err != nil {
		server.log.Info("failed to create employee", "error", err)
		return nil, status.Errorf(codes.Internal, "unable to add employee payroll: %v", err)
	}

	res := &hrm.PayrollResponse{
		Username: req.GetUsername(),
	}

	return res, nil
}

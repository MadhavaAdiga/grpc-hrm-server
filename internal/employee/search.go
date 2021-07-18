package employee

import (
	"context"
	"database/sql"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// search for an employees in the organization
// return an employee response with limited data information
func (server *EmployeeServer) FindEmployee(ctx context.Context, req *hrm.FindEmployeeRequest) (*hrm.FindEmployeeResponse, error) {

	filter := req.GetFilter()
	if filter == nil {
		server.log.Info("invalid filter is nil")
		return nil, status.Errorf(codes.InvalidArgument, "filter is required")
	}
	if len(filter.GetOrganizationName()) == 0 {
		server.log.Info("invalid organization name is empty")
		return nil, status.Errorf(codes.InvalidArgument, "organization name is required")
	}
	if len(filter.GetUserName()) == 0 {
		server.log.Info("invalid user name is empty")
		return nil, status.Errorf(codes.InvalidArgument, "user name is required")
	}

	// handele context error
	if err := utils.ContextError(ctx); err != nil {
		return nil, err
	}

	// find arg
	arg := db.FindEmployeeUnameAndOrgParam{
		OrganizationName: filter.GetOrganizationName(),
		Username:         filter.GetUserName(),
	}
	employee, err := server.store.FindEmployeeByUnameAndOrg(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("not found", "error", err)
			return nil, status.Errorf(codes.NotFound, "employee is not found: %v", err)
		}
		server.log.Info("request organization not found", "error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	emp := &hrm.Employee{
		Id:           employee.ID.String(),
		User:         &hrm.User{Id: employee.User.ID.String(), UserName: employee.User.UserName},
		Organization: &hrm.Organization{Id: employee.Organization.ID.String(), Name: employee.Organization.Name},
		Role:         &hrm.Role{Id: employee.Role.ID.String(), Name: employee.Role.Name},
		Status:       hrm.Employee_EmployeeStatus(employee.Status),
		Payroll:      &hrm.Payroll{},
		CreateBy:     employee.CreateBy.String(),
		CreatedAt:    timestamppb.New(employee.CreatedAt),
		UpdatedBy:    "",
		UpdatedAt:    &timestamppb.Timestamp{},
	}

	res := &hrm.FindEmployeeResponse{
		Employee: emp,
	}

	return res, nil
}

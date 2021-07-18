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
	"google.golang.org/protobuf/types/known/timestamppb"
)

// find payroll of an employee
func (server *PayrollServer) FindEmployeePayroll(ctx context.Context, req *hrm.FindEmployeePayrollRequest) (*hrm.FindEmployeePayrollResponse, error) {

	filter := req.GetFilter()
	if filter == nil {
		server.log.Error("filter is nil")
		return nil, status.Errorf(codes.InvalidArgument, "filter is required")
	}

	// handele context error
	if err := utils.ContextError(ctx); err != nil {
		return nil, err
	}

	var payroll db.Payroll
	var err error
	// find record base on type of key
	switch filter.GetKey().(type) {
	case *hrm.PayrollFilter_Id:
		// check for valid uuid
		var id uuid.UUID
		id, err = uuid.Parse(filter.GetId())
		if err != nil {
			server.log.Info("invalid uuid", "error", err)
			return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %v", err)
		}
		// call helper func
		payroll, err = findByEmployeeId(id, ctx, server)
	case *hrm.PayrollFilter_EmployeeName:
		if len(req.Filter.GetEmployeeName()) == 0 {
			server.log.Info("invalid employee name, empty")
			return nil, status.Errorf(codes.InvalidArgument, "employee is required")
		}
		// call helper func
		payroll, err = findByEmployeeName(filter.GetEmployeeName(), ctx, server)
	}

	if err != nil {
		return nil, err
	}

	p := &hrm.Payroll{
		Id:         payroll.ID.String(),
		EmployeeId: payroll.Employee.ID.String(),
		Ctc:        payroll.Ctc,
		Allowance:  payroll.Allowance,
		CreateBy:   payroll.CreateBy.String(),
		CreatedAt:  timestamppb.New(payroll.CreatedAt),
		UpdatedBy:  payroll.UpdatedBy.String(),
		UpdatedAt:  timestamppb.New(payroll.UpdatedAt),
	}
	res := &hrm.FindEmployeePayrollResponse{
		Payroll: p,
	}

	return res, nil
}

// Helper function to perform find operation based of employee id
func findByEmployeeId(id uuid.UUID, ctx context.Context, server *PayrollServer) (db.Payroll, error) {
	p, err := server.store.FindPayrollByEmpID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("not found", "error", err)
			return db.Payroll{}, status.Errorf(codes.NotFound, "payroll of the employee is not found: %v", err)
		}
		server.log.Info("request payroll not found", "error", err)
		return db.Payroll{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return p, nil
}

// Helper function to perform find operation based of employee name
func findByEmployeeName(employeeName string, ctx context.Context, server *PayrollServer) (db.Payroll, error) {
	p, err := server.store.FindPayrollByEmpName(ctx, employeeName)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("not found", "error", err)
			return db.Payroll{}, status.Errorf(codes.NotFound, "payroll of the employee is not found: %v", err)
		}
		server.log.Info("request payroll not found", "error", err)
		return db.Payroll{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return p, nil
}

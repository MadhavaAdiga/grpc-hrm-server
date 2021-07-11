package payroll

import (
	"context"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/hashicorp/go-hclog"
)

type PayrollServer struct {
	store db.Store
	log   hclog.Logger
	hrm.UnimplementedPayrollServiceServer
}

func NewPayrollServe(store db.Store, l hclog.Logger) hrm.PayrollServiceServer {
	return &PayrollServer{
		store: store,
		log:   l.Named("payroll_server"),
	}
}

// update payroll of an existing employee
func (server *PayrollServer) UpdatePayroll(ctx context.Context, req *hrm.UpdatePayrollRequest) (*hrm.PayrollResponse, error) {
	return nil, nil
}

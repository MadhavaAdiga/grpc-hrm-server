package employee

import (
	"context"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/hashicorp/go-hclog"
)

type EmployeeServer struct {
	store db.Store
	log   hclog.Logger
	hrm.UnimplementedEmployeeServiceServer
}

func NewEmployeeServer(store db.Store, l hclog.Logger) hrm.EmployeeServiceServer {
	return &EmployeeServer{
		store: store,
		log:   l.Named("employee_server"),
	}
}

// search for an employees in the organization
func (server *EmployeeServer) FindEmployee(ctx context.Context, req *hrm.FindEmployeeRequest) (*hrm.FindEmployeeResponse, error) {
	return nil, nil
}

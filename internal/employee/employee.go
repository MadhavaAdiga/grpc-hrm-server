package employee

import (
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

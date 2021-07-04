package role

import (
	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/hashicorp/go-hclog"
)

type RoleServer struct {
	store db.Store
	log   hclog.Logger
	hrm.UnimplementedRoleServiceServer
}

func NewRoleServer(store db.Store, l hclog.Logger) hrm.RoleServiceServer {
	return &RoleServer{
		store: store,
		log:   l.Named("role_server"),
	}
}

package role

import (
	"context"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// create a role for the orgaization
func (server *RoleServer) CreateRole(ctx context.Context, req *hrm.CreateRoleRequest) (*hrm.CreateRoleResponse, error) {

	// create should check for admin permission

	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "role name is required")
	}

	if len(req.Permissions) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "permissions are not specified for the role")
	}

	return nil, nil
}

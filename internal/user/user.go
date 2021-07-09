package user

import (
	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/hashicorp/go-hclog"
)

type UserServer struct {
	store db.Store
	log   hclog.Logger
	hrm.UnimplementedUserServiceServer
}

func NewUserServer(s db.Store, l hclog.Logger) hrm.UserServiceServer {
	return &UserServer{
		store: s,
		log:   l.Named("user_server"),
	}
}

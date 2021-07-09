package organization

import (
	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/hashicorp/go-hclog"
)

type OrganizationServer struct {
	store db.Store
	log   hclog.Logger
	hrm.UnimplementedOrganizationServiceServer
}

func NewOrganizationServer(store db.Store, l hclog.Logger) hrm.OrganizationServiceServer {
	return &OrganizationServer{
		store: store,
		log:   l.Named("organization_server"),
	}
}

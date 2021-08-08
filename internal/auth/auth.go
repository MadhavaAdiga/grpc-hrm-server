package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const tokenDuration = time.Minute * 10

type AuthServer struct {
	store        db.Store
	log          hclog.Logger
	tokenManager TokenManager
	hrm.UnimplementedAuthServiceServer
}

func NewAuthServer(store db.Store, l hclog.Logger, manager TokenManager) hrm.AuthServiceServer {
	return &AuthServer{
		store:        store,
		log:          l.Named("auth_server"),
		tokenManager: manager,
	}
}

func (server *AuthServer) Login(ctx context.Context, req *hrm.LoginRequest) (*hrm.LoginResponse, error) {
	if len(req.GetUsername()) == 0 {
		server.log.Error("Username is empty")
		return nil, status.Errorf(codes.InvalidArgument, "Username is a required field")
	}
	if len(req.GetPassword()) == 0 {
		server.log.Error("Password is empty")
		return nil, status.Errorf(codes.InvalidArgument, "Password is a required field")
	}

	user, err := server.store.FindUserByName(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Error("not found", "error", err)
			return nil, status.Errorf(codes.NotFound, "user is not found: %v", err)
		}
		server.log.Error("request user not found", "error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// check if the user is an employee to get an employee previlages
	var isEmployee bool = true
	employee, err := server.store.FindAdminEmployeeByUserID(ctx, user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Error("employee not found", "error", err)
			isEmployee = false
		} else {
			server.log.Error("request user not found", "error", err)
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}
	}

	err = utils.ValidatePassword(user.HashedPassword, req.GetPassword())
	if err != nil {
		server.log.Error("invalid password", "error", err)
		return nil, status.Errorf(codes.Unauthenticated, "username: %s and password does not match", req.GetUsername())
	}

	var permission []hrm.Permission
	if isEmployee {
		for i, v := range employee.Role.Permissions {
			permission[i] = hrm.Permission(v)
		}
	} else {
		// Add a default permission for user to view profiles and salary
		permission = []hrm.Permission{hrm.Permission_CAN_VIEW_EMPLOYEE, hrm.Permission_CAN_VIEW_SALARIES}
	}

	// replace userid with short UID
	token, err := server.tokenManager.CreateToken(user.ID, tokenDuration, permission)
	if err != nil {
		server.log.Error("failed to generate access token", "error", err)
		return nil, status.Errorf(codes.Internal, "token creation failed")
	}

	res := &hrm.LoginResponse{
		AcessToken: token,
	}

	return res, nil
}

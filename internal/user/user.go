package user

import (
	"context"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// create a user of the system
func (server *UserServer) CreateUser(ctx context.Context, req *hrm.CreateUserRequest) (*hrm.CreateUserResponse, error) {
	// errs := []error{}

	if req.GetFirstName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "first name is a required field")
	}
	if req.GetUserName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user name is required")
	}
	// calculate the length of a number
	count := utils.NumberLen(int(req.GetContactNumber()))
	if count < 10 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid contact number, should be atleast of 10 digits")
	}

	password := req.GetPassword()
	// TODO check for strong password
	if len(password) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "password is a required field")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	// create a param to store value to db
	arg := db.CreateUserParam{
		FirstName:      req.GetFirstName(),
		LastName:       req.GetLastName(),
		UserName:       req.GetUserName(),
		HashedPassword: hashedPassword,
		Address:        req.GetAddress(),
		Email:          req.GetEmailId(),
		ContactNumber:  req.GetContactNumber(),
	}
	// store to db
	id, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create new user: %v", err)
	}
	// build response object
	res := &hrm.CreateUserResponse{
		Id: id.String(),
	}

	return res, nil
}

// find user of the system by name
func (server *UserServer) FindUser(ctx context.Context, req *hrm.FindUserRequest) (*hrm.FindUserResponse, error) {
	return nil, nil
}

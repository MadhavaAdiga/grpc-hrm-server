package user

import (
	"context"
	"errors"

	"github.com/MadhavaAdiga/grpc-hrm-server/db"
	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/MadhavaAdiga/grpc-hrm-server/utils"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create a user of the system
// RPC implemtation to add a user to the system
func (server *UserServer) CreateUser(ctx context.Context, req *hrm.CreateUserRequest) (*hrm.CreateUserResponse, error) {
	invalidArgErrs := []error{}

	// check for valid arguments
	if req.GetFirstName() == "" {
		server.log.Info("first name is empty")
		invalidArgErrs = append(invalidArgErrs, errors.New("first name is a required field"))
	}
	if req.GetUserName() == "" {
		server.log.Info("user name is empty")
		invalidArgErrs = append(invalidArgErrs, errors.New("user name is required"))
	}
	// calculate the length of a number
	count := utils.NumberLen(int(req.GetContactNumber()))
	if count < 10 {
		server.log.Info("not a valid contact number, must be atleast 10 digits")
		invalidArgErrs = append(invalidArgErrs, errors.New("invalid contact number, should be atleast of 10 digits"))
	}

	// check for valid emailId if passed
	if len(req.GetEmailId()) > 0 {
		if !utils.ValidateMail(req.EmailId) {
			server.log.Info("eamil id is invalid format")
			return nil, status.Errorf(codes.InvalidArgument, "invalid email-id format,must in a proper email-id format")
		}
	}

	password := req.GetPassword()
	// TODO check for strong password
	if len(password) == 0 {
		server.log.Info("password is empty")
		invalidArgErrs = append(invalidArgErrs, errors.New("password is a required field"))
	}
	// check if there are errors in the fields
	if len(invalidArgErrs) != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "missing arguments: %v", invalidArgErrs)
	}

	//  encrypt password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		server.log.Error("failed to hash password", "error", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	// handele context error
	if err := utils.ContextError(ctx); err != nil {
		return nil, err
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
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		pqError, ok := err.(*pq.Error)
		if ok {
			// check for db error
			switch pqError.Code.Name() {
			case "unique_violation":
				server.log.Info("failed to create user", "error", err)
				return nil, status.Errorf(codes.AlreadyExists, "user with user name:%s alreaddy exists", req.UserName)
			}
		}
		server.log.Info("failed to create user", "error", err)
		return nil, status.Errorf(codes.Internal, "unable to create new user: %v", err)
	}
	// build response object
	res := &hrm.CreateUserResponse{
		Id: user.ID.String(),
	}

	return res, nil
}

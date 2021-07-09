package user

import (
	"context"
	"database/sql"

	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//	Find user of the system by name
//	RPC implementation to find user existing in the system
func (server *UserServer) FindUser(ctx context.Context, req *hrm.FindUserRequest) (*hrm.FindUserResponse, error) {

	userName := req.UserName
	if len(userName) == 0 {
		server.log.Info("user is empty")
		return nil, status.Errorf(codes.InvalidArgument, "username is required")
	}

	user, err := server.store.FindUserByName(ctx, userName)
	if err != nil {
		if err == sql.ErrNoRows {
			server.log.Info("failed to find user", "error", err)
			return nil, status.Errorf(codes.NotFound, "username: %s does not exists", userName)
		}
		server.log.Error("server error", "error", err)
		return nil, status.Errorf(codes.Internal, "error: %v", err)
	}

	// build response
	res := &hrm.FindUserResponse{
		User: &hrm.User{
			Id:            user.ID.String(),
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			UserName:      user.UserName,
			Password:      "",
			Address:       user.Address,
			EmailId:       user.Email,
			ContactNumber: user.ContactNumber,
			Createdat:     timestamppb.New(user.CreatedAt),
			UpdatedAt:     timestamppb.New(user.UpdatedAt),
		},
	}

	return res, nil
}

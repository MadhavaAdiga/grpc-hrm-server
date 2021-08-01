package auth

import (
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
)

/*
  interface for managing a validation token
*/
type TokenManager interface {
	// generate a new token
	CreateToken(username string, durattion time.Duration, permissions []hrm.Permission) (string, error)
	// check if token is valid
	VerifyToken(token string) (*Payload, error)
}

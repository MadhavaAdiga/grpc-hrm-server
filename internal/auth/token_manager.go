package auth

import (
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/google/uuid"
)

/*
  interface for managing a validation token
*/
type TokenManager interface {
	// generate a new token
	CreateToken(shortUid uuid.UUID, durattion time.Duration, permissions []hrm.Permission) (string, error)
	// check if token is valid
	VerifyToken(token string) (*Payload, error)
}

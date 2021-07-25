package auth

import "time"

/*
  interface for managing a validation token
*/
type TokenManager interface {
	// generate a new token
	CreateToken(username string, durattion time.Duration) (string, error)
	// check if token is valid
	VerifyToken(token string) (*Payload, error)
}

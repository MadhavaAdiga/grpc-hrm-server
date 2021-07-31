package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpriedToken = errors.New("token expried")
var ErrInvalidToken = errors.New("token is invalid")

// implements jwt.Claims interface
type Payload struct {
	ID       uuid.UUID // id specific to the payload
	UserName string
	IssuedAt time.Time
	ExpireAt time.Time
}

/*
  constructor to create a new payload
  duration is used to specify the deadline
*/
func NewPayload(userName string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       id,
		UserName: userName,
		IssuedAt: time.Now(),
		ExpireAt: time.Now().Add(duration),
	}

	return payload, nil
}

// interface method implemetation
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpireAt) {
		return ErrExpriedToken
	}
	return nil
}

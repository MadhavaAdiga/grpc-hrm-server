package auth

import (
	"time"

	"github.com/MadhavaAdiga/grpc-hrm-server/protos/hrm"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

// implements maker interface
// v2
type PasetoManager struct {
	paseto        *paseto.V2
	asymmetricKey *AsymmetricStore
}

func NewPasetoManager() (TokenManager, error) {
	key := NewAsymmetricStore()

	manager := &PasetoManager{
		paseto:        paseto.NewV2(),
		asymmetricKey: key,
	}

	return manager, nil
}

// generate a new paseto token
func (manager *PasetoManager) CreateToken(shortUid uuid.UUID, duration time.Duration, permissions []hrm.Permission) (string, error) {
	// create a new payload
	payload, err := NewPayload(shortUid, duration, permissions)
	if err != nil {
		return "", err
	}

	// encrypt token with asymmetric key
	return manager.paseto.Sign(manager.asymmetricKey.privatekey, payload, nil)
}

// check if token is valid
func (manager *PasetoManager) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	// decrypt the token with asymmetric key
	err := manager.paseto.Verify(token, manager.asymmetricKey.publicKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	// validate payload
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

package auth

import (
	"time"

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
	// check if the key length is equal to required by the algorithm
	// if len(key.privatekey) != chacha20poly1305.KeySize {
	// 	return nil, fmt.Errorf("invalid key size: must be minimum of %d char", chacha20poly1305.KeySize)
	// }

	// log.Fatal((key.privatekey))

	manager := &PasetoManager{
		paseto:        paseto.NewV2(),
		asymmetricKey: key,
	}

	return manager, nil
}

// generate a new paseto token
func (manager *PasetoManager) CreateToken(username string, durattion time.Duration) (string, error) {
	// create a new payload
	payload, err := NewPayload(username, durattion)
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
	err := manager.paseto.Verify(token, manager.asymmetricKey.certificate, payload, nil)
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

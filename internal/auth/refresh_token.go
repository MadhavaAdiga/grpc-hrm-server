package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JSON web token manager
type RefreshTokenManager struct {
	secretKey     string
	tokenDuration time.Duration
}

// constructor
func NewRefreshTokenManager(secretKey string, duration time.Duration) *RefreshTokenManager {
	return &RefreshTokenManager{
		secretKey:     secretKey,
		tokenDuration: duration,
	}
}

// method to generate and sign a new token for a user
func (manager *RefreshTokenManager) Generate() (string, error) {
	// create claims for jwt
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
	}
	// in production use stronger methods of signing like RSA
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign and return
	return token.SignedString([]byte(manager.secretKey))
}

// method to verify jwt token
func (manager *RefreshTokenManager) Verify(jwtToken string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&jwt.StandardClaims{},
		// key function
		func(t *jwt.Token) (interface{}, error) {
			// check if the signing method matches
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// check claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

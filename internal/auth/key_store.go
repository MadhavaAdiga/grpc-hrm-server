package auth

import (
	"crypto"
	"encoding/hex"

	"golang.org/x/crypto/ed25519"
)

type AsymmetricStore struct {
	publicKey  crypto.PublicKey
	privatekey crypto.PrivateKey
}

func NewAsymmetricStore() *AsymmetricStore {
	b, _ := hex.DecodeString("b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	privKey := ed25519.PrivateKey(b)

	b, _ = hex.DecodeString("1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	pubKey := ed25519.PublicKey(b)

	return &AsymmetricStore{
		publicKey:  pubKey,
		privatekey: privKey,
	}
}

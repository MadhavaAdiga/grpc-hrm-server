package auth

import (
	"crypto"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/pkcs12"
)

type AsymmetricStore struct {
	certificate []byte
	privatekey  crypto.PrivateKey
}

func NewAsymmetricStore() *AsymmetricStore {
	p12_data, err := ioutil.ReadFile("../../cert/keystore.pkcs12")
	if err != nil {
		log.Fatal(err)
	}

	key, cert, err := pkcs12.Decode(p12_data, "grpcstore")
	if err != nil {
		log.Fatal(err)
	}

	return &AsymmetricStore{
		certificate: cert.Raw,
		privatekey:  key.(crypto.PrivateKey),
	}
}

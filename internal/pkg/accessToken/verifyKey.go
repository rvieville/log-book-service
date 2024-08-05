package accessToken

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

var RsaPublicKey *rsa.PublicKey

func ValidateKey() {
	b, _ := base64.StdEncoding.DecodeString(os.Getenv("JWT_PUBLIC_KEY"))
	block, _ := pem.Decode(b)
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err.Error())
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		panic("failed to parse RSA public key")
	}

	RsaPublicKey = rsaPublicKey
}

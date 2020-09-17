package security

import "encoding/base64"

// SecretKey is a struct.
type SecretKey struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// GenerateSecret will return base64 encoded string of private key and public key.
func GenerateSecret() (*SecretKey, error) {
	privateKey, publicKey, err := GenerateKey64()
	secretPriPubKey := SecretKey{
		PrivateKey: base64.StdEncoding.EncodeToString([]byte(privateKey)),
		PublicKey:  base64.StdEncoding.EncodeToString([]byte(publicKey)),
	}
	if err != nil {
		return nil, err
	}
	return &secretPriPubKey, nil
}

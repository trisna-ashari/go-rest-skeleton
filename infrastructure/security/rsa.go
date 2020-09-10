package security

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"go-rest-skeleton/infrastructure/message/exception"
)

const (
	rsaKeySize = 2048
)

func hash(data []byte) []byte {
	s := sha512.Sum512(data)
	return s[:]
}

// GenerateKey will generate private key with type *rsa.PrivateKey and public key with type *rsa.PublicKey.
func GenerateKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	pri, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return nil, nil, err
	}
	return pri, &pri.PublicKey, nil
}

// GenerateKeyBytes will generate private key and public key on byte type.
func GenerateKeyBytes() (privateBytes, publicBytes []byte, err error) {
	pri, pub, err := GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	priBytes, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		return nil, nil, err
	}
	pubBytes := x509.MarshalPKCS1PublicKey(pub)
	return priBytes, pubBytes, nil
}

// GenerateKey64 will generate private key and public key on base64 encoded string type.
func GenerateKey64() (pri64, pub64 string, err error) {
	pri, pub, err := GenerateKeyBytes()
	if err != nil {
		return "", "", nil
	}

	priKey := base64.StdEncoding.EncodeToString(pri)
	pubKey := base64.StdEncoding.EncodeToString(pub)
	priKeyPem := fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----\n", priKey)
	pubKeyPem := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----\n", pubKey)

	return priKeyPem,
		pubKeyPem,
		nil
}

// PublicKeyFrom will generate public key from byte.
func PublicKeyFrom(key []byte) (*rsa.PublicKey, error) {
	pubInterface, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return nil, err
	}
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, exception.ErrorTextInvalidPublicKey
	}
	return pub, nil
}

// PublicKeyFrom64 will generate public key from string.
func PublicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return PublicKeyFrom(b)
}

// PrivateKeyFrom will generate private key from byte.
func PrivateKeyFrom(key []byte) (*rsa.PrivateKey, error) {
	pri, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}
	p, ok := pri.(*rsa.PrivateKey)
	if !ok {
		return nil, exception.ErrorTextInvalidPrivateKey
	}
	return p, nil
}

// PrivateKeyFrom64 will generate private key from string.
func PrivateKeyFrom64(key string) (*rsa.PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFrom(b)
}

// PublicEncrypt will encrypt data with public key.
func PublicEncrypt(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, key, data)
}

// PublicSign will sign data with public key.
func PublicSign(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return PublicEncrypt(key, hash(data))
}

// PublicVerify will verify signed data with public key.
func PublicVerify(key *rsa.PublicKey, sign, data []byte) error {
	return rsa.VerifyPKCS1v15(key, crypto.SHA1, hash(data), sign)
}

// PrivateDecrypt will decrypt data with private key.
func PrivateDecrypt(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, key, data)
}

// PrivateSign will sign data with private key.
func PrivateSign(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, hash(data))
}

// PrivateVerify will verify signed data with private key.
func PrivateVerify(key *rsa.PrivateKey, sign, data []byte) error {
	h, err := PrivateDecrypt(key, sign)
	if err != nil {
		return err
	}
	if !bytes.Equal(h, hash(data)) {
		return rsa.ErrVerification
	}
	return nil
}

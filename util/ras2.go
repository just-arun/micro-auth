package util

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

const (
	rsaKeySize2 = 4096
)

type rsaStruct2 struct{}

func Rsa2() rsaStruct2 {
	return rsaStruct2{}
}

func (r rsaStruct2) Hash(data []byte) []byte {
	s := sha1.Sum(data)
	return s[:]

}
func (r rsaStruct2) GenerateKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	pri, err := rsa.GenerateKey(rand.Reader, rsaKeySize2)
	if err != nil {
		return nil, nil, err
	}
	return pri, &pri.PublicKey, nil
}

func (r rsaStruct2) GenerateKeyBytes() (privateBytes, publicBytes []byte, err error) {
	pri, pub, err := r.GenerateKey()
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

func (r rsaStruct2) GenerateKey64() (pri64, pub64 string, err error) {
	pri, pub, err := r.GenerateKeyBytes()
	if err != nil {
		return "", "", nil
	}
	return base64.StdEncoding.EncodeToString(pri),
		base64.StdEncoding.EncodeToString(pub),
		nil
}

func (r rsaStruct2) PublicKeyFrom(key []byte) (*rsa.PublicKey, error) {
	pubInterface, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return nil, err
	}
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid public key")
	}
	return pub, nil
}

func (r rsaStruct2) PublicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return r.PublicKeyFrom(b)
}

func (r rsaStruct2) PrivateKeyFrom(key []byte) (*rsa.PrivateKey, error) {
	pri, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}
	p, ok := pri.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("invalid private key")
	}
	return p, nil
}

func (r rsaStruct2) PrivateKeyFrom64(key string) (*rsa.PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return r.PrivateKeyFrom(b)
}

func (r rsaStruct2) PublicEncrypt(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, key, data)
}

func (r rsaStruct2) PublicSign(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return r.PublicEncrypt(key, hash(data))
}

func (r rsaStruct2) PublicVerify(key *rsa.PublicKey, sign, data []byte) error {
	return rsa.VerifyPKCS1v15(key, crypto.SHA1, hash(data), sign)
}

func (r rsaStruct2) PrivateDecrypt(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, key, data)
}

func (r rsaStruct2) PrivateSign(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, hash(data))
}

func (r rsaStruct2) PrivateVerify(key *rsa.PrivateKey, sign, data []byte) error {
	h, err := r.PrivateDecrypt(key, sign)
	if err != nil {
		return err
	}
	if !bytes.Equal(h, hash(data)) {
		return rsa.ErrVerification
	}
	return nil
}

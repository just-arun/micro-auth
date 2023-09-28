package util

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"log"

	"golang.org/x/crypto/ssh"
)

const (
	rsaKeySize = 2048
)

type rsaStruct struct{}

func Rsa() rsaStruct {
	return rsaStruct{}
}

func (r rsaStruct) Encrypt(base64PublicKey, valueToEncrypt string) (encryptedBase64String string, err error) {
	pubKey, err := r.PublicKeyFrom64(base64PublicKey)
	if err != nil {
		return "", err
	}
	encByte, err := r.PublicEncrypt(pubKey, []byte(valueToEncrypt))
	if err != nil {
		return "", err
	}
	encryptedBase64String = base64.StdEncoding.EncodeToString(encByte)
	return
}

func (r rsaStruct) Decrypt(base64PrivateKey, encryptedString string) (decryptedString string, err error) {
	privateKey, err := r.PrivateKeyFrom64(base64PrivateKey)
	if err != nil {
		return "", err
	}
	decByte, err := r.PrivateDecrypt(privateKey, []byte(encryptedString))
	if err != nil {
		return "", err
	}
	decryptedString = base64.StdEncoding.EncodeToString(decByte)
	return
}

func (r rsaStruct) GenerateKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	pri, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return nil, nil, err
	}
	return pri, &pri.PublicKey, nil
}

func (r rsaStruct) GenerateKeyBytes() (privateBytes, publicBytes []byte, err error) {
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

func (r rsaStruct) GenerateKey64() (pri64, pub64 string, err error) {
	pri, pub, err := r.GenerateKeyBytes()
	if err != nil {
		return "", "", nil
	}
	return base64.StdEncoding.EncodeToString(pri),
		base64.StdEncoding.EncodeToString(pub),
		nil
}

func (r rsaStruct) GeneratePublicKey(key *rsa.PrivateKey) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(&key.PublicKey)
}

func (r rsaStruct) GeneratePublicKeyBase64(key *rsa.PrivateKey) (string, error) {
	da, err := r.GeneratePublicKey(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(da), nil
}

func (r rsaStruct) PublicKeyFrom(key []byte) (*rsa.PublicKey, error) {
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

func GeneratePublicKey(privatekey *rsa.PublicKey) (string, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return "", err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	log.Println("Public key generated")
	return base64.StdEncoding.Strict().EncodeToString(pubKeyBytes), nil
}

func (r rsaStruct) PublicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return r.PublicKeyFrom(b)
}

func (r rsaStruct) PrivateKeyFrom(key []byte) (*rsa.PrivateKey, error) {
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

func (r rsaStruct) PrivateKeyFrom64(key string) (*rsa.PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return r.PrivateKeyFrom(b)
}

func (r rsaStruct) PublicEncrypt(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, key, data)
}

func (r rsaStruct) PublicEncryptBase64(key *rsa.PublicKey, data []byte) (string, error) {
	da, err := r.PublicEncrypt(key, data)
	return base64.StdEncoding.EncodeToString(da), err
}

func (r rsaStruct) PublicSign(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return r.PublicEncrypt(key, hash(data))
}

func (r rsaStruct) PublicVerify(key *rsa.PublicKey, sign, data []byte) error {
	return rsa.VerifyPKCS1v15(key, crypto.SHA1, hash(data), sign)
}

func (r rsaStruct) PrivateDecrypt(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, key, data)
}

func (r rsaStruct) PrivateDecryptFromBase64(key *rsa.PrivateKey, base64String string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	return r.PrivateDecrypt(key, data)
}

func (r rsaStruct) PrivateSign(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, hash(data))
}

func (r rsaStruct) PrivateVerify(key *rsa.PrivateKey, sign, data []byte) error {
	h, err := r.PrivateDecrypt(key, sign)
	if err != nil {
		return err
	}
	if !bytes.Equal(h, hash(data)) {
		return rsa.ErrVerification
	}
	return nil
}

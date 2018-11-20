package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RSAEncrypt 加密
func RSAEncrypt(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("无效的publicKey")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, key.(*rsa.PublicKey), origData)
}

// RSADecrypt 解密
func RSADecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("无效的privateKey")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, key, ciphertext)
}

func PKCS1PrivateKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("无效的privateKey")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func PKCS1PublicKey(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("无效的privateKey")
	}
	pkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pkey.(*rsa.PublicKey), nil
}

//
func PKCS8PrivateKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("无效的privateKey")
	}
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	} else {
		return key.(*rsa.PrivateKey), nil
	}
}

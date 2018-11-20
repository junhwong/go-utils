package crypto

import (
	"errors"
	"strings"
)

var ciphers = map[string]Cipher{}

type Cipher interface {
	Encrypt(data []byte, key ...[]byte) ([]byte, error)
	Decrypt(data []byte, key ...[]byte) ([]byte, error)
}

func init() {
	ciphers["BASE64"] = &Base64Cipher{}
	ciphers["AES-128-CBC_PKCS#7"] = &AESCipher{
		Len:        128,
		PaddingLen: 7,
	}
	ciphers["AES-128-CBC"] = ciphers["AES-128-CBC_PKCS#7"]
}

func Encrypt(data []byte, method string, key ...[]byte) ([]byte, error) {
	if cipher, ok := ciphers[strings.ToUpper(method)]; ok {
		return cipher.Encrypt(data, key...)
	}
	return nil, errors.New("Cipher Undefined:" + method)
}

func Decrypt(data []byte, method string, key ...[]byte) ([]byte, error) {
	if cipher, ok := ciphers[strings.ToUpper(method)]; ok {
		return cipher.Decrypt(data, key...)
	}
	return nil, errors.New("Cipher Undefined:" + method)
}

package crypto

import (
	"encoding/base64"
	"fmt"
)

type Base64Cipher struct {
}

func (*Base64Cipher) Encrypt(data []byte, key ...[]byte) ([]byte, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("参数不匹配:%d", len(key))
	}
	s := base64.StdEncoding.EncodeToString(data)
	return []byte(s), nil
}

func (*Base64Cipher) Decrypt(data []byte, key ...[]byte) ([]byte, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("参数不匹配:%d", len(key))
	}
	return base64.StdEncoding.DecodeString(string(data))
}

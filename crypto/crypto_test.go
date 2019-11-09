package crypto_test

import (
	"encoding/base64"
	"testing"

	"github.com/junhwong/go-utils/crypto"
)

func TestAes(t *testing.T) {
	text := "hello,世界"
	key, _ := crypto.Decrypt([]byte("tiihtNczf5v6AKRyjwEUhQ=="), "base64")
	iv, _ := base64.StdEncoding.DecodeString("r7BXXKkLb8qrSNn05n0qiA==")
	eData, err := crypto.Encrypt([]byte(text), "AES-128-CBC", key, iv)
	if err != nil {
		t.Fatal(err)
	}
	pData, err := crypto.Decrypt(eData, "AES-128-CBC", key, iv)
	if err != nil {
		t.Fatal(err)
	}
	if string(pData) != text {
		t.Fatal(string(pData))
	}
}

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type AESCipher struct {
	Len        int
	PaddingLen int
}

/*CBC加密 按照golang标准库的例子代码
不过里面没有填充的部分,所以补上
*/

//使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func aesCBCEncrypt(rawData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//填充原文
	blockSize := block.BlockSize()
	rawData = PKCS7Padding(rawData, blockSize)
	//初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, len(rawData))
	//block大小 16
	if iv == nil {
		iv = key[:blockSize]
	}

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	//block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, rawData)

	return cipherText, nil
}

func aesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	if iv == nil {
		iv = key[:blockSize]
	}
	//encryptData = encryptData[blockSize:]
	decryptData := make([]byte, len(encryptData))
	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(decryptData, encryptData)
	//解填充
	decryptData = PKCS7UnPadding(decryptData)
	return decryptData, nil
}

func (c *AESCipher) Encrypt(data []byte, key ...[]byte) ([]byte, error) {
	if len(key) == 2 {
		return aesCBCEncrypt(data, key[0], key[1])
	} else if len(key) == 1 {
		return aesCBCEncrypt(data, key[0], nil)
	}

	return nil, fmt.Errorf("参数不匹配:%d", len(key))
}

func (c *AESCipher) Decrypt(data []byte, key ...[]byte) ([]byte, error) {
	if len(key) == 2 {
		return aesCBCDecrypt(data, key[0], key[1])
	} else if len(key) == 1 {
		return aesCBCDecrypt(data, key[0], nil)
	}

	return nil, fmt.Errorf("参数不匹配:%d", len(key))
}

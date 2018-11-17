package gohelper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"
	"time"
)

type AESCipher struct {
	cipher cipher.Block
}

func NewAESCipher(key string) (*AESCipher, error) {
	block, err := aes.NewCipher([]byte(key))
	if nil != err {
		return nil, err
	}

	return &AESCipher{block}, nil
}

func (a *AESCipher) CBCEncrypt(plaintext string) string {
	pt := padPlaintext([]byte(plaintext), a.cipher.BlockSize())
	iv := createIV(a.cipher.BlockSize())

	encrypter := cipher.NewCBCEncrypter(a.cipher, iv)
	dst := make([]byte, len(pt))
	encrypter.CryptBlocks(dst, pt)

	dst = append(iv, dst...)
	return base64.StdEncoding.EncodeToString(dst)
}

func (a *AESCipher) CBCDecrypt(ciphertext string) string {
	ct, err := base64.StdEncoding.DecodeString(ciphertext)

	if nil != err {
		return ""
	}

	iv := ct[:a.cipher.BlockSize()]
	decrypter := cipher.NewCBCDecrypter(a.cipher, iv)
	plaintext := make([]byte, len(ct)-a.cipher.BlockSize())

	decrypter.CryptBlocks(plaintext, ct[a.cipher.BlockSize():])
	l := len(plaintext) - int(plaintext[len(plaintext)-1])

	return string(plaintext[:l])
}

func padPlaintext(plaintext []byte, size int) []byte {
	padding := size - len(plaintext)%size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(plaintext, padtext...)
}

func createIV(size int) []byte {
	iv := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		iv[i] = byte(rand.Intn(256))
	}

	return iv
}

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getCipherCode() string {
	if err := godotenv.Load(); err != nil {
		log.Println("NO .env found")
	}

	code := os.Getenv("CIPHER_CODE")
	return code
}

func EncryptCipher(textToEncrypt string) (string, string) {
	code := getCipherCode()

	key, _ := hex.DecodeString(code)
	text := []byte(textToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesGcm, err := cipher.NewGCMWithNonceSize(block, 12)
	if err != nil {
		panic(err.Error())
	}

	cipherText := aesGcm.Seal(nil, nonce, text, nil)
	return string(nonce), string(cipherText)
}

func DecryptCipher(nonceKey string, cipherKey string) string {
	code := getCipherCode()

	key, _ := hex.DecodeString(code)
	text, _ := hex.DecodeString(cipherKey)
	nonce, _ := hex.DecodeString(nonceKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainText, err := aesGcm.Open(nil, nonce, text, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plainText)
}

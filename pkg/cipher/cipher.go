package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func Ecnrypt(key []byte, data []byte) string {
	byteMsg := []byte(data)
	block, err := aes.NewCipher(key)

	if err != nil {
		log.Fatalf("could not create new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalf("could not create new cipher: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText)
}

func Decrypt(key []byte, message string) string {
	cipherText, err := base64.StdEncoding.DecodeString(message)

	if err != nil {
		log.Fatalf("could not create new cipher: %v", err)
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		log.Fatalf("could not create new cipher: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		log.Fatalf("could not create new cipher: %v", err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}

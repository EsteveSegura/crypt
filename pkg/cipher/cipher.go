package cipher

import (
	b64 "encoding/base64"
)

// This is not encryption
// Just encode o decode
func Ecnrypt(data []byte) string {
	ciphertext := b64.StdEncoding.EncodeToString([]byte(data))
	return ciphertext
}

func Decrypt(data string) []byte {
	plaintext, _ := b64.StdEncoding.DecodeString(data)
	return plaintext
}
